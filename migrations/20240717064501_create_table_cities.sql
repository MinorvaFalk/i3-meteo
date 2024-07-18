-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS cities (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    name varchar(100) UNIQUE NOT NULL,
    lat decimal(11,7) NOT NULL,
    lon decimal(11,7) NOT NULL,
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT now(),
    deleted_at timestamptz
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS city;
-- +goose StatementEnd
