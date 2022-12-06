package routes

import (
	"depd_go_echo/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		var res models.Response
		res.Code = http.StatusOK
		res.Status = true
		res.Message = "Hello, this is echo!"

		return c.JSON(http.StatusOK, res)
	})

	return e
}
