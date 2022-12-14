package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v4"
)

var db *sql.DB
var err error

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/ping", handlePing)
	r.Post("/login", handleLogin)

	db, err = sql.Open("mysql", "root:root@/blood_donation?parseTime=true")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	err = http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}

}
func handlePing(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func sendJson(w http.ResponseWriter, statusCode int, response *Response) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)

}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		sendJson(w, http.StatusBadRequest, &Response{
			Error: err.Error(),
		})
		return
	}

	u, err := getUserByEmail(req.Email)
	if err != nil {
		fmt.Println(err)
		sendJson(w, http.StatusInternalServerError, &Response{
			Error: "internal server errorr",
		})
		return
	}
	if u == nil {
		sendJson(w, http.StatusUnauthorized, &Response{
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
			sendJson(w, http.StatusInternalServerError, &Response{
				Error: "internal server errorr",
			})
		}
		sendJson(w, http.StatusOK, &Response{
			Data: LoginResponse{
				Token: tokenString,
			},
		})
		return
	}
	sendJson(w, http.StatusUnauthorized, &Response{
		Error: "invalid credentials",
	})
}

type user struct {
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

// getUserByEmail
func getUserByEmail(email string) (*user, error) {
	query := `SELECT id, email, password, address, blood_group, mobile, created_at, updated_at, deleted_at
	FROM user WHERE email=?`

	var u user
	if err := db.QueryRow(query, email).Scan(
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
