package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aslamaz/blood-donation/repository"
	"github.com/aslamaz/blood-donation/request"
	"github.com/aslamaz/blood-donation/response"
	"github.com/golang-jwt/jwt/v4"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	var req request.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		sendJson(w, http.StatusBadRequest, &response.Response{
			Error: err.Error(),
		})
		return
	}

	u, err := repository.GetUserByEmail(req.Email)
	if err != nil {
		fmt.Println(err)
		sendJson(w, http.StatusInternalServerError, &response.Response{
			Error: "internal server errorr",
		})
		return
	}
	if u == nil {
		sendJson(w, http.StatusUnauthorized, &response.Response{
			Error: "invalid credentials",
		})
		return
	}

	if u.Password == req.Password {

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
			fmt.Println("failed to generate token", err)
			sendJson(w, http.StatusInternalServerError, &response.Response{
				Error: "internal server errorr",
			})
		}
		sendJson(w, http.StatusOK, &response.Response{
			Data: response.LoginResponse{
				Token: tokenString,
			},
		})
		return
	}
	sendJson(w, http.StatusUnauthorized, &response.Response{
		Error: "invalid credentials",
	})
}
