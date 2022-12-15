package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aslamaz/blood-donation/constant"
	"github.com/aslamaz/blood-donation/request"
	"github.com/aslamaz/blood-donation/response"
	"github.com/aslamaz/blood-donation/usecase"
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

	sendJson(w, http.StatusUnauthorized, &response.Response{Data: resp})
}
