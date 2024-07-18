package scheduler

import (
	"context"
	"fmt"
	"i3/internal/entity"
	"i3/internal/meteo"
	"i3/internal/repository"
	"i3/pkg/logger"
	"strconv"

	"go.uber.org/zap"
)

type WeatherJob struct {
	m    *meteo.Meteo
	repo *repository.Repository
}

func NewWeatherJob(m *meteo.Meteo, repo *repository.Repository) *WeatherJob {
	return &WeatherJob{
		m:    m,
		repo: repo,
	}
}

func (j *WeatherJob) Run() {
	ctx := context.Background()

	weathers, err := j.repo.FetchWeatherJobs(ctx)
	if err != nil {
		panic(err)
	}

	for _, weather := range weathers {
		go func(wet *entity.Weathers) {
			lat := strconv.FormatFloat(float64(wet.Lat), 'f', 11, 64)
			lon := strconv.FormatFloat(float64(wet.Lon), 'f', 11, 64)

			data, err := j.m.FetchLocationWeather(lat, lon)
			if err != nil {
				panic(err)
			}

			wet.Summary = &data.Current.Summary
			wet.Temperature = &data.Current.Temperature
			wet.WindSpeed = &data.Current.Wind.Speed
			wet.WindAngle = &data.Current.Wind.Angle
			wet.WindDir = &data.Current.Wind.Dir
			wet.PrecipitationTotal = &data.Current.Precipitation.Total
			wet.PrecipitationType = &data.Current.Precipitation.Type
			wet.CloudCover = &data.Current.CloudCover

			if err := j.repo.UpdateWeatherJob(ctx, *wet); err != nil {
				panic(err)
			}

			logger.Zap().Info("weather_job", zap.String("msg", fmt.Sprintf("data %s sucessfully saved asynchronously", wet.City)))
		}(weather)
	}
}
