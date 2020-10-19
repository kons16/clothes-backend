
-- +migrate Up
ALTER TABLE clothes ADD in_shop BIGINT;

-- +migrate Down
ALTER TABLE clothes DROP clothes;
