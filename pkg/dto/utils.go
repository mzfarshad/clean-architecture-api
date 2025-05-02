package dto

// Ptr returns a mapper function (T->*T) which returns the allocated address of a T type value.
func Ptr[T any]() func(T) *T {
	return func(t T) *T { return &t }
}

// IndirectFunc returns a mapper function (*T->T) which indirect a non-nil value of type *T to its value.
// In case of nil input and safe mode returns the zero value of T, else panics.
func IndirectFunc[T any](safe bool) func(*T) T {
	return func(t *T) T {
		if safe {
			if t == nil {
				var zero T
				return zero
			}
		}
		return *t
	}
}

// Address returns the address of an input value.
func Address[T any](input T) *T { return &input }

// Indirect makes indirect a non-nil value of type *T to its value.
// In case of nil input and safe mode returns the zero value of T, else panics.
func Indirect[T any](input *T, safe bool) T {
	if safe && input == nil {
		var zero T
		return zero
	}
	return *input
}

// Is reports weather any comparable type value equals the given targets of same type.
func Is[T comparable](t, target T, or ...T) bool {
	if t == target {
		return true
	}
	for i := range or {
		if t == or[i] {
			return true
		}
	}
	return false
}
