package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-ark/internal/service"
	"github.com/lubosgarancovsky/eden-ark/pkg/helpers"
)

type UserAvatarHandler struct {
	service *service.UserAvatarService
}

func NewUserAvatarHandler(service *service.UserAvatarService) *UserAvatarHandler {
	return &UserAvatarHandler{service}
}

func (h *UserAvatarHandler) RemoveAvatar(c *gin.Context) {
	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	if err = h.service.RemoveAvatar(user.ID); err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *UserAvatarHandler) UploadAvatar(c *gin.Context) {
	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	result, err := h.service.UploadAvatar(c, user.ID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, result)
}
