package models

import "time"

type User struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name" validate:"required"`
	Email     string    `json:"email" db:"email" validate:"required,email"`
	Phone     string    `json:"phone" db:"phone" validate:"required"`
	Website   string    `json:"website" db:"website"`
	About     string    `json:"about" db:"about"`
	State     string    `json:"state" db:"state" validate:"required"`
	PartnerID string    `json:"partner_id" db:"partner_id"`
	Password  string    `json:"password" db:"password" validate:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
