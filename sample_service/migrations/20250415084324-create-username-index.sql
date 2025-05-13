
-- +migrate Up
CREATE INDEX username_idx ON users(username);
-- +migrate Down
DROP INDEX username_idx;