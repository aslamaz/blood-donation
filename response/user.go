package response

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterUser struct {
	Id int `json:"id"`
}

type GetMatchingBloodGroups struct {
	Recieves   []string `json:"recieves"`
	Gives      []string `json:"gives"`
	BloodGroup string   `json:"bloodGroup"`
}
type GetBloodGroups struct {
	BloodGroups []string `json:"bloodGroup"`
}
