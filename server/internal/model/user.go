package model

import (
	"time"

	"github.com/google/uuid"
)

type IsAvailableRequest struct {
	Value string `json:"value"`
}

type IsAvailableResponse struct {
	IsAvailable bool `json:"isAvailable"`
}

type CreateUserRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	FirsName string `json:"firsName"`
	LastName string `json:"lastName"`
	IsAdmin  bool   `json:"isAdmin"`
}

type UpdateUserRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	IsAdmin   bool   `json:"isAdmin"`
}

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Role         string    `json:"role"`
	IsActive     bool      `json:"isActive"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	DeletedAt    time.Time `json:"deletedAt"`
}

func (u *User) TableName() string {
	return "iam_users"
}
