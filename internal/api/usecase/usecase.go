package apiusecase

import (
	"context"
	"i3/internal/api"
	"i3/internal/model"
)

type usecase struct {
	repo api.Repository
}

func New(repo api.Repository) api.Usecase {
	return &usecase{
		repo: repo,
	}
}

func (u *usecase) ScheduleLocationData(ctx context.Context, req model.CityRequest) (data model.ScheduleLocationData, err error) {
	return u.repo.InsertCity(ctx, req)
}

func (u *usecase) FetchWeatherData(ctx context.Context) ([]*model.WeatherData, error) {
	return u.repo.FetchWeatherData(ctx)
}
