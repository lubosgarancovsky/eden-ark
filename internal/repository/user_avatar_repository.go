package repository

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-ark/internal/model"
	"gorm.io/gorm"
)

type UserAvatarRepository struct {
	db *gorm.DB
}

func NewUserAvatarRepository(db *gorm.DB) *UserAvatarRepository {
	return &UserAvatarRepository{db}
}

func (r *UserAvatarRepository) Delete(userID uuid.UUID) error {
	return r.db.
		Model(model.User{}).
		Where("id = ?", userID).
		Updates(map[string]interface{}{
			"avatarMime":    nil,
			"avatarVersion": nil,
		}).
		Error
}
