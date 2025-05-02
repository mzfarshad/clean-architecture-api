package dto

type optionalTypes interface {
	~bool | ~uint | ~string
}

// NewOptional creates a new Optional with the given value and marks it as populated.
// If you want an unpopulated Optional value, use Optional[T]{}.
func NewOptional[T optionalTypes](value T) Optional[T] {
	return Optional[T]{value: value, populated: true}
}

type Optional[T optionalTypes] struct {
	value     T
	populated bool
}

// IsPopulated returns true if the value of the Optional is populated.
func (t *Optional[T]) IsPopulated() bool {
	return t.populated
}

// Value returns the value of the Optional type.
// If the value is not populated, it will return the zero value of the type.
// Check if the value is populated using IsPopulated.
func (t *Optional[T]) Value() T { return t.value }
