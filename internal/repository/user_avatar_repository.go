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

func (r *UserAvatarRepository) Change(userID uuid.UUID, avatarMime string, avatarVersion int) error {
	return r.db.
		Model(model.User{}).
		Where("id = ?", userID).
		Updates(map[string]interface{}{
			"avatar_mime":    avatarMime,
			"avatar_version": avatarVersion,
		}).
		Error
}

func (r *UserAvatarRepository) Delete(userID uuid.UUID) error {
	return r.db.
		Model(model.User{}).
		Where("id = ?", userID).
		Updates(map[string]interface{}{
			"avatar_mime":    nil,
			"avatar_version": nil,
		}).
		Error
}
