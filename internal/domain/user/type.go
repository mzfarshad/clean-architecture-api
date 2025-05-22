package user

// Type is the user Type includes TypeCustomer, TypeAdmin.
type Type string

const (
	TypeAdmin    Type = "admin"
	TypeCustomer Type = "customer"
)

func (x Type) Is(target Type, or ...Type) bool {
	if x == target {
		return true
	}
	for i := range or {
		if x == or[i] {
			return true
		}
	}
	return false
}

func (x Type) String() string { return string(x) }
