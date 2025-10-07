package service

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-arc/internal/config"
	"github.com/lubosgarancovsky/eden-arc/internal/model"
	"github.com/lubosgarancovsky/eden-arc/internal/repository"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	r                    *repository.UserRepository
	cfg                  *config.Config
	recoveryTokenService *RecoveryTokenService
	emailService         *EmailService
}

func NewUserService(
	cfg *config.Config,
	r *repository.UserRepository,
	recoveryTokenService *RecoveryTokenService,
	emailService *EmailService) *UserService {
	return &UserService{r: r, cfg: cfg, recoveryTokenService: recoveryTokenService, emailService: emailService}
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

func (s *UserService) Insert(userRequest *model.CreateUserRequest) (*model.User, error) {
	passwordHash, err := s.HashPassword(userRequest.Password)
	if err != nil {
		return nil, err
	}

	role := "USER"
	if userRequest.IsAdmin {
		role = "ADMIN"
	}

	user := &model.User{
		Username:     userRequest.Username,
		Email:        userRequest.Email,
		PasswordHash: passwordHash,
		FirstName:    userRequest.FirsName,
		LastName:     userRequest.LastName,
		IsActive:     true,
		Role:         role,
	}

	return s.r.Insert(user)
}

func (s *UserService) Update(userID uuid.UUID, userRequest *model.UpdateUserRequest) (*model.User, error) {
	role := "USER"
	if userRequest.IsAdmin {
		role = "ADMIN"
	}

	user := &model.User{
		ID:        userID,
		FirstName: userRequest.FirstName,
		LastName:  userRequest.LastName,
		Role:      role,
	}

	return s.r.Update(user)
}

func (s *UserService) Delete(userID uuid.UUID) error {
	return s.r.Delete(userID)
}

func (s *UserService) RequestPasswordChange(email string) error {
	user, err := s.FindByEmail(email)
	if err != nil {
		return err
	}

	token, err := s.recoveryTokenService.Insert(user.ID, "password", email)
	if err != nil {
		return err
	}

	emailTemplateData := &model.PasswordResetTemplate{
		Name:             user.FirstName,
		AppName:          "Eden",
		Year:             time.Now().Year(),
		Token:            token.Token,
		ExpiresAt:        token.ExpiresAt.Format(time.RFC3339Nano),
		ExpiresInMinutes: int(time.Until(token.ExpiresAt).Minutes()) + 1,
		ResetURL:         fmt.Sprintf("%s/reset-password", s.cfg.PublicURL),
	}

	templatePath := "templates/reset-password.html"
	return s.emailService.SendTemplateEmail(email, "Password reset", templatePath, emailTemplateData)
}

func (s *UserService) ResetPassword(tokenString string, password string) error {
	token, err := s.recoveryTokenService.FindOne(tokenString)
	if err != nil {
		return err
	}

	user, err := s.FindByID(token.UserID)
	if err != nil {
		return err
	}
	if token.UserID != user.ID {
		return api_err.ErrBadRequest.WithMessage("invalid token")
	}

	passwordHash, err := s.HashPassword(password)
	if err != nil {
		return err
	}

	user.PasswordHash = passwordHash
	if _, err := s.r.Update(user); err != nil {
		return err
	}

	s.recoveryTokenService.Delete(token.Token)
	return nil
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
