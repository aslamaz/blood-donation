package repository

import (
	"database/sql"
	"fmt"

	"github.com/aslamaz/blood-donation/model"
)

var Db *sql.DB

// getUserByEmail
func GetUserByEmail(email string) (*model.User, error) {
	query := `SELECT id, email, password, address, blood_group, mobile, created_at, updated_at, deleted_at
	FROM user WHERE email=?`

	var u model.User
	if err := Db.QueryRow(query, email).Scan(
		&u.Id,
		&u.Email,
		&u.Password,
		&u.Address,
		&u.BloodGroup,
		&u.Mobile,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.DeletedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}
	return &u, nil
}
