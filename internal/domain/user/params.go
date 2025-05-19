package user

type CreateParams struct {
	Name     string
	Email    string
	Password string
	Type     Type
}

type SearchParams struct {
	Name  string `query:"name"`
	Email string `query:"email"`
	Limit int    `query:"page_size"`
	Page  int    `query:"page"`
}

type PaginationParams struct {
	TotalData  int
	TotalPages int
	Result     []*Entity
}
