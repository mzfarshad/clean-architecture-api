package user

// Type is the user Type includes TypeCustomer, TypeAdmin.
type Type string

const (
	TypeAdmin    Type = "admin"
	TypeCustomer Type = "customer"
)
