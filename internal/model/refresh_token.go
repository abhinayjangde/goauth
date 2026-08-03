package model

import "time"

type RefreshToken struct {
	ID        int64
	UserId    int64
	Token     string
	ExpiresAt time.Time
	CreatedAt time.Time
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}
