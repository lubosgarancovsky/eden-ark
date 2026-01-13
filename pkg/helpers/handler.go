package helpers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-ark/internal/model"
	"github.com/lubosgarancovsky/go-kit"
)

func CreateListingQuery(c *gin.Context, parser *go_kit.Parser, listingAttr interface{}) (*go_kit.ListingQuery, error) {
	qp := &go_kit.QueryParams{}
	if err := c.ShouldBindQuery(&qp); err != nil {
		return nil, go_kit.Wrap(go_kit.ErrBadRequest, err).WithMessage("Invalid query parameters")
	}

	lq, err := go_kit.NewListingQuery(qp, parser, listingAttr)
	if err != nil {
		return nil, err
	}

	return lq, nil
}

func ExtractID(c *gin.Context, name string) (uuid.UUID, error) {
	ID, ok := c.Params.Get(name)
	if !ok {
		return uuid.Nil, go_kit.ErrParameterMissing.WithMessage(fmt.Sprintf("Path parameter %s is missing", name))
	}

	UID, err := uuid.Parse(ID)
	if err != nil {
		return uuid.Nil, go_kit.Wrap(go_kit.ErrInvalidUUID.WithMessage(fmt.Sprintf("%s is not a valid UUID", ID)), err)
	}

	return UID, nil
}

func GetUserContext(c *gin.Context) (*model.UserContext, error) {
	user, ok := c.Get("user")
	if !ok {
		return nil, go_kit.ErrUnauthorized
	}

	userCtx, ok := user.(*model.UserContext)
	if !ok {
		return nil, go_kit.ErrUnauthorized
	}

	return userCtx, nil
}
