package repository

import (
	"github.com/lubosgarancovsky/eden-ark/internal/model"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AuthCodeRepository struct {
	db *gorm.DB
}

func NewAuthCodeRepository(db *gorm.DB) *AuthCodeRepository {
	return &AuthCodeRepository{db}
}

func (r *AuthCodeRepository) FindByCode(code string) (*model.AuthorizationCode, error) {
	var result model.AuthorizationCode
	if err := r.db.Model(&model.AuthorizationCode{}).Where("code = ?", code).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *AuthCodeRepository) Insert(code *model.AuthorizationCode) (*model.AuthorizationCode, error) {
	if err := r.db.Clauses(clause.Returning{}).Create(code).Error; err != nil {
		return nil, err
	}
	return code, nil
}

func (r *AuthCodeRepository) Delete(code string) error {
	result := r.db.Model(&model.AuthorizationCode{}).Where("code = ?", code).Delete(&model.AuthorizationCode{})
	if result.Error != nil {
		return api_err.Wrap(api_err.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return api_err.ErrNotFound
	}
	return nil
}
