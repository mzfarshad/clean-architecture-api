package search

func NewPagination[T any](size, page int) *Pagination[T] {
	return &Pagination[T]{size: size, page: page}
}

type Pagination[T any] struct {
	// inputs
	Query T
	Param T
	size  int
	page  int
	// results
	total int64
}

func (p *Pagination[T]) WithTotal(count int64) *Pagination[T] {
	p.total = count
	return p
}

func (p *Pagination[T]) Limit() int { return p.size }
func (p *Pagination[T]) Offset() int {
	if p.size < 0 {
		return -1
	}
	return (p.page - 1) * p.size
}

/*
 To be an implementation of rest.Pagination
*/

func (p *Pagination[T]) Size() int    { return p.size }
func (p *Pagination[T]) Page() int    { return p.page }
func (p *Pagination[T]) Total() int64 { return p.total }
func (p *Pagination[T]) Filters() any { return p.Query }
