package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/aslamaz/blood-donation/constant"
	"github.com/aslamaz/blood-donation/repository"
	"github.com/aslamaz/blood-donation/request"
	"github.com/aslamaz/blood-donation/response"
	"github.com/aslamaz/blood-donation/usecase"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
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

	resp, err := usecase.GenerateToken(&req)
	if err != nil {
		if err == constant.ErrInvalidCredentials {
			sendJson(w, http.StatusUnauthorized, &response.Response{
				Error: "invalid credentials",
			})
		} else {
			fmt.Println(err)
			sendJson(w, http.StatusInternalServerError, &response.Response{
				Error: "internal server errorr",
			})
		}
		return
	}

	sendJson(w, http.StatusOK, &response.Response{Data: resp})
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")

	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			sendJson(w, http.StatusForbidden, &response.Response{
				Error: "invalid session",
			})

		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(constant.JwtSigningKey), nil
	})

	if err != nil {
		sendJson(w, http.StatusForbidden, &response.Response{
			Error: "invalid session",
		})
		return
	}
	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		idRaw := claims["id"].(float64) //Type assertion interface{} -> type (float64)
		id := int(idRaw)
		user, err := repository.GetUserById(id)
		if err != nil {
			fmt.Println(err)
			sendJson(w, http.StatusInternalServerError, &response.Response{
				Error: "internal server errorr",
			})
			return
		}
		if user == nil {
			sendJson(w, http.StatusForbidden, &response.Response{
				Error: "invalid session",
			})
			return
		}
		sendJson(w, http.StatusOK, &response.Response{
			Data: user,
		})
	} else {
		sendJson(w, http.StatusForbidden, &response.Response{
			Error: "invalid session",
		})
	}

}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req request.RegisterUser
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		sendJson(w, http.StatusBadRequest, &response.Response{
			Error: err.Error(),
		})
		return
	}
	resp, err := usecase.RegisterUser(&req)

	if err != nil {
		switch err {

		case constant.ErrDuplicateEmail:
			sendJson(w, http.StatusConflict, &response.Response{
				Error: "email already exist",
			})

		case constant.ErrDuplicateMobile:
			sendJson(w, http.StatusConflict, &response.Response{
				Error: "mobile already exist",
			})

		case constant.ErrInvalidBloodGroup:
			sendJson(w, http.StatusBadRequest, &response.Response{
				Error: "invalid bloodgroup",
			})

		default:
			fmt.Println(err)
			sendJson(w, http.StatusInternalServerError, &response.Response{
				Error: "internal server errorr",
			})

		}
		return
	}
	sendJson(w, http.StatusOK, &response.Response{Data: resp})

}
func ChangePassword(w http.ResponseWriter, r *http.Request) {
	var req request.ChangePassword
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		sendJson(w, http.StatusBadRequest, &response.Response{
			Error: err.Error(),
		})
		return
	}

	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")

	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			sendJson(w, http.StatusForbidden, &response.Response{
				Error: "invalid session",
			})

		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(constant.JwtSigningKey), nil
	})

	if err != nil {
		sendJson(w, http.StatusForbidden, &response.Response{
			Error: "invalid session",
		})
		return
	}
	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		idRaw := claims["id"].(float64) //Type assertion interface{} -> type (float64)
		id := int(idRaw)
		user, err := repository.GetUserById(id)
		if err != nil {
			fmt.Println(err)
			sendJson(w, http.StatusInternalServerError, &response.Response{
				Error: "internal server errorr",
			})
			return
		}
		if user == nil {
			sendJson(w, http.StatusForbidden, &response.Response{
				Error: "invalid session",
			})
			return
		}

		if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword)); err != nil {
			sendJson(w, http.StatusForbidden, &response.Response{
				Error: "invalid credentials",
			})
			return
		}
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), 10)
		if err != nil {
			fmt.Println(err)
			sendJson(w, http.StatusInternalServerError, &response.Response{
				Error: "internal server errorr",
			})
			return
		}
		if err = repository.UpdateUserPassword(id, string(passwordHash)); err != nil {
			fmt.Println(err)
			sendJson(w, http.StatusInternalServerError, &response.Response{
				Error: "internal server errorr",
			})
			return
		}
		sendJson(w, http.StatusOK, &response.Response{
			Message: "updated password",
		})
	} else {
		sendJson(w, http.StatusForbidden, &response.Response{
			Error: "invalid session",
		})
	}
}
