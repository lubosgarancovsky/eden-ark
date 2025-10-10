package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type ClientRequest struct {
	Name           string   `json:"name"`
	RedirectUris   []string `gorm:"type:text[]" json:"redirectUris"`
	GrantTypes     []string `gorm:"type:text[]" json:"grantTypes"`
	IsConfidential bool     `json:"isConfidential"`
}

type Client struct {
	ID             uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name           string         `json:"name"`
	RedirectUris   pq.StringArray `gorm:"type:text[]" json:"redirectUris"  swaggertype:"array,string"`
	GrantTypes     pq.StringArray `gorm:"type:text[]" json:"grantTypes"  swaggertype:"array,string"`
	IsConfidential bool           `json:"isConfidential"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
}

func (c *Client) TableName() string {
	return "iam_clients"
}
