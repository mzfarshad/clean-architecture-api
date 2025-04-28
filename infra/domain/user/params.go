package user

type CreateParams struct {
	Name     string
	Email    string
	Password string
	Type     Type
}

type SearchParams struct {
	Name  string
	Email string
	Limit int
	Page  int
}

type PaginationParams struct {
	TotalData  int
	TotalPages int
	Result     []*Entity
}
