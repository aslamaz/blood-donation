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
func GetUserByMobile(mobile string) (*model.User, error) {
	query := `SELECT id, email, password, address, blood_group, mobile, created_at, updated_at, deleted_at
	FROM user WHERE mobile=?`

	var u model.User
	if err := Db.QueryRow(query, mobile).Scan(
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
		return nil, fmt.Errorf("failed to get user by mobile: %w", err)
	}
	return &u, nil
}
func GetUserByBloodGroup(blood_group string) (*model.User, error) {
	query := `SELECT id, email, password, address, blood_group, mobile, created_at, updated_at, deleted_at
	FROM user WHERE blood_group=?`

	var u model.User
	if err := Db.QueryRow(query, blood_group).Scan(
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
		return nil, fmt.Errorf("failed to get user by bloodgroup: %w", err)
	}
	return &u, nil
}

func GetUserById(id int) (*model.User, error) {
	query := `SELECT id, email, password, address, blood_group, mobile, created_at, updated_at, deleted_at
	FROM user WHERE id=?`

	var u model.User
	if err := Db.QueryRow(query, id).Scan(
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
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}
	return &u, nil
}
