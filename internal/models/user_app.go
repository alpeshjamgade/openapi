package models

import "time"

type UserApp struct {
	ID            int       `json:"id" db:"id"`
	Name          string    `json:"name" db:"name"`
	Type          string    `json:"type" db:"type"`
	TradingID     string    `json:"trading_id" db:"trading_id"`
	RedirectURL   string    `json:"redirect_url" db:"redirect_url"`
	PostbackURL   string    `json:"postback_url" db:"postback_url"`
	Description   string    `json:"description" db:"description"`
	AppIconS3Path string    `json:"app_icon_s3_path" db:"app_icon_s3_path"`
	UserID        int       `json:"user_id" db:"user_id"`
	PlanID        int       `json:"plain_id" db:"plain_id"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

type CreateUserAppRequest struct {
	Name          string `json:"name" db:"name" validate:"required"`
	Type          string `json:"type" db:"type" validate:"required"`
	TradingID     string `json:"trading_id" db:"trading_id"`
	RedirectURL   string `json:"redirect_url" db:"redirect_url" validate:"required"`
	PostbackURL   string `json:"postback_url" db:"postback_url" validate:"required"`
	Description   string `json:"description" db:"description" validate:"required"`
	AppIconS3Path string `json:"app_icon_s3_path" db:"app_icon_s3_path"`
	UserID        int    `json:"user_id" db:"user_id" validate:"required"`
	PlanID        int    `json:"plan_id" db:"plan_id" validate:"required"`
}

type UpdateUserAppRequest struct {
	ID            int    `json:"id" db:"id" validate:"required"`
	Name          string `json:"name" db:"name"`
	Type          string `json:"type" db:"type"`
	TradingID     string `json:"trading_id" db:"trading_id"`
	RedirectURL   string `json:"redirect_url" db:"redirect_url"`
	PostbackURL   string `json:"postback_url" db:"postback_url"`
	Description   string `json:"description" db:"description"`
	AppIconS3Path string `json:"app_icon_s3_path" db:"app_icon_s3_path"`
	UserID        int    `json:"user_id" db:"user_id"`
	PlanID        int    `json:"plan_id" db:"plan_id"`
}
