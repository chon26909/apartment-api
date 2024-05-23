package models

type ProfileResponse struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Data    Profile `json:"data"`
}

type Profile struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
