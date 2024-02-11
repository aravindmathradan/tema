-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS valid_access_levels (
	access_level text PRIMARY KEY
);

INSERT INTO valid_access_levels (access_level)
VALUES
	('OWNER'),
	('EDITOR'),
	('VIEWER');

CREATE TABLE IF NOT EXISTS projects_users (
	user_id bigint NOT NULL REFERENCES users ON DELETE CASCADE,
	project_id bigint NOT NULL REFERENCES projects ON DELETE CASCADE,
	PRIMARY KEY (user_id, project_id),
	access_level text NOT NULL REFERENCES valid_access_levels(access_level) ON UPDATE CASCADE ON DELETE CASCADE,
	created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
	updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
	version integer NOT NULL DEFAULT 1
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS projects_users;
DROP TABLE IF EXISTS valid_access_levels;
-- +goose StatementEnd
