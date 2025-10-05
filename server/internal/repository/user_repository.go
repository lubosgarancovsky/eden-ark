package repository

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-arc/internal/model"
	"github.com/lubosgarancovsky/eden-arc/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) FindByID(id uuid.UUID) (*model.User, error) {
	var result model.User
	if err := r.db.Model(&model.User{}).Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	var result model.User
	if err := r.db.Model(&model.User{}).Where("username = ?", username).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var result model.User
	if err := r.db.Model(&model.User{}).Where("email = ?", email).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *UserRepository) Insert(user *model.User) (*model.User, error) {
	if err := r.db.Clauses(clause.Returning{}).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) Update(user *model.User) (*model.User, error) {
	result := r.db.Clauses(clause.Returning{}).Where("id = ?", user.ID).Updates(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.ErrNotFound
	}
	return user, nil
}

func (r *UserRepository) Delete(id uuid.UUID) error {
	result := r.db.Model(&model.User{}).Where("id = ?", id).Delete(&model.User{})
	if result.Error != nil {
		return errors.Wrap(errors.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return errors.ErrNotFound
	}
	return nil
}
