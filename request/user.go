package request

import "github.com/aslamaz/blood-donation/model"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUser struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Mobile       string `json:"mobile"`
	Address      string `json:"address"`
	BloodGroupId int    `json:"bloodGroupId"`
	Password     string `json:"password"`
}

type ChangePassword struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
	User        *model.User
}

type GetMatchingBloodGroups struct {
	BloodGroupId int
}
