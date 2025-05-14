package models

import "time"

type Permission struct {
	ID        int64     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type UpdatePermissionRequest struct {
	Name    string `json:"name" db:"name" validate:"required"`
	NewName string `json:"new_name" validate:"required"`
}
