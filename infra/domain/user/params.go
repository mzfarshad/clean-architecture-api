package user

type CreateParams struct {
	Email    string
	Password string
	Type     Type
}

type SearchParams struct {
	Name  string
	Email string
}
