package helpers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/go-kit/filter"
	"github.com/lubosgarancovsky/go-kit/list"
	"github.com/lubosgarancovsky/go-kit/rsql"
	"github.com/lubosgarancovsky/go-kit/sort"
)

func CreateListingQuery(c *gin.Context, parser *rsql.Parser, filterMap map[string]string, sortMap map[string]string) (*list.ListingQuery, error) {
	var qp list.QueryParms
	err := c.ShouldBindQuery(&qp)
	if err != nil {
		return nil, err
	}

	var limit int = 10
	if qp.PageSize > 0 {
		limit = qp.PageSize
	}

	var page int = 1
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
			return nil, err
		}

		fil, err := filter.BuildFilter(ast, filterMap)
		if err != nil {
			return nil, err
		}

		lq.Filter = fil
	}

	if qp.Sort != "" {
		srt, err := sort.BuildSort(qp.Sort, sortMap)
		if err != nil {
			return nil, err
		}

		lq.Sort = srt
	}

	return lq, nil
}

func ExtractID(c *gin.Context, name string) (uuid.UUID, error) {
	ID, ok := c.Params.Get("clientId")
	if !ok {
		return uuid.Nil, fmt.Errorf("path parameter %s is missing", name)
	}

	UID, err := uuid.Parse(ID)
	if err != nil {
		return uuid.Nil, err
	}

	return UID, nil
}
