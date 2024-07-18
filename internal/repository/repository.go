package repository

import (
	"context"
	"fmt"
	"i3/internal/entity"
	"i3/internal/model"
	"i3/pkg/datasource"
	"time"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	db datasource.Postgres
}

func New(db datasource.Postgres) *Repository {
	return &Repository{db: db}
}

func (s *Repository) InsertWeatherJob(ctx context.Context, req model.WeatherRequest) (id, city, lat, lon string, err error) {
	query := `INSERT INTO weathers (
		city,
		lat,
		lon
	) VALUES ($1, $2, $3) RETURNING id, city, lat, lon;
	`

	err = s.db.QueryRow(ctx, query,
		req.Name,
		req.Lat,
		req.Lon,
	).Scan(&id, &city, &lat, &lon)
	if err != nil {
		return
	}

	return
}

func (s *Repository) UpdateWeatherJob(ctx context.Context, data entity.Weathers) error {
	query := `UPDATE "weathers" SET
		summary = $2,
		temperature = $3,
		wind_speed = $4,
		wind_angle = $5,
		wind_dir = $6,
		precipitation_total = $7,
		precipitation_type = $8,
		cloud_cover = $9,
		updated_at = $10
	WHERE id = $1
	`

	cmd, err := s.db.Exec(ctx, query,
		data.ID,
		data.Summary,
		data.Temperature,
		data.WindSpeed,
		data.WindAngle,
		data.WindDir,
		data.PrecipitationTotal,
		data.PrecipitationType,
		data.CloudCover,
		time.Now(),
	)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("failed to insert data")
	}

	return nil
}

func (s *Repository) FetchWeatherJobs(ctx context.Context) ([]*entity.Weathers, error) {
	query := `SELECT 
		id, city, lat, lon, summary, temperature, 
		wind_speed, wind_angle, wind_dir, 
		precipitation_total, precipitation_type, cloud_cover, 
		created_at, updated_at, deleted_at
		FROM weathers`

	rows, _ := s.db.Query(ctx, query)
	data, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByPos[entity.Weathers])
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Repository) FetchWeatherJob(ctx context.Context, id string) (*entity.Weathers, error) {
	var data entity.Weathers
	query := `SELECT 
		id, city, lat, lon, summary, temperature, 
		wind_speed, wind_angle, wind_dir, 
		precipitation_total, precipitation_type, cloud_cover, 
		created_at, updated_at, deleted_at
		FROM weathers
		WHERE id = $1`

	if err := s.db.QueryRow(ctx, query, id).Scan(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *Repository) CheckJobAvailable(ctx context.Context) (bool, error) {
	query := `SELECT count(1) FROM weathers`

	cmd, err := s.db.Exec(ctx, query)
	if err != nil {
		return false, err
	}

	return cmd.RowsAffected() != 0, nil
}
