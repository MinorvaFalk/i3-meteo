package apirepository

import (
	"context"
	"i3/internal/api"
	"i3/internal/entity"
	"i3/internal/model"
	"i3/pkg/datasource"

	"github.com/jackc/pgx/v5"
)

type repository struct {
	db datasource.Postgres
}

func New(db datasource.Postgres) api.Repository {
	return &repository{db: db}
}

func (s *repository) InsertCity(ctx context.Context, req model.CityRequest) (data model.ScheduleLocationData, err error) {
	query := `INSERT INTO cities (
		name,
		lat, 
		lon
	) VALUES ($1, $2, $3) RETURNING id, name, lat, lon`

	err = s.db.QueryRow(ctx, query,
		req.Name,
		req.Lat,
		req.Lon,
	).Scan(&data.ID, &data.City, &data.Lat, &data.Lon)
	if err != nil {
		return
	}

	return
}

func (s *repository) InsertWeatherJob(ctx context.Context, req entity.Weathers) (err error) {
	query := `INSERT INTO weathers (
		date,
		city_id,
		summary,
		temperature,
		wind_speed,
		wind_angle,
		precipitation_total,
		precipitation_type,
		cloud_cover
	) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		ON CONFLICT (date, city_id) 
		DO UPDATE SET 
			summary = $3,
			temperature = $4,
			wind_speed = $5,
			wind_angle = $6,
			precipitation_total = $7,
			precipitation_type = $8,
			cloud_cover = $9,
			updated_at = now()
	`

	_, err = s.db.Exec(ctx, query,
		req.Date,
		req.CityID,
		req.Summary,
		req.Temperature,
		req.WindSpeed,
		req.WindAngle,
		req.PrecipitationTotal,
		req.PrecipitationType,
		req.CloudCover,
	)
	if err != nil {
		return
	}

	return
}

func (s *repository) FetchCities(ctx context.Context) ([]*entity.Cities, error) {
	query := `SELECT id, name, lat, lon, 
	created_at, updated_at, deleted_at 
	FROM cities`

	rows, _ := s.db.Query(ctx, query)
	data, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByPos[entity.Cities])
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *repository) FetchWeatherData(ctx context.Context) ([]*model.WeatherData, error) {
	query := `SELECT c.id, c.name, c.lat, c.lon,
	date, summary, temperature, 
	wind_speed, wind_angle, wind_dir,
	precipitation_total, precipitation_type, cloud_cover
	FROM cities c LEFT JOIN weathers w ON c.id = w.city_id
	`

	rows, _ := s.db.Query(ctx, query)
	data, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByPos[model.WeatherData])
	if err != nil {
		return nil, err
	}

	return data, nil
}
