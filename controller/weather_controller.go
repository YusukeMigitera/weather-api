package controller

import (
	"fmt"
	"net/http"
	"weather-api/model"
	"weather-api/view_model"

	"github.com/labstack/echo/v4"
)

func GetForecast(c echo.Context) error {
	place := c.QueryParam("q")
	fmt.Printf("%s, ", place)

	var weathers []model.Weather
	weather := model.Weather{}
	weather.Limit(&weathers)
	fmt.Println(weathers)

	forecastList := []view_model.ForecastListData{}
	for _, weather := range weathers {
		forecastList = append(forecastList, view_model.ForecastListData{
			Dt: weather.Dt,
			Main: view_model.Main{
				Temp:      weather.TempAverage,
				TempMin:   weather.TempMin,
				TempMax:   weather.TempMax,
				Pressure:  1013.00,
				SeaLevel:  1013.00,
				GrndLevel: 1011.00,
				Humidity:  weather.Humidity,
				Temp_kf:   1.0,
			},
			Weather: []view_model.Weather{{
				Id:          800,
				Main:        weather.Weather,
				Description: weather.Weather + " sky",
				Icon:        weather.WeatherIcon,
			}},
			Clouds: view_model.Clouds{All: 0},
			Wind: view_model.Wind{
				Speed: weather.WindSpeed,
				Deg:   275.935,
			},
			Sys:   view_model.Sys{Pod: "n"},
			DtTxt: "2019-10-05 03:00:00",
		})
	}

	forecast := view_model.Forecast{
		Cod:     "200",
		Message: 0.011,
		Cnt:     3,
		List:    forecastList,
		City:    view_model.CityForecast{},
	}

	return c.JSON(http.StatusOK, forecast)
}
