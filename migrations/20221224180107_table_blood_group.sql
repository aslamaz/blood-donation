-- +goose Up
CREATE TABLE blood_group(
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(3) NOT NULL
);
-- +goose StatementBegin
INSERT INTO blood_group (id, name) VALUES
(1,'A+'),
(2,'B+'),
(3,'AB+'),
(4,'O+'),
(5,'A-'),
(6,'B-'),
(7,'AB-'),
(8,'O-');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE blood_group;
-- +goose StatementEnd
