-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS permissions (
	id bigserial PRIMARY KEY,
	code text NOT NULL
);

CREATE TABLE IF NOT EXISTS users_permissions (
	user_id bigint NOT NULL REFERENCES users ON DELETE CASCADE,
	permission_id bigint NOT NULL REFERENCES permissions ON DELETE CASCADE,
	PRIMARY KEY (user_id, permission_id),
	created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
	updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);

INSERT INTO permissions (code)
VALUES
	('projects:read'),
	('projects:write');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users_permissions;
DROP TABLE IF EXISTS permissions;
-- +goose StatementEnd
