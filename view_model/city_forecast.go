package view_model

type CityForecast struct {
    Id int `json:"id"`
    Name string `json:"name"`
    Country string `json:"country"`
    Population int `json:"population"`
    Timezone int `json:"timezon`
    Sunrise int `json:"sunrise"`
    Sunset int `json:"sunset"`
}