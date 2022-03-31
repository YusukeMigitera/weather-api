package view_model

type Main struct {
	Temp      float32 `json:"temp"`
	TempMin   float32 `json:"temp_min"`
	TempMax   float32 `json:"temp_max"`
	Pressure  float32 `json:"pressure"`
	SeaLevel  float32 `json:"sea_level"`
	GrndLevel float32 `json:"grnd_level"`
	Humidity  int     `json:"humidity"`
	Temp_kf   float32 `json:"temp_kf"`
}
