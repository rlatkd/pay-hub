package models

import (
	"database/sql"
	"time"
)

type User struct {
	UserID       string         `json:"user_id"`
	Email        string         `json:"email"`
	PasswordHash string         `json:"-"`
	Name         string         `json:"name"`
	Status       string         `json:"status"`
	RoleID       string         `json:"role_id"`
	LastLoginAt  sql.NullTime   `json:"last_login_at"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

type RefreshToken struct {
	TokenID         string    `json:"token_id"`
	UserID          string    `json:"user_id"`
	RotationCounter int       `json:"rotation_counter"`
	IssuedAt        time.Time `json:"issued_at"`
	ExpiresAt       time.Time `json:"expires_at"`
	DeviceInfo      string    `json:"device_info"`
}
