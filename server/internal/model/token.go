package model

import (
	"time"

	"github.com/google/uuid"
)

type Token struct {
	ID           uuid.UUID `json:"id"`
	UserID       uuid.UUID `json:"userId"`
	ClientID     uuid.UUID `json:"clientId"`
	AccessToken  string    `json:"accessToken"`
	RefreshToken *string   `json:"refreshToken,omitempty"`
	Scopes       []string  `json:"scopes"`
	ExpiresAt    time.Time `json:"expiresAt"`
	CreatedAt    time.Time `json:"createdAt"`
	Revoked      bool      `json:"revoked"`
}

func (u *Token) TableName() string {
	return "iam_tokens"
}
