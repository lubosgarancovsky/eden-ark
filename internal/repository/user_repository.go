package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-ark/internal/model"
	"github.com/lubosgarancovsky/eden-ark/pkg/helpers"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"github.com/lubosgarancovsky/go-kit/list"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) FindAll(lq *list.ListingQuery) ([]model.User, int64, error) {
	query := r.db.Model(&model.User{})
	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	items, total, err := helpers.List[model.User](query, lq)
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
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
		return nil, api_err.ErrNotFound
	}
	return user, nil
}

func (r *UserRepository) MarkAsDeleted(id uuid.UUID) error {
	result := r.db.Model(&model.User{}).Where("id = ?", id).Updates(map[string]time.Time{
		"deleted_at": time.Now(),
	})

	if result.Error != nil {
		return api_err.Wrap(api_err.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return api_err.ErrNotFound
	}
	return nil
}

func (r *UserRepository) Delete(id uuid.UUID) error {
	result := r.db.Where("id = ?", id).Delete(&model.User{})
	if result.Error != nil {
		return api_err.Wrap(api_err.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return api_err.ErrNotFound
	}
	return nil
}
