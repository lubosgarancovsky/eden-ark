package service

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-ark/internal/config"
	"github.com/lubosgarancovsky/eden-ark/internal/model"
	"github.com/lubosgarancovsky/eden-ark/internal/repository"
	"github.com/lubosgarancovsky/go-kit"
)

// 5MB
const MaxSize = 5 << 20

type UserAvatarService struct {
	cfg         *config.Config
	userService *UserService
	repository  *repository.UserAvatarRepository
}

func NewUserAvatarService(cfg *config.Config, userService *UserService, repository *repository.UserAvatarRepository) *UserAvatarService {
	return &UserAvatarService{cfg, userService, repository}
}

func (s *UserAvatarService) RemoveAvatar(userID uuid.UUID) error {
	return s.repository.Delete(userID)
}

func (s *UserAvatarService) UploadAvatar(c *gin.Context, userID uuid.UUID) (*model.UserAvatar, error) {
	user, err := s.userService.FindByID(userID)
	if err != nil {
		return nil, err
	}

	file, err := c.FormFile("file")
	if err != nil {
		return nil, go_kit.ErrBadRequest.WithMessage("No file found")
	}

	if file.Size > MaxSize {
		return nil, go_kit.ErrBadRequest.WithMessage("Uploaded file is too big.")
	}

	avatarMime, err := validateMimeType(file)
	if err != nil {
		return nil, err
	}

	avatarVersion := 1
	if user.AvatarVersion != nil {
		avatarVersion = *user.AvatarVersion + 1
	}

	folder := fmt.Sprintf("user-%s", user.ID.String())
	fileName := getFileName(avatarMime, avatarVersion)

	if fileName == "" {
		return nil, go_kit.ErrBadRequest.WithMessage("Invalid file name")
	}

	dst := filepath.Join(s.cfg.UploadsFolder, "avatars", folder, fileName)
	if err = c.SaveUploadedFile(file, dst); err != nil {
		return nil, err
	}

	if err := s.repository.Change(userID, avatarMime, avatarVersion); err != nil {
		return nil, err
	}

	return &model.UserAvatar{
		MimeType: avatarMime,
		Version:  avatarVersion,
	}, nil
}

func getFileName(avatarMime string, avatarVersion int) string {
	parts := strings.Split(avatarMime, "/")
	if len(parts) != 2 {
		return ""
	}

	extension := parts[1]
	return fmt.Sprintf("v%d.%s", avatarVersion, extension)
}

func validateMimeType(file *multipart.FileHeader) (string, error) {
	openedFile, err := file.Open()
	if err != nil {
		return "", err
	}
	defer openedFile.Close()

	buffer := make([]byte, 512) // first 512 bytes for MIME sniffing
	if _, err := openedFile.Read(buffer); err != nil {
		return "", go_kit.ErrBadRequest.WithMessage("Failed to read multipart file")
	}

	mimeType := http.DetectContentType(buffer)
	allowed := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/webp": true,
	}
	if !allowed[mimeType] {
		return "", go_kit.ErrBadRequest.WithMessage("File of this type is not supported.")
	}

	return mimeType, nil
}
