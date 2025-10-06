package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-arc/internal/listing"
	"github.com/lubosgarancovsky/eden-arc/internal/model"
	"github.com/lubosgarancovsky/eden-arc/internal/service"
	"github.com/lubosgarancovsky/eden-arc/pkg/helpers"
	"github.com/lubosgarancovsky/go-kit/rsql"
)

type ClientSecretHandler struct {
	s      *service.ClientSecretService
	parser *rsql.Parser
}

type ClientSecretPage struct {
	Items      []model.ClientSecret
	Page       int
	PageSize   int
	TotalCount int64
}

func NewClientSecretHandler(parser *rsql.Parser, s *service.ClientSecretService) *ClientSecretHandler {
	return &ClientSecretHandler{s, parser}
}

// FindAll @Summary      List client secrets
// @Description  Returns a paginated list of all secrets by client ID
// @Tags         Client secrets
// @Accept       json
// @Produce      json
// @Param        page      query     int     false  "Page number"     default(1)
// @Param        pageSize  query     int     false  "Items per page"  default(10)
// @Param        filter      query     string     false  "RSQL filter query"
// @Param        sort  query     string     false  "Sort query"
// @Param        clientId   path      string  true  "Client ID"
// @Success      200  {array}   ClientSecretPage
// @Router       /v1/arc/admin/clients/{clientId}/secrets [get]
func (h *ClientSecretHandler) FindAll(c *gin.Context) {
	lq, apiErr := helpers.CreateListingQuery(c, h.parser, listing.ClientSecretFilter, listing.ClientSecretFilter)
	if apiErr != nil {
		c.Error(apiErr)
		return
	}

	clientID, err := helpers.ExtractID(c, "clientId")
	if err != nil {
		c.Error(err)
		return
	}

	result, err := h.s.FindAll(clientID, lq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, result)
}

// Create @Summary      Create a new client secret
// @Description  Registers a new OAuth client secret by client ID
// @Tags         Client secrets
// @Accept       json
// @Produce      json
// @Param        client  body  model.ClientSecretRequest  true  "Client data"
// @Param        clientId   path      string  true  "Client ID"
// @Success      201  {object}  model.ClientSecret
// @Router       /v1/arc/admin/clients/{clientId}/secrets [post]
func (h *ClientSecretHandler) Create(c *gin.Context) {
	var input model.ClientSecretRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.Error(err)
		return
	}

	clientID, err := helpers.ExtractID(c, "clientId")
	if err != nil {
		c.Error(err)
		return
	}

	result, err := h.s.Create(clientID, &input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, result)
}

// Delete @Summary      Delete a client secret
// @Description  Deletes an OAuth client secret belonging to a specific client
// @Tags         Client secrets
// @Accept       json
// @Produce      json
// @Param        clientId        path      string  true  "Client ID"
// @Param        clientSecretId  path      string  true  "Secret ID"
// @Success      204  "No Content"
// @Router       /v1/arc/admin/clients/{clientId}/secrets/{clientSecretId} [delete]
func (h *ClientSecretHandler) Delete(c *gin.Context) {
	clientID, err := helpers.ExtractID(c, "clientId")
	if err != nil {
		c.Error(err)
		return
	}

	secretID, err := helpers.ExtractID(c, "clientSecretId")
	if err != nil {
		c.Error(err)
		return
	}

	if err := h.s.Delete(clientID, secretID); err != nil {
		c.Error(err)
		return
	}

	c.Status(204)
}
