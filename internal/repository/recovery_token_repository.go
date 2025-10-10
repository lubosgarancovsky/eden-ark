package repository

import (
	"github.com/lubosgarancovsky/eden-ark/internal/model"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RecoveryTokenRepository struct {
	db *gorm.DB
}

func NewRecoveryTokenRepository(db *gorm.DB) *RecoveryTokenRepository {
	return &RecoveryTokenRepository{db}
}

func (r *RecoveryTokenRepository) FindOne(token string) (*model.RecoveryToken, error) {
	var result model.RecoveryToken
	if err := r.db.Model(&model.RecoveryToken{}).Where("token = ?", token).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *RecoveryTokenRepository) Insert(token *model.RecoveryToken) (*model.RecoveryToken, error) {
	if err := r.db.Clauses(clause.Returning{}).Create(token).Error; err != nil {
		return nil, err
	}
	return token, nil
}

func (r *RecoveryTokenRepository) Delete(token string) error {
	result := r.db.Where("token = ?", token).Delete(&model.RecoveryToken{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return api_err.ErrNotFound
	}
	return nil
}
