package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-arc/internal/config"
	"github.com/lubosgarancovsky/eden-arc/internal/model"
	"github.com/lubosgarancovsky/eden-arc/internal/repository"
	"github.com/lubosgarancovsky/eden-arc/pkg/errors"
	"github.com/lubosgarancovsky/eden-arc/pkg/helpers"
)

type SessionService struct {
	r   *repository.SessionRepository
	cfg *config.Config
}

func NewSessionService(cfg *config.Config, r *repository.SessionRepository) *SessionService {
	return &SessionService{cfg: cfg, r: r}
}

func (s *SessionService) FindByToken(token string) (*model.Session, error) {
	return s.r.FindByToken(token)
}

func (s *SessionService) Insert(userID uuid.UUID, ipAddr string, userAgent string) (*model.Session, error) {
	token := helpers.SHA256(32)
	if token == "" {
		return nil, errors.ErrInternalServer
	}

	session := &model.Session{
		UserID:       userID,
		SessionToken: token,
		IPAddress:    &ipAddr,
		UserAgent:    &userAgent,
		ExpiresAt:    time.Now().Add(time.Duration(s.cfg.SessionExp) * time.Second),
	}

	return s.r.Insert(session)
}

func (s *SessionService) Delete(sessionID uuid.UUID) error {
	return s.r.Delete(sessionID)
}
