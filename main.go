package main

import (
	"cinema/controllers"
	"cinema/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// Routes
	e.GET("/positions", controllers.GetSeats)
	e.POST("/positions", controllers.RegisterSeats)

	e.Logger.Fatal(e.Start(":8080"))
}

func init() {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "cinema"
	models.ConnectToDB(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
}
