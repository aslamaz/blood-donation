-- +goose Up
-- +goose StatementBegin
ALTER TABLE user
ADD name VARCHAR(100) NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE user
DROP COLUMN name;
-- +goose StatementEnd
