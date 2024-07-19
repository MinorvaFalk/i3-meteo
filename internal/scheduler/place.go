package scheduler

import (
	"context"
	"fmt"
	"i3/internal/api"
	"i3/internal/entity"
	"i3/internal/meteo"
	"i3/pkg/logger"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

type PlaceJob struct {
	meteo meteo.Meteo
	repo  api.Repository
}

func NewPlaceJob(meteo meteo.Meteo, repo api.Repository) *PlaceJob {
	return &PlaceJob{
		meteo: meteo,
		repo:  repo,
	}
}

func (j *PlaceJob) Run() {
	ctx := context.Background()

	cities, err := j.repo.FetchCities(ctx)
	if err != nil {
		logger.Zap().Error("place_job", zap.Error(err))
	}

	if len(cities) == 0 {
		logger.Zap().Info("place_job", zap.String("msg", "no job found"))
		return
	}

	for _, city := range cities {
		go func(c *entity.Cities) {
			places, err := j.meteo.FetchPlacesPrefix(c.ID.String(), c.Name)
			if err != nil {
				logger.Zap().Error("place_job", zap.Error(err))
			}

			lat := places[0].Lat
			lon := places[0].Lon

			if strings.HasSuffix(lat, "S") {
				lat = "-" + strings.TrimSuffix(lat, "S")
			} else {
				lat = strings.TrimSuffix(lat, "N")
			}

			if strings.HasSuffix(lon, "W") {
				lon = "-" + strings.TrimSuffix(lon, "W")
			} else {
				lon = strings.TrimSuffix(lon, "E")
			}

			latF, _ := strconv.ParseFloat(lat, 32)
			lonF, _ := strconv.ParseFloat(lon, 32)

			// Check if latitude and longitude is the same
			if float32(latF) == c.Lat && float32(lonF) == c.Lon {
				return
			}

			if err := j.repo.UpdateCityCoordinate(ctx, c.ID.String(), float32(latF), float32(lonF)); err != nil {
				logger.Zap().Error("place_job", zap.Error(err))
			}

			logger.Zap().Info("places_job", zap.String("msg", fmt.Sprintf("city %s sucessfully updated asynchronously", c.Name)))
		}(city)
	}
}
