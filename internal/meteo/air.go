package meteo

import "math/big"

type Air struct {
	Lat       string    `json:"lat"`
	Lon       string    `json:"lon"`
	Elevation string    `json:"elevation"`
	Timezone  string    `json:"timezone"`
	Data      []AirData `json:"data"`
}

type AirData struct {
	Date              string    `json:"date"` // date format 2024-05-03T11:00:00
	Aerosol550        float32   `json:"aerosol_550"`
	Airquality        int       `json:"air_quality"`
	Cosurface         float32   `json:"co_surface"`
	Dust550nm         float32   `json:"dust_550nm"`
	Dustmixingratio05 big.Float `json:"dust_mixing_ratio_05"`
	No2surface        float32   `json:"no2_surface"`
	Nosurface         float32   `json:"no_surface"`
	Ozonesurface      float32   `json:"ozone_surface"`
	Ozonetotal        float32   `json:"ozone_total"`
	Pm10              float32   `json:"pm10"`
	Pm25              float32   `json:"pm25"`
	So2surface        float32   `json:"so2_surface"`
}
