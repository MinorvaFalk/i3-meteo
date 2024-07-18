package model

type Weather struct {
	Lat       string         `json:"lat"`
	Lon       string         `json:"lon"`
	Elevation int            `json:"elevation"`
	Timezone  string         `json:"timezone"`
	Units     string         `json:"units"`
	Current   WeatherCurrent `json:"current"`
	Hourly    struct {
		Data []WeatherHourly `json:"data"`
	} `json:"hourly"`
	Daily struct {
		Data []WeatherDaily `json:"data"`
	} `json:"daily"`
}

type WeatherCurrent struct {
	Icon          string        `json:"icon"`
	IconNum       int           `json:"icon_num"`
	Summary       string        `json:"summary"`
	Temperature   float32       `json:"temperature"`
	Wind          Wind          `json:"wind"`
	Precipitation Precipitation `json:"precipitation"`
	CloudCover    int           `json:"cloud_cover"`
}

type WeatherHourly struct {
	Date          string        `json:"date"` // date format 2024-07-17T14:00:00
	Weather       string        `json:"weather"`
	Icon          int           `json:"icon"`
	Summary       string        `json:"summary"`
	Temperature   float32       `json:"temperature"`
	Wind          Wind          `json:"wind"`
	CloudCover    CloudCover    `json:"cloud_cover"`
	Precipitation Precipitation `json:"precipitation"`
}

type WeatherDaily struct {
	Day       string  `json:"date"` // day format 2024-07-17
	Weather   string  `json:"weather"`
	Icon      int     `json:"icon"`
	Summary   string  `json:"summary"`
	AllDay    DayData `json:"all_day"`
	Morning   DayData `json:"morning"`
	Afternoon DayData `json:"afternoon"`
	Evening   DayData `json:"evening"`
}

type DayData struct {
	Weather        string        `json:"weather"`
	Icon           int           `json:"icon"`
	Temperature    float32       `json:"temperature"`
	TemperatureMin float32       `json:"temperature_min"`
	TemperatureMax float32       `json:"temperature_max"`
	Wind           Wind          `json:"wind"`
	CloudCover     CloudCover    `json:"cloud_cover"`
	Precipitation  Precipitation `json:"precipitation"`
}

type Wind struct {
	Speed float32 `json:"speed"`
	Angle int     `json:"angle"`
	Dir   string  `json:"dir"`
}

type Precipitation struct {
	Total float32 `json:"total"`
	Type  string  `json:"type"`
}

type CloudCover struct {
	Total int `json:"total"`
}
