package model

import "time"

type User struct {
	Id         int        `json:"id"`
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	Password   string     `json:"password"`
	Address    string     `json:"address"`
	BloodGroup string     `json:"bloodgroup"`
	Mobile     string     `json:"mobile"`
	CreatedAt  time.Time  `json:"createdat"`
	UpdatedAt  time.Time  `json:"updatedat"`
	DeletedAt  *time.Time `json:"deletedat"`
}
