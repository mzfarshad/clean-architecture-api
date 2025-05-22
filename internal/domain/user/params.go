package user

type PaginationParams struct {
	TotalData  int
	TotalPages int
	Result     []*Entity
}
