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
