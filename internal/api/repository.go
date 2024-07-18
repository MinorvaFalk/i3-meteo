package api

import (
	"context"
	"i3/internal/entity"
	"i3/internal/model"
)

type Repository interface {
	InsertCity(ctx context.Context, req model.CityRequest) (data model.ScheduleLocationData, err error)
	InsertWeatherJob(ctx context.Context, req entity.Weathers) (err error)

	FetchCities(ctx context.Context) ([]*entity.Cities, error)
	FetchWeatherData(ctx context.Context) ([]*model.WeatherData, error)
}
