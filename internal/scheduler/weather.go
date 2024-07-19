package scheduler

import (
	"context"
	"fmt"
	"i3/internal/api"
	"i3/internal/entity"
	"i3/internal/meteo"
	"i3/pkg/logger"
	"strconv"
	"time"

	"go.uber.org/zap"
)

type WeatherJob struct {
	meteo meteo.Meteo
	repo  api.Repository
}

func NewWeatherJob(meteo meteo.Meteo, repo api.Repository) *WeatherJob {
	return &WeatherJob{
		meteo: meteo,
		repo:  repo,
	}
}

func (j *WeatherJob) Run() {
	ctx := context.Background()

	cities, err := j.repo.FetchCities(ctx)
	if err != nil {
		logger.Zap().Error("weather_job", zap.Error(err))
	}

	if len(cities) == 0 {
		logger.Zap().Info("weather_job", zap.String("msg", "no job found"))
		return
	}

	for _, city := range cities {
		go func(c *entity.Cities) {
			lat := strconv.FormatFloat(float64(c.Lat), 'f', 11, 32)
			lon := strconv.FormatFloat(float64(c.Lon), 'f', 11, 32)

			data, err := j.meteo.FetchLocationWeather(c.ID.String(), lat, lon)
			if err != nil {
				logger.Zap().Error("weather_job", zap.Error(err))
			}

			// TODO: create function for converting object
			var weather entity.Weathers
			weather.Date = time.Now().Format(time.DateOnly)
			weather.CityID = c.ID.String()
			weather.Summary = &data.Current.Summary
			weather.Temperature = &data.Current.Temperature
			weather.WindSpeed = &data.Current.Wind.Speed
			weather.WindAngle = &data.Current.Wind.Angle
			weather.WindDir = &data.Current.Wind.Dir
			weather.PrecipitationTotal = &data.Current.Precipitation.Total
			weather.PrecipitationType = &data.Current.Precipitation.Type
			weather.CloudCover = &data.Current.CloudCover

			if err := j.repo.InsertWeatherJob(ctx, weather); err != nil {
				logger.Zap().Error("weather_job", zap.Error(err))
			}

			logger.Zap().Info("weather_job", zap.String("msg", fmt.Sprintf("data %s sucessfully saved asynchronously", c.Name)))
		}(city)
	}
}
