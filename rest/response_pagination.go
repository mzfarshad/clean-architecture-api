package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mzfarshad/music_store_api/pkg/search"
	"strconv"
)

const (
	DefaultPageSize   uint32 = 50
	DefaultPageNumber uint32 = 1
)

func NewPagination[T any](ctx *fiber.Ctx) (*search.Pagination[T], error) {
	pageSize, err := strconv.Atoi(ctx.Query("page_size"))
	if err != nil {
		pageSize = int(DefaultPageSize)
	}
	pageNumber, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		pageNumber = int(DefaultPageNumber)
	}

	p := search.NewPagination[T](pageSize, pageNumber)
	if err = ctx.QueryParser(&p.Query); err != nil {
		return nil, err
	}

	return p, nil
}

type Pagination interface {
	Size() int
	Page() int
	Total() int64
	Filters() any
}

type pagination struct {
	PageSize   int   `json:"page_size"`
	Page       int   `json:"page"`
	TotalCount int64 `json:"total_count"`
	Filters    any   `json:"filters"`
}
