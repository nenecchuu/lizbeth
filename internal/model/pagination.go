package model

import "fmt"

var (
	defaultLimit uint64 = 10
)

type PaginationOpts struct {
	Limit  uint64
	Offset uint64
}

type Pagination struct {
	Limit  uint64
	Offset uint64
}

func NewPagination(o PaginationOpts) *Pagination {
	if o.Limit == 0 {
		o.Limit = defaultLimit
	}

	return &Pagination{
		Limit:  o.Limit,
		Offset: o.Offset,
	}
}

func (pg *Pagination) BuildPaginationQuery() string {
	if pg.Limit > 0 {
		return fmt.Sprintf(" LIMIT %d OFFSET %d", pg.Limit, pg.Offset)
	}

	return ""
}
