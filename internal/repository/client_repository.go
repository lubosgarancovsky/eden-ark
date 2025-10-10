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

type ClientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) *ClientRepository {
	return &ClientRepository{db: db}
}

func (r *ClientRepository) FindAll(lq *list.ListingQuery) ([]model.Client, int64, error) {
	query := r.db.Model(&model.Client{})
	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	items, total, err := helpers.List[model.Client](query, lq)
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

func (r *ClientRepository) FindByID(clientID uuid.UUID) (*model.Client, error) {
	var result model.Client
	if err := r.db.Model(&model.Client{}).Where("id = ?", clientID).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *ClientRepository) Insert(client *model.Client) (*model.Client, error) {
	if err := r.db.Clauses(clause.Returning{}).Create(client).Error; err != nil {
		return nil, err
	}
	return client, nil
}

func (r *ClientRepository) Update(client *model.Client) (*model.Client, error) {
	result := r.db.Clauses(clause.Returning{}).Where("id = ?", client.ID).Updates(&client)
	if result.Error != nil {
		return nil, api_err.Wrap(api_err.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, api_err.Wrap(api_err.ErrNotFound, result.Error).WithMessage(fmt.Sprintf("Client with id %s does not exist", client.ID))
	}
	return client, nil
}

func (r *ClientRepository) Delete(clientID uuid.UUID) error {
	result := r.db.Where("id = ?", clientID).Delete(&model.Client{})
	if result.Error != nil {
		return api_err.Wrap(api_err.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return api_err.ErrNotFound.WithMessage(fmt.Sprintf("Client with id %s does not exist", clientID))
	}
	return nil
}
