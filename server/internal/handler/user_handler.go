package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-arc/internal/listing"
	"github.com/lubosgarancovsky/eden-arc/internal/model"
	"github.com/lubosgarancovsky/eden-arc/internal/service"
	"github.com/lubosgarancovsky/eden-arc/pkg/helpers"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"github.com/lubosgarancovsky/go-kit/rsql"
)

type UserHandler struct {
	s      *service.UserService
	parser *rsql.Parser
}

func NewUserHandler(parser *rsql.Parser, s *service.UserService) *UserHandler {
	return &UserHandler{s: s, parser: parser}
}

func (h *UserHandler) FindAll(c *gin.Context) {
	lq, err := helpers.CreateListingQuery(c, h.parser, listing.UserFilter, listing.UserSort)
	if err != nil {
		c.Error(err)
		return
	}

	page, err := h.s.FindAll(lq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, page)
}

func (h *UserHandler) FindByID(c *gin.Context) {
	ID, err := helpers.ExtractID(c, "userId")
	if err != nil {
		c.Error(err)
		return
	}

	result, err := h.s.FindByID(ID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, result)
}

func (h *UserHandler) Create(c *gin.Context) {
	var input model.CreateUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(api_err.Wrap(api_err.ErrBadRequest, err))
		return
	}

	user, err := h.s.Insert(&input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, user)
}

func (h *UserHandler) Update(c *gin.Context) {
	var input model.UpdateUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(api_err.Wrap(api_err.ErrBadRequest, err))
		return
	}

	ID, err := helpers.ExtractID(c, "userId")
	if err != nil {
		c.Error(err)
		return
	}

	user, err := h.s.Update(ID, &input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, user)
}

func (h *UserHandler) Delete(c *gin.Context) {
	ID, err := helpers.ExtractID(c, "userId")
	if err != nil {
		c.Error(err)
		return
	}

	if err := h.s.Delete(ID); err != nil {
		c.Error(err)
		return
	}

	c.Status(204)
}

func (h *UserHandler) GeneratePassword(c *gin.Context) {
	ID, err := helpers.ExtractID(c, "userId")
	if err != nil {
		c.Error(err)
		return
	}

	if err := h.s.GeneratePassword(ID); err != nil {
		c.Error(err)
		return
	}

	c.Status(204)
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

	c.Status(204)
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

func (h *UserHandler) RequestChangeEmail(c *gin.Context) {
	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	if err := h.s.RequestEmailChange(user.ID); err != nil {
		c.Error(err)
		return
	}

	c.Status(204)
}

func (h *UserHandler) ChangeEmail(c *gin.Context) {
	var input model.EmailChangeRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(api_err.Wrap(api_err.ErrBadRequest, err))
	}

	user, err := h.s.ChangeEmail(&input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, user)
}

func (h *UserHandler) IsUsernameAvailable(c *gin.Context) {
	var input model.IsAvailableRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(api_err.Wrap(api_err.ErrBadRequest, err))
		return
	}

	response := h.s.IsEmailAvailable(input)
	c.JSON(200, response)
}
