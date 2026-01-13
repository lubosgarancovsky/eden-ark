package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-ark/internal/model"
	"github.com/lubosgarancovsky/eden-ark/internal/service"
	"github.com/lubosgarancovsky/eden-ark/pkg/helpers"
	"github.com/lubosgarancovsky/go-kit"
)

type userListingAttributes struct {
	FirstName string `rsql:"filter,sort"`
	LastName  string `rsql:"filter,sort"`
	Email     string `rsql:"filter,sort"`
	Username  string `rsql:"filter,sort"`
	CreatedAt string `rsql:"filter,sort"`
	IsActive  string `rsql:"filter,sort"`
	Role      string `rsql:"filter,sort"`
	DeletedAt string `rsql:"filter,sort"`
}

type UserHandler struct {
	s      *service.UserService
	parser *go_kit.Parser
}

func NewUserHandler(parser *go_kit.Parser, s *service.UserService) *UserHandler {
	return &UserHandler{s: s, parser: parser}
}

func (h *UserHandler) FindAll(c *gin.Context) {
	lq, err := helpers.CreateListingQuery(c, h.parser, &userListingAttributes{})
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
		c.Error(go_kit.Wrap(go_kit.ErrBadRequest, err))
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
		c.Error(go_kit.Wrap(go_kit.ErrBadRequest, err))
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
		c.Error(go_kit.Wrap(go_kit.ErrBadRequest, err))
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
		c.Error(go_kit.Wrap(go_kit.ErrBadRequest, err))
		return
	}

	if err := h.s.ResetPassword(input.Token, input.Password); err != nil {
		c.Error(err)
	}

	c.Status(204)
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
		c.Error(go_kit.Wrap(go_kit.ErrBadRequest, err))
	}

	user, err := h.s.ChangeEmail(&input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, user)
}

func (h *UserHandler) IsUsernameAvailable(c *gin.Context) {
	username, ok := c.Params.Get("username")
	if !ok {
		c.Error(go_kit.ErrParameterMissing.WithMessage(fmt.Sprintf("Path parameter %s is missing", "username")))
	}

	response := h.s.IsUsernameAvailable(username)
	c.JSON(200, response)
}

func (h *UserHandler) IsEmailAvailable(c *gin.Context) {
	email, ok := c.Params.Get("email")
	if !ok {
		c.Error(go_kit.ErrParameterMissing.WithMessage(fmt.Sprintf("Path parameter %s is missing", "email")))
	}

	response := h.s.IsEmailAvailable(email)
	c.JSON(200, response)
}
