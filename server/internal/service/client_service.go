package service

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-arc/internal/model"
	"github.com/lubosgarancovsky/eden-arc/internal/repository"
	"github.com/lubosgarancovsky/eden-arc/pkg/errors"
	"github.com/lubosgarancovsky/go-kit/list"
)

type ClientService struct {
	r *repository.ClientRepository
}

func NewClientService(r *repository.ClientRepository) *ClientService {
	return &ClientService{r}
}

func (s *ClientService) FindAll(lq *list.ListingQuery) (*list.Page[model.Client], error) {
	items, totalCount, err := s.r.FindAll(lq)
	if err != nil {
		return nil, err
	}

	return &list.Page[model.Client]{
		Items:      items,
		Page:       lq.Page,
		PageSize:   lq.Offset,
		TotalCount: totalCount,
	}, nil
}

func (s *ClientService) FindByID(ID uuid.UUID) (*model.Client, *errors.APIError) {
	return s.r.FindByID(ID)
}

func (s *ClientService) Create(input *model.ClientRequest) (*model.Client, *errors.APIError) {
	client := model.Client{
		Name:           input.Name,
		RedirectUris:   input.RedirectUris,
		GrantTypes:     input.GrantTypes,
		IsConfidential: input.IsConfidential,
	}
	return s.r.Insert(&client)
}

func (s *ClientService) Update(ID uuid.UUID, input *model.ClientRequest) (*model.Client, *errors.APIError) {
	client := model.Client{
		ID:             ID,
		Name:           input.Name,
		RedirectUris:   input.RedirectUris,
		GrantTypes:     input.GrantTypes,
		IsConfidential: input.IsConfidential,
	}
	return s.r.Update(&client)
}

func (s *ClientService) Delete(ID uuid.UUID) *errors.APIError {
	return s.r.Delete(ID)
}
