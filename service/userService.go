package service

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (service *Service) GetUser(c echo.Context) error {
	id := c.Param("id")
	userId, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		return err
	}

	user, err := service.user.GetById(uint(userId))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, *user)
}
