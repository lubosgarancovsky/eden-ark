package service

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-ark/internal/model"
	"github.com/lubosgarancovsky/eden-ark/internal/repository"
	"github.com/lubosgarancovsky/go-kit"
)

type ClientService struct {
	r *repository.ClientRepository
}

func NewClientService(r *repository.ClientRepository) *ClientService {
	return &ClientService{r}
}

func (s *ClientService) FindAll(lq *go_kit.ListingQuery) (*go_kit.Page[model.Client], error) {
	items, totalCount, err := s.r.FindAll(lq)
	if err != nil {
		return nil, err
	}

	return &go_kit.Page[model.Client]{
		Items:      items,
		Page:       lq.Page,
		PageSize:   lq.Limit,
		TotalCount: totalCount,
	}, nil
}

func (s *ClientService) FindByID(ID uuid.UUID) (*model.Client, error) {
	return s.r.FindByID(ID)
}

func (s *ClientService) Create(input *model.ClientRequest) (*model.Client, error) {
	client := model.Client{
		Name:           input.Name,
		RedirectUris:   input.RedirectUris,
		GrantTypes:     input.GrantTypes,
		IsConfidential: input.IsConfidential,
	}
	return s.r.Insert(&client)
}

func (s *ClientService) Update(ID uuid.UUID, input *model.ClientRequest) (*model.Client, error) {
	client := model.Client{
		ID:             ID,
		Name:           input.Name,
		RedirectUris:   input.RedirectUris,
		GrantTypes:     input.GrantTypes,
		IsConfidential: input.IsConfidential,
	}
	return s.r.Update(&client)
}

func (s *ClientService) Delete(ID uuid.UUID) error {
	return s.r.Delete(ID)
}
