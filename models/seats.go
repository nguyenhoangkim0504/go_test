package models

import (
	"database/sql"
)

var db *sql.DB

func ConnectToDB(driverName string, dataSourceName string) {
	var err error
	db, err = sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
}

func GetSeats(room string) (*sql.Rows, error) {
	result, err := db.Query("SELECT * FROM seats WHERE room=?", room)
	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func GetRoom(room string) (*sql.Rows, error) {
	result, err := db.Query("SELECT * FROM rooms WHERE room=?", room)
	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func RegisterSeat(room string, row int, col int) (sql.Result, error) {
	result, err := db.Exec("INSERT INTO seats(room, row, col) VALUE (?, ?, ?)", room, row, col)
	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}
