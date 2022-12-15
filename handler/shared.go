package handler

import (
	"encoding/json"
	"net/http"

	"github.com/aslamaz/blood-donation/response"
)

func sendJson(w http.ResponseWriter, statusCode int, response *response.Response) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)

}
