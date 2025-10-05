package service

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-arc/internal/model"
	"github.com/lubosgarancovsky/eden-arc/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	r *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{r: r}
}

func (s *UserService) FindByID(id uuid.UUID) (*model.User, error) {
	return s.r.FindByID(id)
}

func (s *UserService) FindByUsername(username string) (*model.User, error) {
	return s.r.FindByUsername(username)
}

func (s *UserService) FindByEmail(email string) (*model.User, error) {
	return s.r.FindByEmail(email)
}

func (s *UserService) Insert(user *model.User) (*model.User, error) {
	return s.r.Insert(user)
}

func (s *UserService) Update(user *model.User) (*model.User, error) {
	return s.r.Update(user)
}

func (s *UserService) Delete(userID uuid.UUID) error {
	return s.r.Delete(userID)
}

func (s *UserService) MatchPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *UserService) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
