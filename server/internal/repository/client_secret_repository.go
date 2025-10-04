package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-arc/internal/model"
	"github.com/lubosgarancovsky/eden-arc/pkg/errors"
	"github.com/lubosgarancovsky/eden-arc/pkg/helpers"
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
		return nil, 0, errors.Wrap(errors.ErrInternalServer, err)
	}
	return items, total, nil
}

func (r *ClientSecretRepository) Insert(clientSecret *model.ClientSecret) (*model.ClientSecret, error) {
	if err := r.db.Clauses(clause.Returning{}).Create(clientSecret).Error; err != nil {
		return nil, errors.Wrap(errors.ErrInternalServer, err)
	}
	return clientSecret, nil
}

func (r *ClientSecretRepository) Delete(clientID uuid.UUID, clientSecretID uuid.UUID) error {
	result := r.db.Where("id = ? AND client_id = ?", clientSecretID, clientID).Delete(&model.ClientSecret{})
	if result.Error != nil {
		return errors.Wrap(errors.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return errors.Wrap(errors.ErrNotFound, result.Error).WithMessage(fmt.Sprintf("Client secret with id %s does not exist", clientSecretID))
	}
	return nil
}
