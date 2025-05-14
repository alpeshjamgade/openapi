package models

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Email     string `json:"email"`
	UserID    int    `json:"user_id"`
	AuthToken string `json:"auth_token"`
}
