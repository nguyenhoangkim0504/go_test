package main

import (
	"cinema/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// Routes
	e.GET("/positions", controllers.GetSeats)

	e.Logger.Fatal(e.Start(":8080"))
}
