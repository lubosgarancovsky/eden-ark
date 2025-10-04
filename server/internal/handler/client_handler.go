package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-arc/internal/model"
	"github.com/lubosgarancovsky/eden-arc/internal/service"
	"github.com/lubosgarancovsky/go-kit/rsql"
)

type ClientHandler struct {
	s      *service.ClientService
	parser *rsql.Parser
}

func NewClientHandler(parser *rsql.Parser, s *service.ClientService) *ClientHandler {
	return &ClientHandler{s, parser}
}

func (h *ClientHandler) FindAll(c *gin.Context) {
	//var qp list.QueryParms
	//if err := c.ShouldBindQuery(&lq); err != nil {
	//	c.Error(err)
	//	return
	//}
	//
	//lq := list.ListingQuery{
	//	Limit: 0,
	//	Offset: 0,
	//	Filter: nil,
	//	Sort: nil
	//}
	//
	//result, err := h.s.FindAll(&lq)
	//if err != nil {
	//	c.Error(err)
	//	return
	//}
	//
	//c.JSON(200, result)
}

func (h *ClientHandler) FindByID(c *gin.Context) {
	idString, ok := c.Params.Get("id")
	if !ok {
		c.Error(fmt.Errorf("id is required"))
		return
	}

	ID, err := uuid.Parse(idString)
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
