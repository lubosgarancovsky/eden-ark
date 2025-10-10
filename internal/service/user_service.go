package service

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-ark/internal/config"
	"github.com/lubosgarancovsky/eden-ark/internal/model"
	"github.com/lubosgarancovsky/eden-ark/internal/repository"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"github.com/lubosgarancovsky/go-kit/kit"
	"github.com/lubosgarancovsky/go-kit/list"
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

func (s *UserService) FindAll(lq *list.ListingQuery) (*list.Page[model.User], error) {
	items, totalCount, err := s.r.FindAll(lq)
	if err != nil {
		return nil, err
	}

	return &list.Page[model.User]{
		Items:      items,
		Page:       lq.Page,
		PageSize:   lq.Limit,
		TotalCount: totalCount,
	}, nil
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
	user, err := s.FindByID(userID)
	if err != nil {
		return err
	}

	if user.DeletedAt.IsZero() {
		return s.r.MarkAsDeleted(userID)
	}

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

	templateData := &model.PasswordResetTemplate{
		Name:             user.FirstName,
		AppName:          "Eden",
		Year:             time.Now().Year(),
		Token:            token.Token,
		ExpiresAt:        token.ExpiresAt.Format(time.RFC3339Nano),
		ExpiresInMinutes: int(time.Until(token.ExpiresAt).Minutes()) + 1,
		ResetURL:         fmt.Sprintf("%s/reset-password", s.cfg.PublicURL),
	}

	templatePath := "templates/reset-password.html"
	return s.emailService.SendTemplateEmail(email, "Eden - Password reset", templatePath, templateData)
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

func (s *UserService) IsUsernameAvailable(input model.IsAvailableRequest) *model.IsAvailableResponse {
	_, err := s.r.FindByUsername(input.Value)
	return &model.IsAvailableResponse{
		IsAvailable: err != nil,
	}
}

func (s *UserService) IsEmailAvailable(input model.IsAvailableRequest) *model.IsAvailableResponse {
	_, err := s.r.FindByEmail(input.Value)
	return &model.IsAvailableResponse{
		IsAvailable: err != nil,
	}
}

func (s *UserService) RequestEmailChange(userID uuid.UUID) error {
	user, err := s.FindByID(userID)
	if err != nil {
		return err
	}

	token, err := s.recoveryTokenService.Insert(user.ID, "email", user.Email)
	if err != nil {
		return err
	}

	templateData := &model.EmailChangeTemplate{
		Name:             user.FirstName,
		AppName:          "Eden",
		Year:             time.Now().Year(),
		Token:            token.Token,
		ConfirmURL:       fmt.Sprintf("%s/change-email", s.cfg.PublicURL),
		ExpiresAt:        token.ExpiresAt.Format(time.RFC3339Nano),
		ExpiresInMinutes: int(time.Until(token.ExpiresAt).Minutes()) + 1,
	}

	templatePath := "templates/change-email.html"
	return s.emailService.SendTemplateEmail(user.Email, "Eden - Change email", templatePath, templateData)
}

func (s *UserService) GeneratePassword(userID uuid.UUID) error {
	user, err := s.FindByID(userID)
	if err != nil {
		return err
	}

	password, err := kit.GeneratePassword(12)
	if err != nil {
		return err
	}

	passwordHash, err := s.HashPassword(password)

	if err != nil {
		return err
	}

	user.PasswordHash = passwordHash
	_, err = s.r.Update(user)
	if err != nil {
		return err
	}

	templateData := &model.NewUserTemplate{
		Name:     fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		Email:    user.Email,
		Password: password,
		AppName:  "Eden",
		LoginURL: fmt.Sprintf("%s/login", s.cfg.PublicURL),
		Year:     time.Now().Year(),
	}

	templatePath := "templates/new-user-credentials.html"
	return s.emailService.SendTemplateEmail(user.Email, "Eden - Change email", templatePath, templateData)
}

func (s *UserService) ChangeEmail(input *model.EmailChangeRequest) (*model.User, error) {
	token, err := s.recoveryTokenService.FindOne(input.Token)
	if err != nil {
		return nil, err
	}

	user, err := s.FindByID(token.UserID)
	if err != nil {
		return nil, err
	}
	if token.UserID != user.ID {
		return nil, api_err.ErrBadRequest.WithMessage("invalid token")
	}

	user.Email = input.Email
	updatedUser, err := s.r.Update(user)
	if err != nil {
		return nil, err
	}

	s.recoveryTokenService.Delete(token.Token)
	return updatedUser, nil
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
