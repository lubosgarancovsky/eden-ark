package model

import (
	"time"

	"github.com/google/uuid"
)

type ClientSecretRequest struct {
	ClientID  uuid.UUID `json:"clientId"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type ClientSecret struct {
	ID           uuid.UUID `json:"id"`
	ClientSecret string    `json:"clientSecret"`
	ClientID     uuid.UUID `json:"clientId"`
	ExpiresAt    time.Time `json:"expiresAt"`
	CreatedAt    time.Time `json:"createdAt"`
}

func (c *ClientSecret) TableName() string {
	return "iam_client_secrets"
}
