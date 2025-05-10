package models

import "time"

type Plan struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Type      string    `json:"type" db:"type"`
	Status    string    `json:"status" db:"status"`
	Amount    int32     `json:"amount" db:"amount"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
