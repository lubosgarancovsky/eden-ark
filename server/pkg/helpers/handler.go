package helpers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"github.com/lubosgarancovsky/go-kit/filter"
	"github.com/lubosgarancovsky/go-kit/list"
	"github.com/lubosgarancovsky/go-kit/rsql"
	"github.com/lubosgarancovsky/go-kit/sort"
)

func CreateListingQuery(c *gin.Context, parser *rsql.Parser, filterMap map[string]string, sortMap map[string]string) (*list.ListingQuery, error) {
	var qp list.QueryParms
	err := c.ShouldBindQuery(&qp)
	if err != nil {
		return nil, api_err.Wrap(api_err.ErrBadRequest, err).WithMessage("Invalid query parameters")
	}

	var limit = 10
	if qp.PageSize > 0 {
		limit = qp.PageSize
	}

	var page = 1
	if qp.Page > 0 {
		page = qp.Page
	}

	lq := &list.ListingQuery{
		Filter: nil,
		Sort:   []sort.Sort{},
		Limit:  limit,
		Offset: (page - 1) * limit,
		Page:   page,
	}

	if qp.Filter != "" {
		ast, err := parser.Parse(qp.Filter)
		if err != nil {
			return nil, api_err.Wrap(api_err.ErrBadRequest, err).WithMessage("Invalid filter parameter")
		}

		fil, err := filter.BuildFilter(ast, filterMap)
		if err != nil {
			return nil, api_err.Wrap(api_err.ErrBadRequest, err).WithMessage("Invalid filter parameter")
		}

		lq.Filter = fil
	}

	if qp.Sort != "" {
		srt, err := sort.BuildSort(qp.Sort, sortMap)
		if err != nil {
			return nil, api_err.Wrap(api_err.ErrBadRequest, err).WithMessage("Invalid sort parameter")
		}

		lq.Sort = srt
	}

	return lq, nil
}

func ExtractID(c *gin.Context, name string) (uuid.UUID, error) {
	ID, ok := c.Params.Get(name)
	if !ok {
		return uuid.Nil, api_err.ErrParameterMissing.WithMessage(fmt.Sprintf("Path parameter %s is missing", name))
	}

	UID, err := uuid.Parse(ID)
	if err != nil {
		return uuid.Nil, api_err.Wrap(api_err.ErrInvalidUUID.WithMessage(fmt.Sprintf("%s is not a valid UUID", ID)), err)
	}

	return UID, nil
}
