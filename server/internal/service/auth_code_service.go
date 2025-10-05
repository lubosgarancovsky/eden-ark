package service

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-arc/internal/repository"
	"github.com/lubosgarancovsky/eden-arc/pkg/errors"
)

type AuthCodeService struct {
	r *repository.AuthCodeRepository
}

func NewAuthCodeService(r *repository.AuthCodeRepository) *AuthCodeService {
	return &AuthCodeService{r}
}

//func (s *AuthCodeService) CreateAuthCode(userID uuid.UUID, input interface{}) (*model.AuthorizationCode, error) {
//	codeHash := generateAuthCode()
//	if codeHash == "" {
//		return nil, errors.ErrInternalServer
//	}
//
//	code := &model.AuthorizationCode{
//		Code:                codeHash,
//		UserID:              userID,
//		ClientID:            input.ClientID,
//		RedirectURI:         input.RedirectURI,
//		Scopes:              []string{},
//		CodeChallenge:       input.CodeChallenge,
//		CodeChallengeMethod: input.CodeChallengeMethod,
//		ExpiresAt:           time.Now().Add(time.Minute * 2),
//	}
//
//	return s.r.Insert(code)
//}

func (s *AuthCodeService) DeleteAuthCode(userID uuid.UUID, clientID uuid.UUID, code string) error {
	authCode, err := s.r.FindByCode(code)
	if err != nil {
		return err
	}

	if authCode.UserID != userID {
		return errors.ErrNotFound
	}

	if authCode.ClientID != clientID {
		return errors.ErrNotFound
	}

	return s.r.Delete(code)
}

func generateAuthCode() string {
	randomBytes := make([]byte, 48)
	if _, err := rand.Read(randomBytes); err != nil {
		return ""
	}

	rawCode := base64.RawURLEncoding.EncodeToString(randomBytes)

	hash := sha256.Sum256([]byte(rawCode))
	codeHash := fmt.Sprintf("%x", hash[:])

	return codeHash
}
