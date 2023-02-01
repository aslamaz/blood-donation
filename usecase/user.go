package usecase

import (
	"fmt"
	"time"

	"github.com/aslamaz/blood-donation/constant"
	"github.com/aslamaz/blood-donation/model"
	"github.com/aslamaz/blood-donation/repository"
	"github.com/aslamaz/blood-donation/request"
	"github.com/aslamaz/blood-donation/response"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func GenerateToken(req *request.LoginRequest) (*response.LoginResponse, error) {
	u, err := repository.GetUserByEmail(req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}
	if u == nil {
		return nil, constant.ErrInvalidCredentials
	}
	if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		return nil, constant.ErrInvalidCredentials
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, jwt.MapClaims{
		"sub": u.Email,
		"id":  u.Id,
		"exp": time.Now().Add(30 * time.Minute).Unix(),
	})
	hmacSampleSecret := []byte(constant.JwtSigningKey)
	// Sign and get the complete encoded token as a string using the secret9
	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &response.LoginResponse{
		Token: tokenString,
	}, nil

}

func RegisterUser(req *request.RegisterUser) (*response.RegisterUser, error) {
	existingUser, err := repository.GetUserByEmail(req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}
	if existingUser != nil {
		return nil, constant.ErrDuplicateEmail

	}
	existingUser, err = repository.GetUserByMobile(req.Mobile)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by mobile: %w", err)
	}
	if existingUser != nil {
		return nil, constant.ErrDuplicateMobile
	}
	bloodGroup, err := repository.GetBloodGroupById(req.BloodGroupId)
	if err != nil {
		return nil, fmt.Errorf("failed to get bloodgroup by id: %w", err)
	}

	if bloodGroup == nil {
		return nil, constant.ErrInvalidBloodGroup
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return nil, fmt.Errorf("failed to generate hashed password:%w", err)
	}

	var user model.User
	user.Name = req.Name
	user.Email = req.Email
	user.BloodGroupId = req.BloodGroupId
	user.Address = req.Address
	user.Password = string(passwordHash)
	user.Mobile = req.Mobile

	id, err := repository.InsertUser(user)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user:%w", err)
	}

	return &response.RegisterUser{
		Id: id,
	}, nil
}
