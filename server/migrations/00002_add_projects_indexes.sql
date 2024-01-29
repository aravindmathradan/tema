-- +goose Up
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS projects_name_idx ON projects USING GIN (to_tsvector('english', name));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS projects_name_idx;
-- +goose StatementEnd
