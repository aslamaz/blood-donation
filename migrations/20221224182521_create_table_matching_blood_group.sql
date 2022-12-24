-- +goose Up
CREATE TABLE matching_blood_group(
    id INT PRIMARY KEY AUTO_INCREMENT,
    recipient_blood_group INT NOT NULL,
    donor_blood_group INT NOT NULL,
    CONSTRAINT fk_mbg_rbg_bg_id FOREIGN KEY (recipient_blood_group) REFERENCES blood_group(id),
    CONSTRAINT fk_mbg_dbg_bg_id FOREIGN KEY (donor_blood_group) REFERENCES blood_group(id),
    CONSTRAINT uk_mbg_rbg_dbg UNIQUE(recipient_blood_group,donor_blood_group)
);
-- +goose StatementBegin
INSERT INTO matching_blood_group (recipient_blood_group,donor_blood_group) VALUES
(1,1),
(1,5),
(1,4),
(1,8),
(2,2),
(2,6),
(2,4),
(2,8),
(3,1),
(3,2),
(3,3),
(3,4),
(3,5),
(3,6),
(3,7),
(3,8),
(4,4),
(4,8),
(5,5),
(5,8),
(6,6),
(6,8),
(7,7),
(7,5),
(7,6),
(7,8),
(8,8);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE matching_blood_group;
-- +goose StatementEnd
