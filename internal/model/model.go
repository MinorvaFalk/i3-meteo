package model

type WeatherData struct {
	ID                 string   `json:"id"`
	Name               string   `json:"name"`
	Lat                float32  `json:"lat"`
	Lon                float32  `json:"lon"`
	Date               *string  `json:"date"`
	Summary            *string  `json:"summary"`
	Temperature        *float32 `json:"temperature"`
	WindSpeed          *float32 `json:"wind_speed"`
	WindAngle          *int     `json:"wind_angle"`
	WindDir            *string  `json:"wind_dir"`
	PrecipitationTotal *float32 `json:"precipitation_total"`
	PrecipitationType  *string  `json:"precipitation_type"`
	CloudCover         *int     `json:"cloud_cover"`
}

type ScheduleLocationData struct {
	ID   string  `json:"id"`
	City string  `json:"city"`
	Lat  float32 `json:"lat"`
	Lon  float32 `json:"lon"`
}
