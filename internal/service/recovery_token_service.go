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

type RecoveryTokenService struct {
	cfg *config.Config
	r   *repository.RecoveryTokenRepository
}

func NewRecoveryTokenService(cfg *config.Config, r *repository.RecoveryTokenRepository) *RecoveryTokenService {
	return &RecoveryTokenService{cfg: cfg, r: r}
}

func (s *RecoveryTokenService) FindOne(token string) (*model.RecoveryToken, error) {
	return s.r.FindOne(token)
}

func (s *RecoveryTokenService) Insert(userID uuid.UUID, recoveryType string, metadata string) (*model.RecoveryToken, error) {
	token := kit.SHA256(32)
	if token == "" {
		return nil, api_err.ErrInternalServer
	}

	recoveryToken := &model.RecoveryToken{
		UserID:       userID,
		Token:        token,
		RecoveryType: recoveryType,
		Metadata:     metadata,
		ExpiresAt:    time.Now().Add(time.Duration(s.cfg.RecoveryTokenExp) * time.Second),
	}

	return s.r.Insert(recoveryToken)
}

func (s *RecoveryTokenService) Delete(token string) error {
	return s.r.Delete(token)
}
