-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS projects (
	id bigserial PRIMARY KEY,
	created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
	updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
	name text NOT NULL,
	description text NOT NULL,
	archived boolean NOT NULL DEFAULT FALSE,
	owner_id bigint NOT NULL REFERENCES users ON DELETE CASCADE,
	version integer NOT NULL DEFAULT 1
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS projects;
-- +goose StatementEnd
