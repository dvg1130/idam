package models

import "time"

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Login struct {
	HashedPW string
	JWTToken string
}

type RefreshSession struct {
	RefreshToken string    `json:"refresh_token"`
	DeviceID     string    `json:"device_id"`
	ExpiresAt    time.Time `json:"expires_at"`
}
