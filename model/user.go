package model

import "time"

type User struct {
	Id           int        `json:"id"`
	Name         string     `json:"name"`
	Email        string     `json:"email"`
	Password     string     `json:"password"`
	Address      string     `json:"address"`
	BloodGroupId int        `json:"bloodGroupId"`
	Mobile       string     `json:"mobile"`
	CreatedAt    time.Time  `json:"createdat"`
	UpdatedAt    time.Time  `json:"updatedat"`
	DeletedAt    *time.Time `json:"deletedat"`
}
