-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS weathers (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    city varchar(100) NOT NULL,
    lat decimal(11, 7) NOT NULL,
    lon decimal(11, 7) NOT NULL,
    summary varchar(255),
    temperature real,
    wind_speed real,
    wind_angle numeric,
    wind_dir varchar(3),
    precipitation_total real,
    precipitation_type varchar(15),
    cloud_cover numeric,
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT now(),
    deleted_at timestamptz
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS weathers;
-- +goose StatementEnd
