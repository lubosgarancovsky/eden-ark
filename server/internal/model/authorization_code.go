package model

import (
	"time"

	"github.com/google/uuid"
)

type AuthorizationCode struct {
	Code                string    `json:"code"`
	UserID              uuid.UUID `json:"userId"`
	ClientID            uuid.UUID `json:"clientId"`
	RedirectURI         string    `json:"redirectUri"`
	Scopes              []string  `json:"scopes"`
	CodeChallenge       *string   `json:"codeChallenge,omitempty"`
	CodeChallengeMethod *string   `json:"codeChallengeMethod,omitempty"`
	ExpiresAt           time.Time `json:"expiresAt"`
	CreatedAt           time.Time `json:"createdAt"`
}

func (a *AuthorizationCode) TableName() string {
	return "iam_auth_codes"
}
