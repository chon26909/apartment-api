package models

// type Common struct {
// 	Code    int
// 	Message string
// }

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    LoginResponseData `json:"data"`
}

type LoginResponseData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RequestTokenRequest struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
