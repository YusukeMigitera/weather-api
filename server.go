package main

import (
    "net/http"

    "github.com/labstack/echo"
    "./view_model"
)

func main() {
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
	e.GET("/forecast", Forecast)

    e.Logger.Fatal(e.Start(":1323"))
}

func Forecast(c echo.Context) error {
	city := view_model.CityForecast{}

    forecast := view_model.Forecast{
        "200",
        0.011,
        40,
        []string{"0", "1", "2"},
        city,
    }

	if err := c.Bind(forecast); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, forecast)
}