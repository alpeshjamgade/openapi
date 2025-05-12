package models

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Email     string `json:"email"`
	AuthToken string `json:"auth_token"`
}
