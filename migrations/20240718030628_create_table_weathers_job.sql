-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS weathers_job (
    id serial PRIMARY KEY,
    job_id uuid NOT NULL,
    job_spec varchar(20) NOT NULL,
    CONSTRAINT fk_job_id FOREIGN KEY (job_id) REFERENCES weathers(id)
        ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS weathers_job;
-- +goose StatementEnd
