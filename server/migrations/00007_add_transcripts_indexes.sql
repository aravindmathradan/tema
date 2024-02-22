-- +goose Up
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS transcripts_name_idx ON transcripts USING GIN (to_tsvector('english', name));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS transcripts_name_idx;
-- +goose StatementEnd
