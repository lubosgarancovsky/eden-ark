package helpers

import (
	"fmt"
	"sync"

	"github.com/lubosgarancovsky/go-kit/list"
	"gorm.io/gorm"
)

type ListState[T any] struct {
	items      []T
	totalCount int64
	wg         sync.WaitGroup
	errList    chan error
	lq         *list.ListingQuery
}

func OrderBy[T any](query *gorm.DB, state *ListState[T]) *gorm.DB {
	for _, clause := range state.lq.Sort {
		query = query.Order(fmt.Sprintf("%s %s", clause.Field, clause.Direction))
	}
	return query
}

func Paginate[T any](query *gorm.DB, state *ListState[T]) {
	defer state.wg.Done()
	q := OrderBy(query, state)
	if err := q.Limit(state.lq.Limit).Offset(state.lq.Offset).Find(&state.items).Error; err != nil {
		state.errList <- err
	}
}

func Count[T any](query *gorm.DB, state *ListState[T]) {
	defer state.wg.Done()
	if err := query.Count(&state.totalCount).Error; err != nil {
		state.errList <- err
	}
}

func List[T any](query *gorm.DB, lq *list.ListingQuery) ([]T, int64, error) {
	state := &ListState[T]{
		items:      make([]T, 0),
		totalCount: 0,
		wg:         sync.WaitGroup{},
		errList:    make(chan error, 2),
		lq:         lq,
	}

	state.wg.Add(2)
	go Paginate(query.Session(&gorm.Session{}), state)
	go Count(query.Session(&gorm.Session{}), state)

	state.wg.Wait()
	close(state.errList)

	for err := range state.errList {
		if err != nil {
			return nil, 0, err
		}
	}

	return state.items, state.totalCount, nil
}
