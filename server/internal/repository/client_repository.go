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
		return nil, 0, errors.Wrap(errors.ErrInternalServer, err)
	}
	return items, total, nil
}

func (r *ClientRepository) FindByID(clientID uuid.UUID) (*model.Client, *errors.APIError) {
	var result model.Client
	if err := r.db.Model(&model.Client{}).Where("id = ?", clientID).First(&result).Error; err != nil {
		return nil, errors.Wrap(errors.ErrInternalServer, err)
	}
	return &result, nil
}

func (r *ClientRepository) Insert(client *model.Client) (*model.Client, *errors.APIError) {
	if err := r.db.Clauses(clause.Returning{}).Create(client).Error; err != nil {
		return nil, errors.Wrap(errors.ErrInternalServer, err)
	}
	return client, nil
}

func (r *ClientRepository) Update(client *model.Client) (*model.Client, *errors.APIError) {
	result := r.db.Clauses(clause.Returning{}).Where("id = ?", client.ID).Updates(&client)
	if result.Error != nil {
		return nil, errors.Wrap(errors.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, errors.Wrap(errors.ErrNotFound, result.Error).WithMessage(fmt.Sprintf("Client with id %s does not exist", client.ID))
	}
	return client, nil
}

func (r *ClientRepository) Delete(clientID uuid.UUID) *errors.APIError {
	result := r.db.Where("id = ?", clientID).Delete(&model.Client{})
	if result.Error != nil {
		return errors.Wrap(errors.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return errors.ErrNotFound.WithMessage(fmt.Sprintf("Client with id %s does not exist", clientID))
	}
	return nil
}
