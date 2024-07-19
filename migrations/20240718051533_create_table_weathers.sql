-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS weathers (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    date varchar(20),
    city_id uuid NOT NULL,
    summary varchar(255),
    temperature real,
    wind_speed real,
    wind_angle numeric,
    wind_dir varchar(3),
    precipitation_total real,
    precipitation_type varchar(15),
    cloud_cover real,
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT now(),
    deleted_at timestamptz,
    UNIQUE (date, city_id),
    CONSTRAINT fk_city_id FOREIGN KEY (city_id) REFERENCES cities (id)
        ON UPDATE CASCADE ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS weathers;
-- +goose StatementEnd
