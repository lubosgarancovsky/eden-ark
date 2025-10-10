package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-ark/internal/model"
	"github.com/lubosgarancovsky/eden-ark/internal/repository"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"github.com/lubosgarancovsky/go-kit/kit"
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
	codeHash := kit.SHA256(48)
	if codeHash == "" {
		return nil, api_err.ErrInternalServer
	}

	clientID, err := uuid.Parse(input.ClientID)
	if err != nil {
		return nil, api_err.Wrap(api_err.ErrInvalidUUID, err)
	}

	code := &model.AuthorizationCode{
		Code:                codeHash,
		UserID:              session.UserID,
		ClientID:            clientID,
		SessionID:           session.ID,
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
