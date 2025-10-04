package model

import (
	"time"

	"github.com/google/uuid"
)

type RecoveryToken struct {
	ID           uuid.UUID `json:"id"`
	UserID       uuid.UUID `json:"userId"`
	Token        string    `json:"token"`
	RecoveryType string    `json:"recoveryType"` // "password" or "email"
	Metadata     *string   `json:"metadata,omitempty"`
	ExpiresAt    time.Time `json:"expiresAt"`
	Used         bool      `json:"used"`
	CreatedAt    time.Time `json:"createdAt"`
}

func (u *RecoveryToken) TableName() string {
	return "iam_recovery_tokens"
}
