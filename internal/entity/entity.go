package entity

import (
	"time"

	"github.com/google/uuid"
)

type Cities struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Lat       float32    `json:"lat"`
	Lon       float32    `json:"lon"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func (Cities) TableName() string {
	return "cities"
}

type Weathers struct {
	ID                 int64      `json:"id"`
	Date               string     `json:"date"`
	CityID             string     `json:"city_id"`
	Summary            *string    `json:"summary"`
	Temperature        *float32   `json:"temperature"`
	WindSpeed          *float32   `json:"wind_speed"`
	WindAngle          *int       `json:"wind_angle"`
	WindDir            *string    `json:"wind_dir"`
	PrecipitationTotal *float32   `json:"precipitation_total"`
	PrecipitationType  *string    `json:"precipitation_type"`
	CloudCover         *int       `json:"cloud_cover"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
	DeletedAt          *time.Time `json:"deleted_at,omitempty"`
}

func (Weathers) TableName() string {
	return "weathers"
}
