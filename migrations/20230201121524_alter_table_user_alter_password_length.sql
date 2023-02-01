-- +goose Up
-- +goose StatementBegin
ALTER TABLE user
MODIFY password VARCHAR(200) NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE user
MODIFY password VARCHAR(10) NOT NULL;
-- +goose StatementEnd
