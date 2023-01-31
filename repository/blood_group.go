package repository

import (
	"database/sql"
	"fmt"

	"github.com/aslamaz/blood-donation/model"
)

func GetBloodGroupById(id int) (*model.BloodGroup, error) {
	query := `SELECT id,name FROM blood_group WHERE id=?`

	var u model.BloodGroup
	if err := Db.QueryRow(query, id).Scan(
		&u.Id,
		&u.Name,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get bloodgroup by id: %w", err)
	}
	return &u, nil
}
