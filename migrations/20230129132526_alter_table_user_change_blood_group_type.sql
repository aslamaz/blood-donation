-- +goose Up
ALTER TABLE user
ADD blood_group_id INT NOT NULL;

UPDATE `user` set blood_group_id=(SELECT id from blood_group WHERE name=user.blood_group);

ALTER TABLE user
DROP COLUMN blood_group;

ALTER TABLE user ADD CONSTRAINT fk_user_blood_group_id_blood_group_id FOREIGN KEY (blood_group_id) REFERENCES blood_group(id);
-- +goose StatementBegin
-- +goose StatementEnd

-- +goose Down
ALTER TABLE user
ADD blood_group VARCHAR(10) NOT NULL;

UPDATE `user` set blood_group=(SELECT name from blood_group WHERE id=user.blood_group_id);

ALTER TABLE user
DROP CONSTRAINT fk_user_blood_group_id_blood_group_id;

ALTER TABLE user
DROP COLUMN blood_group_id;
-- +goose StatementBegin

-- +goose StatementEnd
