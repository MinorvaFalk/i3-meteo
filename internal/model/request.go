package model

type CityRequest struct {
	Name string  `json:"name" binding:"required"`
	Lat  float32 `json:"lat" binding:"required"`
	Lon  float32 `json:"lon" binding:"required"`
}
