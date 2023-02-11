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

func GetDonorBloodGroups(recipientBloodGroupId int) ([]string, error) {
	query := `SELECT bg.name 
	FROM matching_blood_group mbg 
	INNER JOIN blood_group bg ON mbg.donor_blood_group =bg.id 
	WHERE recipient_blood_group=?`
	rows, err := Db.Query(query, recipientBloodGroupId)
	if err != nil {
		return nil, fmt.Errorf("failed to get donor bloodgroup by id:%w", err)
	}
	defer rows.Close()
	var donorBloodGroups []string
	var bloodGroup string
	for rows.Next() {
		if err := rows.Scan(&bloodGroup); err != nil {
			return nil, fmt.Errorf("failed to scan donor bloodgroup by id:%w", err)
		}
		donorBloodGroups = append(donorBloodGroups, bloodGroup)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate donor bloodgroup by id:%w", err)
	}
	return donorBloodGroups, nil
}

func GetRecipientBloodGroups(donorBloodGroupId int) ([]string, error) {
	query := `SELECT bg.name 
	FROM matching_blood_group mbg 
	INNER JOIN blood_group bg ON mbg.recipient_blood_group =bg.id 
	WHERE donor_blood_group =?`
	rows, err := Db.Query(query, donorBloodGroupId)
	if err != nil {
		return nil, fmt.Errorf("failed to scan recipient bloodgroup by id:%w", err)
	}
	defer rows.Close()
	var recipientBloodGroups []string
	var bloodGroup string
	for rows.Next() {
		if err := rows.Scan(&bloodGroup); err != nil {
			return nil, fmt.Errorf("failed to scan donor bloodgroup by id:%w", err)
		}
		recipientBloodGroups = append(recipientBloodGroups, bloodGroup)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate recipient bloodgroup by id:%w", err)
	}
	return recipientBloodGroups, nil
}
