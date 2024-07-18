package entity

import (
	"time"

	"github.com/google/uuid"
)

type Weathers struct {
	ID                 uuid.UUID  `json:"id"`
	City               string     `json:"city"`
	Lat                float32    `json:"lat"`
	Lon                float32    `json:"lon"`
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
