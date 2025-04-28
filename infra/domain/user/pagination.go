package user

type ResponsePagination struct {
	TotalData  int
	TotalPages int
	Result     []*Entity
}
