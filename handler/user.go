package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aslamaz/blood-donation/constant"
	"github.com/aslamaz/blood-donation/model"
	"github.com/aslamaz/blood-donation/request"
	"github.com/aslamaz/blood-donation/response"
	"github.com/aslamaz/blood-donation/shared"
	"github.com/aslamaz/blood-donation/usecase"
	"github.com/go-chi/chi/v5"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	var req request.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		shared.SendJson(w, http.StatusBadRequest, &response.Response{
			Error: err.Error(),
		})
		return
	}

	resp, err := usecase.GenerateToken(&req)
	if err != nil {
		if err == constant.ErrInvalidCredentials {
			shared.SendJson(w, http.StatusUnauthorized, &response.Response{
				Error: "invalid credentials",
			})
		} else {
			fmt.Println(err)
			shared.SendJson(w, http.StatusInternalServerError, &response.Response{
				Error: "internal server errorr",
			})
		}
		return
	}

	shared.SendJson(w, http.StatusOK, &response.Response{Data: resp})
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*model.User)
	shared.SendJson(w, http.StatusOK, &response.Response{
		Data: user,
	})

}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req request.RegisterUser
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		shared.SendJson(w, http.StatusBadRequest, &response.Response{
			Error: err.Error(),
		})
		return
	}
	resp, err := usecase.RegisterUser(&req)

	if err != nil {
		switch err {

		case constant.ErrDuplicateEmail:
			shared.SendJson(w, http.StatusConflict, &response.Response{
				Error: "email already exist",
			})

		case constant.ErrDuplicateMobile:
			shared.SendJson(w, http.StatusConflict, &response.Response{
				Error: "mobile already exist",
			})

		case constant.ErrInvalidBloodGroup:
			shared.SendJson(w, http.StatusBadRequest, &response.Response{
				Error: "invalid bloodgroup",
			})

		default:
			fmt.Println(err)
			shared.SendJson(w, http.StatusInternalServerError, &response.Response{
				Error: "internal server errorr",
			})

		}
		return
	}
	shared.SendJson(w, http.StatusOK, &response.Response{Data: resp})

}
func ChangePassword(w http.ResponseWriter, r *http.Request) {
	var req request.ChangePassword
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		shared.SendJson(w, http.StatusBadRequest, &response.Response{
			Error: err.Error(),
		})
		return
	}

	req.User = r.Context().Value("user").(*model.User)
	err = usecase.ChangePassword(&req)
	if err != nil {
		switch err {
		case constant.ErrInvalidCredentials:
			shared.SendJson(w, http.StatusUnauthorized, &response.Response{
				Error: "invalid credentials",
			})
		default:
			fmt.Println(err)
			shared.SendJson(w, http.StatusInternalServerError, &response.Response{
				Error: "internal server error",
			})
		}
		return
	}

	shared.SendJson(w, http.StatusOK, &response.Response{
		Message: "updated password",
	})
}
func GetMatchingBloodGroupsOfUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*model.User)
	resp, err := usecase.GetMatchingBloodGroups(&request.GetMatchingBloodGroups{
		BloodGroupId: user.BloodGroupId,
	})
	if err != nil {
		shared.SendJson(w, http.StatusInternalServerError, &response.Response{
			Error: "internal sever error",
		})
		return
	}
	shared.SendJson(w, http.StatusOK, &response.Response{
		Data: resp})
}
func GetMatchingBloodGroups(w http.ResponseWriter, r *http.Request) {
	bgId := chi.URLParam(r, "bgid")
	bloodGroupId, err := strconv.ParseInt(bgId, 10, 64)
	if err != nil {
		shared.SendJson(w, http.StatusNotFound, &response.Response{
			Error: "not found",
		})
		return
	}
	resp, err := usecase.GetMatchingBloodGroups(&request.GetMatchingBloodGroups{
		BloodGroupId: int(bloodGroupId),
	})
	if err != nil {
		switch err {
		case constant.ErrInvalidBloodGroup:
			shared.SendJson(w, http.StatusNotFound, &response.Response{
				Error: "not found",
			})
		default:
			shared.SendJson(w, http.StatusInternalServerError, &response.Response{
				Error: "internal sever error",
			})
		}
		return
	}
	shared.SendJson(w, http.StatusOK, &response.Response{
		Data: resp})
}
