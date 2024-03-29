-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS transcripts (
	id bigserial PRIMARY KEY,
	created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
	updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
	title text NOT NULL,
	description text NOT NULL,
	content text NOT NULL,
	archived boolean NOT NULL DEFAULT FALSE,
	project_id bigint NOT NULL REFERENCES projects ON DELETE CASCADE,
	version integer NOT NULL DEFAULT 1
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS transcripts;
-- +goose StatementEnd
