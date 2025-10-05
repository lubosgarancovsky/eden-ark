package repository

import (
	"github.com/lubosgarancovsky/eden-arc/internal/model"
	"github.com/lubosgarancovsky/eden-arc/pkg/errors"
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
		return nil, errors.Wrap(errors.ErrInternalServer, err)
	}
	return &result, nil
}

func (r *AuthCodeRepository) Insert(code *model.AuthorizationCode) (*model.AuthorizationCode, error) {
	if err := r.db.Clauses(clause.Returning{}).Create(code).Error; err != nil {
		return nil, errors.Wrap(errors.ErrInternalServer, err)
	}
	return code, nil
}

func (r *AuthCodeRepository) Delete(code string) error {
	result := r.db.Model(&model.AuthorizationCode{}).Where("code = ?", code).Delete(&model.AuthorizationCode{})
	if result.Error != nil {
		return errors.Wrap(errors.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return errors.ErrNotFound
	}
	return nil
}
