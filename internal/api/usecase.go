package api

import (
	"context"
	"i3/internal/entity"
	"i3/internal/model"
)

type Usecase interface {
	ScheduleLocationData(ctx context.Context, req model.CityRequest) (model.ScheduleLocationData, error)

	FetchWeatherData(ctx context.Context) ([]*model.WeatherData, error)
	FetchCities(ctx context.Context) ([]*entity.Cities, error)
}
