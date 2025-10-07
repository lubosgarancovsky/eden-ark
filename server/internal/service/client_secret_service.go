package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-arc/internal/model"
	"github.com/lubosgarancovsky/eden-arc/internal/repository"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"github.com/lubosgarancovsky/go-kit/list"

	"crypto/rand"
	"encoding/base64"
)

type ClientSecretService struct {
	r *repository.ClientSecretRepository
}

func NewClientSecretService(r *repository.ClientSecretRepository) *ClientSecretService {
	return &ClientSecretService{r}
}

func (s *ClientSecretService) FindAll(clientID uuid.UUID, lq *list.ListingQuery) (*list.Page[model.ClientSecret], error) {
	items, totalCount, err := s.r.FindAll(clientID, lq)
	if err != nil {
		return nil, err
	}

	var masked []model.ClientSecret
	for _, item := range items {
		masked = append(masked, model.ClientSecret{
			ID:           item.ID,
			ClientID:     item.ClientID,
			ClientSecret: maskClientSecret(item.ClientSecret, 4, 4),
			ExpiresAt:    item.ExpiresAt,
		})
	}

	return &list.Page[model.ClientSecret]{
		Items:      items,
		Page:       lq.Page,
		PageSize:   lq.Limit,
		TotalCount: totalCount,
	}, nil
}

func (s *ClientSecretService) ValidateSecret(clientID uuid.UUID, secretStr string) (*model.ClientSecret, error) {
	secret, err := s.r.FindOne(clientID, secretStr)
	if err != nil {
		return nil, err
	}

	if secret.ExpiresAt.Before(time.Now()) {
		return nil, api_err.ErrUnauthorized.WithMessage("secret expired")
	}

	return secret, nil
}

func (s *ClientSecretService) Create(clientID uuid.UUID, input *model.ClientSecretRequest) (*model.ClientSecret, error) {
	secretString, err := generateClientSecret(32)
	if err != nil {
		return nil, api_err.Wrap(api_err.ErrInternalServer, err)
	}

	secret := &model.ClientSecret{
		ClientID:     clientID,
		ClientSecret: secretString,
		ExpiresAt:    input.ExpiresAt,
	}
	return s.r.Insert(secret)
}

func (s *ClientSecretService) Delete(clientID uuid.UUID, clientSecretID uuid.UUID) error {
	return s.r.Delete(clientID, clientSecretID)
}

func generateClientSecret(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(bytes), nil
}

func maskClientSecret(secret string, visibleStart, visibleEnd int) string {
	length := len(secret)
	if length == 0 {
		return ""
	}
	if visibleStart+visibleEnd >= length {
		return secret
	}

	masked := make([]byte, length)
	copy(masked[:visibleStart], secret[:visibleStart])
	for i := visibleStart; i < length-visibleEnd; i++ {
		masked[i] = '*'
	}
	copy(masked[length-visibleEnd:], secret[length-visibleEnd:])
	return string(masked)
}
