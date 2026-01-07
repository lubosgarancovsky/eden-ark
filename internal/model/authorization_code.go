package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type AuthorizationCode struct {
	Code                string         `json:"code"`
	UserID              uuid.UUID      `json:"userId"`
	ClientID            uuid.UUID      `json:"clientId"`
	SessionID           uuid.UUID      `json:"sessionId"`
	RedirectURI         string         `json:"redirectUri"`
	Scopes              pq.StringArray `json:"scopes" gorm:"type:text[]" swaggertype:"array,string"`
	CodeChallenge       string         `json:"codeChallenge,omitempty"`
	CodeChallengeMethod string         `json:"codeChallengeMethod,omitempty"`
	ExpiresAt           time.Time      `json:"expiresAt"`
	CreatedAt           time.Time      `json:"createdAt"`
}

func (a *AuthorizationCode) TableName() string {
	return "iam_auth_codes"
}
