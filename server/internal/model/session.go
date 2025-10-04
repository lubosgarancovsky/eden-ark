package model

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID           uuid.UUID `json:"id"`
	UserID       uuid.UUID `json:"userId"`
	SessionToken string    `json:"sessionToken"`
	IPAddress    *string   `json:"ipAddress,omitempty"`
	UserAgent    *string   `json:"userAgent,omitempty"`
	ExpiresAt    time.Time `json:"expiresAt"`
	CreatedAt    time.Time `json:"createdAt"`
}

func (u *Session) TableName() string {
	return "iam_sessions"
}
