package model

import "time"

type User struct {
	Id         int
	Email      string
	Password   string
	Address    string
	BloodGroup string
	Mobile     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}
