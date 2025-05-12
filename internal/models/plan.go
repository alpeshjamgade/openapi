package models

import "time"

type Plan struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name" validate:"required"`
	Type      string    `json:"type" db:"type" validate:"required"`
	Status    string    `json:"status" db:"status" validate:"required"`
	Amount    int32     `json:"amount" db:"amount" validate:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type CreatePlanRequest struct {
	Name   string `json:"name" db:"name" validate:"required"`
	Type   string `json:"type" db:"type" validate:"required"`
	Status string `json:"status" db:"status" validate:"required"`
	Amount int32  `json:"amount" db:"amount" validate:"required"`
}

type UpdatePlanRequest struct {
	ID     int    `json:"id" db:"id" validate:"required"`
	Name   string `json:"name" db:"name"`
	Type   string `json:"type" db:"type"`
	Status string `json:"status" db:"status"`
	Amount int32  `json:"amount" db:"amount"`
}
