package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-arc/internal/model"
	"github.com/lubosgarancovsky/eden-arc/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) *SessionRepository {
	return &SessionRepository{db}
}

func (r *SessionRepository) FindByToken(token string) (*model.Session, error) {
	var result model.Session
	if err := r.db.Model(&model.Session{}).Where("session_token = ?", token).First(&result).Error; err != nil {
		return nil, errors.Wrap(errors.ErrInternalServer, err)
	}
	return &result, nil
}

func (r *SessionRepository) Insert(session *model.Session) (*model.Session, error) {
	if err := r.db.Clauses(clause.Returning{}).Create(session).Error; err != nil {
		return nil, errors.Wrap(errors.ErrInternalServer, err)
	}
	return session, nil
}

func (r *SessionRepository) Delete(sessionID uuid.UUID) error {
	result := r.db.Where("id = ?", sessionID).Delete(&model.Session{})
	if result.Error != nil {
		return errors.Wrap(errors.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return errors.Wrap(errors.ErrNotFound, result.Error).WithMessage(fmt.Sprintf("Session with id %s does not exist", sessionID))
	}
	return nil
}
