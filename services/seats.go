package services

import (
	"math"
	"strconv"
)

const MIN_DISTANCE = 7

// Define position struct
type Position struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

// HandleGetSeats: Handle get list seats empty
func HandleGetSeats(seats string) (map[int]interface{}, int, int) {
	numberOfSeats, _ := strconv.Atoi(seats)
	listSeatSuitable := make(map[int]Position)

	switch numberOfSeats {
	case 0:
		return nil, 0, 0
	case 1:
		listSeatSuitable[1] = Position{0, 0}
		return mappingInterface(listSeatSuitable), 1, 1
	default:
		totalRows := 1
		totalCols := 1
		seat := 0

		colChecking := 0
		for row := 0; row < totalRows && seat < numberOfSeats; row++ {
			for col := 0; col <= row+1 && seat < numberOfSeats; col++ {
				// Check that the point under consideration meets the min distance
				newPoint := Position{row, col}
				if checkMinDistance(listSeatSuitable, newPoint) {
					listSeatSuitable[seat] = newPoint
					seat = seat + 1
					if totalCols < col {
						totalCols = col
					}
				}
				colChecking = col
			}

			// Loop running backwards
			for rowChecking := row; colChecking != 0 && rowChecking >= 0 && seat < numberOfSeats; rowChecking-- {
				// Check that the point under consideration meets the min distance
				newPoint := Position{rowChecking, colChecking}
				if checkMinDistance(listSeatSuitable, newPoint) {
					listSeatSuitable[seat] = newPoint
					seat += 1
					if totalCols < colChecking {
						totalCols = colChecking
					}
				}
			}

			totalRows += 1
		}

		return mappingInterface(listSeatSuitable), totalRows - 1, totalCols + 1
	}
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

func mappingInterface(list map[int]Position) map[int]interface{} {
	result := make(map[int]interface{})
	for key, position := range list {
		result[key] = position
	}

	return result
}
