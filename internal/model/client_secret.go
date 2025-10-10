package model

import (
	"time"

	"github.com/google/uuid"
)

type ClientSecretRequest struct {
	ExpiresAt time.Time `json:"expiresAt"`
}

type ClientSecret struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	ClientSecret string    `json:"clientSecret"`
	ClientID     uuid.UUID `json:"clientId"`
	ExpiresAt    time.Time `json:"expiresAt"`
	CreatedAt    time.Time `json:"createdAt"`
}

func (c *ClientSecret) TableName() string {
	return "iam_client_secrets"
}
