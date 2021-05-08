package services

import (
	"cinema/entities"
	"cinema/models"
	"database/sql"
	"fmt"
	"math"
	"strconv"

	"github.com/mitchellh/mapstructure"
)

const MIN_DISTANCE = 7

// Define position struct
type Position struct {
	Row int `json:"row"`
	Col int `json:"rol"`
}

// HandleGetSeats: Handle get list seats empty
func HandleGetSeats(room string, seats string) map[int]interface{} {
	listSeatEmpty := map[int]interface{}{}

	// Get room info
	roomSelected, err := models.GetRoom(room)
	if err != nil {
		panic(err)
	}
	defer roomSelected.Close()
	roomEntity := scanDataRoom(roomSelected)

	// Select seat from DB
	seatSelected, err := models.GetSeats(room)
	if err != nil {
		panic(err)
	}
	defer seatSelected.Close()
	listSeatEntity := scanDataSeat(seatSelected)

	numberOfSeats, _ := strconv.Atoi(seats)
	if numberOfSeats == 1 {
		listSeatEmpty = getListSeatEmpty(roomEntity, listSeatEntity, false)
	} else if numberOfSeats > 1 {

		listSeatEmpty = getListSeatEmpty(roomEntity, listSeatEntity, true)
	}

	return listSeatEmpty
}

// HandleRegisterSeats: Handle register seat
func HandleRegisterSeats(room string, body map[string]interface{}) bool {
	seats := body["seats"]

	listNewPoint := make(map[int]Position)
	index := 0
	for _, seat := range seats.([]interface{}) {
		point := Position{}
		mapstructure.Decode(seat, &point)
		listNewPoint[index] = point
		index++
	}

	// Get room info
	roomSelected, err := models.GetRoom(room)
	if err != nil {
		panic(err)
	}
	defer roomSelected.Close()
	roomEntity := scanDataRoom(roomSelected)

	// Select seat from DB
	seatSelected, err := models.GetSeats(room)
	if err != nil {
		panic(err)
	}
	defer seatSelected.Close()
	listSeatEntity := scanDataSeat(seatSelected)

	listSeatClosed := make(map[int]Position)
	seatClosedNumber := len(listSeatEntity)
	for i := 0; i < seatClosedNumber; i++ {
		point := Position{listSeatEntity[i].Row, listSeatEntity[i].Col}
		listSeatClosed[i] = point
	}

	numberOfSeats := len(listNewPoint)
	if numberOfSeats == 1 {
		if listNewPoint[0].Row > roomEntity.Rows || listNewPoint[0].Col > roomEntity.Cols {
			return false
		}

		if !checkMinDistance(listSeatClosed, listNewPoint[0]) {
			return false
		}

		_, err := models.RegisterSeat(room, listNewPoint[0].Row, listNewPoint[0].Col)
		if err != nil {
			panic(err)
		}
	}

	if numberOfSeats > 1 {
		for i := 0; i < numberOfSeats; i++ {
			if listNewPoint[0].Row > roomEntity.Rows || listNewPoint[0].Col > roomEntity.Cols {
				return false
			}

			if !checkMinDistance(listSeatClosed, listNewPoint[i]) {
				return false
			}

			_, err := models.RegisterSeat(room, listNewPoint[i].Row, listNewPoint[i].Col)
			if err != nil {
				panic(err)
			}
		}
	}

	return true
}

// scanDataSeat: Mapping data to entity
func scanDataSeat(rows *sql.Rows) map[int]entities.Seat {
	seats := make(map[int]entities.Seat)
	index := 0
	for rows.Next() {
		seat := entities.Seat{}
		err := rows.Scan(&seat.Room, &seat.Seat, &seat.Row, &seat.Col)
		if err != nil {
			fmt.Println(err)
			return map[int]entities.Seat{}
		}
		seats[index] = seat
		index++
	}

	return seats
}

// scanDataRoom: Mapping data to entity
func scanDataRoom(rows *sql.Rows) entities.Room {
	room := entities.Room{}
	for rows.Next() {
		err := rows.Scan(&room.Room, &room.Rows, &room.Cols)
		if err != nil {
			fmt.Println(err)
			return entities.Room{}
		}
	}

	return room
}

// getListSeatEmpty: get suitable empty seats
func getListSeatEmpty(roomEntity entities.Room, listSeatEntity map[int]entities.Seat, group bool) map[int]interface{} {
	rows := roomEntity.Rows
	cols := roomEntity.Cols

	listSeatClosed := make(map[int]Position)
	listSeatEmpty := make(map[int]interface{})

	seat := 0
	seatClosedNumber := len(listSeatEntity)
	if seatClosedNumber == 0 {
		listSeatEmpty[seat] = Position{0, 0}
	} else {
		for i := 0; i < seatClosedNumber; i++ {
			point := Position{listSeatEntity[i].Row, listSeatEntity[i].Col}
			listSeatClosed[i] = point
		}
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// Check that the point under consideration meets the min distance
			newPoint := Position{row, col}
			if checkMinDistance(listSeatClosed, newPoint) {
				listSeatEmpty[seat] = newPoint
				seat += 1
				if !group {
					listSeatClosed[seat] = newPoint
				}
			}
		}
	}

	return listSeatEmpty
}

// checkMinDistance: Check the distance between new arrivals and approved points
func checkMinDistance(listPosition map[int]Position, newPoint Position) bool {
	status := true
	for _, oldPoint := range listPosition {
		distance := math.Abs(float64(newPoint.Row)-float64(oldPoint.Row)) +
			math.Abs(float64(newPoint.Col)-float64(oldPoint.Col))
		if distance < MIN_DISTANCE {
			status = false
		}
	}

	return status
}
