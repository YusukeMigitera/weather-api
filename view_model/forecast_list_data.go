package view_model

type ForecastListData struct {
	Dt      int       `json:"dt"`
	Main    Main      `json:"main"`
	Weather []Weather `json:"weather"`
	Clouds  Clouds    `json:"clouds"`
	Wind    Wind      `json:"wind"`
	Sys     Sys       `json:"sys"`
	DtTxt   string    `json:"dt_txt"`
}
