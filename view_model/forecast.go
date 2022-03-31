package view_model

type Forecast struct {
	Cod     string             `json:"cod"`
	Message float64            `json:"message"`
	Cnt     int                `json:"cnt"`
	List    []ForecastListData `json:"list"`
	City    CityForecast       `json:"city"`
}
