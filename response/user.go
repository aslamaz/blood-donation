package response

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterUser struct {
	Id int `json:"id"`
}
