package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-arc/internal/model"
	"github.com/lubosgarancovsky/eden-arc/internal/repository"
	"github.com/lubosgarancovsky/eden-arc/pkg/errors"
	"github.com/lubosgarancovsky/eden-arc/pkg/helpers"
)

type AuthCodeService struct {
	r *repository.AuthCodeRepository
}

func NewAuthCodeService(r *repository.AuthCodeRepository) *AuthCodeService {
	return &AuthCodeService{r}
}

func (s *AuthCodeService) FindByCode(code string) (*model.AuthorizationCode, error) {
	return s.r.FindByCode(code)
}

func (s *AuthCodeService) CreateAuthCode(session *model.Session, input *model.AuthorizeQuery) (*model.AuthorizationCode, error) {
	codeHash := helpers.SHA256(48)
	if codeHash == "" {
		return nil, errors.ErrInternalServer
	}

	clientID, err := uuid.Parse(input.ClientID)
	if err != nil {
		return nil, errors.Wrap(errors.ErrInvalidUUID, err)
	}

	code := &model.AuthorizationCode{
		Code:                codeHash,
		UserID:              session.UserID,
		ClientID:            clientID,
		RedirectURI:         input.RedirectURI,
		Scope:               input.Scope,
		CodeChallenge:       input.CodeChallenge,
		CodeChallengeMethod: input.CodeChallengeMethod,
		ExpiresAt:           time.Now().Add(time.Minute * 2), // TODO: Load from config
	}

	return s.r.Insert(code)
}

func (s *AuthCodeService) DeleteAuthCode(code string) error {
	return s.r.Delete(code)
}
