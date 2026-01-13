package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-ark/internal/model"
	"github.com/lubosgarancovsky/go-kit"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) *SessionRepository {
	return &SessionRepository{db}
}

func (r *SessionRepository) FindByID(id uuid.UUID) (*model.Session, error) {
	var result model.Session
	if err := r.db.Model(&model.Session{}).Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *SessionRepository) FindByToken(token string) (*model.Session, error) {
	var result model.Session
	if err := r.db.Model(&model.Session{}).Where("session_token = ?", token).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *SessionRepository) Insert(session *model.Session) (*model.Session, error) {
	if err := r.db.Clauses(clause.Returning{}).Create(session).Error; err != nil {
		return nil, err
	}
	return session, nil
}

func (r *SessionRepository) Delete(sessionID uuid.UUID) error {
	result := r.db.Where("id = ?", sessionID).Delete(&model.Session{})
	if result.Error != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return go_kit.Wrap(go_kit.ErrNotFound, result.Error).WithMessage(fmt.Sprintf("Session with id %s does not exist", sessionID))
	}
	return nil
}

func (r *SessionRepository) DeleteByToken(sessionToken string) error {
	result := r.db.Where("session_token = ?", sessionToken).Delete(&model.Session{})
	if result.Error != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return go_kit.Wrap(go_kit.ErrNotFound, result.Error)
	}
	return nil
}

func (r *SessionRepository) UpdateNonce(sessionID uuid.UUID, nonce string) (*model.Session, error) {
	session := &model.Session{}

	result := r.db.Model(session).
		Clauses(clause.Returning{}).
		Where("id = ?", sessionID).
		Update("nonce", nonce)
	if result.Error != nil {
		return nil, go_kit.Wrap(go_kit.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, go_kit.Wrap(go_kit.ErrNotFound, nil)
	}

	return session, nil
}
