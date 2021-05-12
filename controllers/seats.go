package controllers

import (
	"cinema/services"
	"net/http"

	"github.com/labstack/echo"
)

// GetSeats: Get the seats that match the number of seats to be booked
func GetSeats(c echo.Context) error {
	seats := c.QueryParam("seats")

	listSeatSuitable, totalRows, totalCols := services.HandleGetSeats(seats)
	response := handleResponse(listSeatSuitable, totalRows, totalCols)

	return c.JSON(http.StatusOK, response)
}

// handleResponse: Create response to return
func handleResponse(listSeatSuitable map[int]interface{}, totalRows int, totalCols int) map[string]interface{} {
	response := make(map[string]interface{})
	response["rows"] = totalRows
	response["cols"] = totalCols

	seats := make(map[int]interface{})
	for key := 0; key < len(listSeatSuitable); key++ {
		seats[key+1] = listSeatSuitable[key]
	}

	response["seat_position"] = seats
	return response
}
