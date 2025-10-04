package model

import (
	"time"

	"github.com/google/uuid"
)

type ClientRequest struct {
	Name           string   `json:"name"`
	RedirectUris   []string `json:"redirectUris"`
	GrantTypes     []string `json:"grantTypes"`
	IsConfidential bool     `json:"isConfidential"`
}

type Client struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	RedirectUris   []string  `json:"redirectUris"`
	GrantTypes     []string  `json:"grantTypes"`
	IsConfidential bool      `json:"isConfidential"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

func (c *Client) TableName() string {
	return "iam_clients"
}
