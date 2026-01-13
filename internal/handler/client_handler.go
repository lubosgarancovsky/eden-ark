package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-ark/internal/model"
	"github.com/lubosgarancovsky/eden-ark/internal/service"
	"github.com/lubosgarancovsky/eden-ark/pkg/helpers"
	"github.com/lubosgarancovsky/go-kit"
)

type ClientHandler struct {
	s      *service.ClientService
	parser *go_kit.Parser
}

type ClientPage struct {
	Items      []model.Client
	Page       int
	PageSize   int
	TotalCount int64
}

type clientListingAttributes struct {
	Name      string `rsql:"filter,sort"`
	CreatedAt string `rsql:"filter,sort"`
}

func NewClientHandler(parser *go_kit.Parser, s *service.ClientService) *ClientHandler {
	return &ClientHandler{s, parser}
}

// FindAll @Summary      List clients
// @Description  Returns a paginated list of all clients
// @Tags         Clients
// @Accept       json
// @Produce      json
// @Param        page      query     int     false  "Page number"     default(1)
// @Param        pageSize  query     int     false  "Items per page"  default(10)
// @Param        filter      query     string     false  "RSQL filter query"
// @Param        sort  query     string     false  "Sort query"
// @Success      200  {array}   ClientPage
// @Router       /v1/ark/admin/clients [get]
func (h *ClientHandler) FindAll(c *gin.Context) {
	lq, apiErr := helpers.CreateListingQuery(c, h.parser, &clientListingAttributes{})
	if apiErr != nil {
		c.Error(apiErr)
		return
	}

	result, err := h.s.FindAll(lq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, result)
}

// FindByID @Summary      Get client by ID
// @Description  Returns a client by its ID
// @Tags         Clients
// @Accept       json
// @Produce      json
// @Param        clientId   path      string  true  "Client ID"
// @Success      200  {array}   model.Client
// @Router       /v1/ark/admin/clients/{clientId} [get]
func (h *ClientHandler) FindByID(c *gin.Context) {
	UID, err := helpers.ExtractID(c, "clientId")
	if err != nil {
		c.Error(err)
		return
	}

	result, err := h.s.FindByID(UID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, result)
}

// Create @Summary      Create a new client
// @Description  Registers a new OAuth client in the IAM system
// @Tags         Clients
// @Accept       json
// @Produce      json
// @Param        client  body  model.ClientRequest  true  "Client data"
// @Success      201  {object}  model.Client
// @Router       /v1/ark/admin/clients [post]
func (h *ClientHandler) Create(c *gin.Context) {
	var input model.ClientRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.Error(err)
		return
	}

	result, err := h.s.Create(&input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, result)
}

// Update @Summary      Update a client
// @Description  Updates an existing OAuth client in the IAM system
// @Tags         Clients
// @Accept       json
// @Produce      json
// @Param        client  body  model.ClientRequest  true  "Client data"
// @Param        clientId   path      string  true  "Client ID"
// @Success      200  {object}  model.Client
// @Router       /v1/ark/admin/clients/{clientId} [put]
func (h *ClientHandler) Update(c *gin.Context) {
	var input model.ClientRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.Error(err)
	}

	UID, err := helpers.ExtractID(c, "clientId")
	if err != nil {
		c.Error(err)
		return
	}

	result, err := h.s.Update(UID, &input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, result)
}

// Delete @Summary      Delete a client
// @Description  Deletes the OAuth client from the IAM system
// @Tags         Clients
// @Accept       json
// @Produce      json
// @Param        clientId   path      string  true  "Client ID"
// @Success      204  {string}  string  "No Content"
// @Router       /v1/ark/admin/clients/{clientId} [delete]
func (h *ClientHandler) Delete(c *gin.Context) {
	UID, err := helpers.ExtractID(c, "clientId")
	if err != nil {
		c.Error(err)
		return
	}

	err = h.s.Delete(UID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(204, nil)
}
