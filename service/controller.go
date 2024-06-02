package service

import "github.com/labstack/echo/v4"

const USER_GROUP = "/user"

func (service *Service) RegisterController(apiVersion *echo.Group) {
	userApi := apiVersion.Group(USER_GROUP)
	userApi.GET("/:id", service.GetUser)

}
