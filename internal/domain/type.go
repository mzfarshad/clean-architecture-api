package domain

// UserType is the user type includes Customer, Admin.
type UserType string

const (
	Admin    UserType = "admin"
	Customer UserType = "customer"
)

func (x UserType) Is(target UserType, or ...UserType) bool {
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

func (x UserType) String() string { return string(x) }
