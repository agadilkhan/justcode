package dto

type LoginResponse struct {
	ID          int    `json:"id"`
	Login       string `json:"login"`
	AccessToken string `json:"access_token"`
}
