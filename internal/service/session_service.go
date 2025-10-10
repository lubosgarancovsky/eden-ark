package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-ark/internal/config"
	"github.com/lubosgarancovsky/eden-ark/internal/model"
	"github.com/lubosgarancovsky/eden-ark/internal/repository"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"github.com/lubosgarancovsky/go-kit/kit"
)

type SessionService struct {
	r   *repository.SessionRepository
	cfg *config.Config
}

func NewSessionService(cfg *config.Config, r *repository.SessionRepository) *SessionService {
	return &SessionService{cfg: cfg, r: r}
}

func (s *SessionService) FindByID(id uuid.UUID) (*model.Session, error) {
	return s.r.FindByID(id)
}

func (s *SessionService) FindByToken(token string) (*model.Session, error) {
	return s.r.FindByToken(token)
}

func (s *SessionService) Insert(userID uuid.UUID, ipAddr string, userAgent string, nonce string) (*model.Session, error) {
	token := kit.SHA256(32)
	if token == "" {
		return nil, api_err.ErrInternalServer
	}

	session := &model.Session{
		UserID:       userID,
		SessionToken: token,
		Nonce:        nonce,
		IPAddress:    &ipAddr,
		UserAgent:    &userAgent,
		ExpiresAt:    time.Now().Add(time.Duration(s.cfg.SessionExp) * time.Second),
	}

	return s.r.Insert(session)
}

func (s *SessionService) Delete(sessionID uuid.UUID) error {
	return s.r.Delete(sessionID)
}

func (s *SessionService) DeleteByToken(sessionToken string) error {
	return s.r.DeleteByToken(sessionToken)
}

func (s *SessionService) UpdateNonce(sessionID uuid.UUID, nonce string) (*model.Session, error) {
	return s.r.UpdateNonce(sessionID, nonce)
}
