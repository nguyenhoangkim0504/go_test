package controllers

import (
	"cinema/services"
	"cinema/utils"
	"net/http"

	"github.com/labstack/echo"
)

// GetSeats: Get the seats that match the number of seats to be booked
func GetSeats(c echo.Context) error {
	room := c.QueryParam("room")
	seats := c.QueryParam("seats")

	listSeatEmpty := services.HandleGetSeats(room, seats)
	response := handleResponse(listSeatEmpty)

	return c.JSON(http.StatusOK, response)
}

// RegisterSeats: Register seats
func RegisterSeats(c echo.Context) error {
	room := c.QueryParam("room")
	seats, err := utils.ParseReqBody(c)
	checkError(err)

	status := services.HandleRegisterSeats(room, seats)
	responseStatus := http.StatusOK
	responseMsg := make(map[string]string)
	responseMsg["Status"] = "Register successfully"
	if !status {
		responseStatus = http.StatusBadRequest
		responseMsg["Status"] = "Register failed"
	}

	return c.JSON(responseStatus, responseMsg)
}

// handleResponse: Create response to return
func handleResponse(listSeatEmpty map[int]interface{}) map[string]interface{} {
	response := make(map[string]interface{})
	response["seat_number"] = len(listSeatEmpty)

	seats := make(map[int]interface{})
	for key := 0; key < len(listSeatEmpty); key++ {
		seats[key+1] = listSeatEmpty[key]
	}

	response["seat_position"] = seats
	return response
}

// checkError: Check error
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
