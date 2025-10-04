package repository

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-arc/internal/model"
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
	return helpers.List[model.Client](query, lq)
}

func (r *ClientRepository) FindByID(clientID uuid.UUID) (*model.Client, error) {
	var result model.Client
	err := r.db.Model(&model.Client{}).Where("id = ?", clientID).First(&result).Error
	return &result, err
}

func (r *ClientRepository) Insert(client *model.Client) (*model.Client, error) {
	err := r.db.Clauses(clause.Returning{}).Create(client).Error
	return client, err
}

func (r *ClientRepository) Update(client *model.Client) (*model.Client, error) {
	err := r.db.Clauses(clause.Returning{}).Where("id = ?", client.ID).Updates(&client).Error
	return client, err
}

func (r *ClientRepository) Delete(clientID uuid.UUID) error {
	result := r.db.Where("id = ?", clientID).Delete(&model.Client{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
