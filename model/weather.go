package model

import (
	"gorm.io/gorm"
)

type Weather struct {
	Id          int     `json:"id"`
	Dt          int     `json:"dt"`
	Humidity    int     `json:"humidity"`
	WindSpeed   float32 `json:"wind_speed"`
	TempAverage float32 `json:"temp_average"`
	TempMax     float32 `json:"temp_max"`
	TempMin     float32 `json:"temp_min"`
	Weather     string  `json:"weather"`
	WeatherIcon string  `json:"weather_icon"`
}

func (p *Weather) Limit(slice *[]Weather) (tx *gorm.DB) {
	return DB.Limit(3).Find(slice)
}
