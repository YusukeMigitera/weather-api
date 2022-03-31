package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"weather-api/model"
	"weather-api/view_model"

	"github.com/labstack/echo/v4"
)

func GetForecast(c echo.Context) error {
	fmt.Print("id:")
	i, _ := strconv.Atoi(c.QueryParam("q"))
	id := uint(i)
	fmt.Println(id)
	weather := model.Weather{}
	weather.FirstById(id)

	city := view_model.CityForecast{}

	forecast := view_model.Forecast{
		Cod:     "200",
		Message: 0.011,
		Cnt:     40,
		List:    []string{"0", "1", "2"},
		City:    city,
	}

	return c.JSON(http.StatusOK, forecast)
}
