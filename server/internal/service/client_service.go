package service

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-arc/internal/model"
	"github.com/lubosgarancovsky/eden-arc/internal/repository"
	"github.com/lubosgarancovsky/go-kit/list"
)

type ClientService struct {
	r *repository.ClientRepository
}

func NewClientService(r *repository.ClientRepository) *ClientService {
	return &ClientService{r}
}

func (s *ClientService) FindAll(ls *list.ListingQuery) (list.Page[model.Client], error) {

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
