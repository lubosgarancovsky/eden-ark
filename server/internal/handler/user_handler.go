package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-arc/internal/model"
	"github.com/lubosgarancovsky/eden-arc/internal/service"
	"github.com/lubosgarancovsky/go-kit/api_err"
)

type UserHandler struct {
	s *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{s}
}

func (h *UserHandler) RequestResetPassword(c *gin.Context) {
	var input model.EmailRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(api_err.Wrap(api_err.ErrBadRequest, err))
		return
	}

	if err := h.s.RequestPasswordChange(input.Email); err != nil {
		c.Error(err)
		return
	}

	c.JSON(204, "")
}

func (h *UserHandler) ResetPassword(c *gin.Context) {
	var input model.PasswordRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(api_err.Wrap(api_err.ErrBadRequest, err))
		return
	}

	if err := h.s.ResetPassword(input.Token, input.Password); err != nil {
		c.Error(err)
	}

	c.JSON(204, "")
}
