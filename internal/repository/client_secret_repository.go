package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-ark/internal/model"
	"github.com/lubosgarancovsky/eden-ark/pkg/helpers"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"github.com/lubosgarancovsky/go-kit/list"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ClientSecretRepository struct {
	db *gorm.DB
}

func NewClientSecretRepository(db *gorm.DB) *ClientSecretRepository {
	return &ClientSecretRepository{db}
}

func (r *ClientSecretRepository) FindAll(clientID uuid.UUID, lq *list.ListingQuery) ([]model.ClientSecret, int64, error) {
	query := r.db.Model(&model.ClientSecret{}).Where("client_id = ?", clientID)
	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	items, total, err := helpers.List[model.ClientSecret](query, lq)
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

func (r *ClientSecretRepository) FindOne(clientID uuid.UUID, secret string) (*model.ClientSecret, error) {
	var result model.ClientSecret
	if err := r.db.Model(&model.ClientSecret{}).Where("client_id = ? AND client_secret = ?", clientID, secret).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *ClientSecretRepository) Insert(clientSecret *model.ClientSecret) (*model.ClientSecret, error) {
	if err := r.db.Clauses(clause.Returning{}).Create(clientSecret).Error; err != nil {
		return nil, err
	}
	return clientSecret, nil
}

func (r *ClientSecretRepository) Delete(clientID uuid.UUID, clientSecretID uuid.UUID) error {
	result := r.db.Where("id = ? AND client_id = ?", clientSecretID, clientID).Delete(&model.ClientSecret{})
	if result.Error != nil {
		return api_err.Wrap(api_err.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return api_err.Wrap(api_err.ErrNotFound, result.Error).WithMessage(fmt.Sprintf("Client secret with id %s does not exist", clientSecretID))
	}
	return nil
}
