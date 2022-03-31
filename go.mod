module weather-api

go 1.16

require (
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/echo/v4 v4.7.2
	golang.org/x/crypto v0.0.0-20220321153916-2c7772ba3064 // indirect
	gorm.io/driver/mysql v1.3.2
	gorm.io/gorm v1.23.3
)

replace weather-api/controller => ./controller
replace weather-api/view_model => ./view_model