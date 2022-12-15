package usecase

import (
	"fmt"
	"time"

	"github.com/aslamaz/blood-donation/constant"
	"github.com/aslamaz/blood-donation/repository"
	"github.com/aslamaz/blood-donation/request"
	"github.com/aslamaz/blood-donation/response"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(req *request.LoginRequest) (*response.LoginResponse, error) {
	u, err := repository.GetUserByEmail(req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}
	if u == nil || u.Password != req.Password {
		return nil, constant.ErrInvalidCredentials
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, jwt.MapClaims{
		"sub": u.Email,
		"id":  u.Id,
		"exp": time.Now().Add(30 * time.Minute).Unix(),
	})
	hmacSampleSecret := []byte("my_secret_key")
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &response.LoginResponse{
		Token: tokenString,
	}, nil

}
