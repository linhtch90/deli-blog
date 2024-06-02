package main

import (
	"deli-blog/dbConnection"
	"deli-blog/persistance"
	"deli-blog/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	db := dbConnection.New()
	dbConnection.AutoMigrate(db)

	userPersistance := persistance.NewUserPersistance(db)
	postPersistance := persistance.NewPostPersistance(db)
	service := service.NewService(userPersistance, postPersistance)

	apiVersion := e.Group("api/v1")
	service.RegisterController(apiVersion)

	e.Logger.Fatal(e.Start(":12000"))
}
