package request

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUser struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Mobile     string `json:"mobile"`
	Address    string `json:"address"`
	Bloodgroup string `json:"bloodgroup"`
	Password   string `json:"password"`
}
