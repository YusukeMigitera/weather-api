package main

import (
	"net/http"

	"weather-api/controller"
	// "weather-api/view_model"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/forecast", controller.GetForecast)

	e.Logger.Fatal(e.Start(":8080"))
}
