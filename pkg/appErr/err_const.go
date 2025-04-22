package apperr

// type error
const (
	TypeApi        = "API"
	TypeInternal   = "INTERNAL"
	TypeValidation = "VALIDATION"
	TypeDatabase   = "DATABASE"
	TypeConfig     = "CONFIG"
)

// Status codes
const (
	StatusOK                  = 200
	StatusBadRequest          = 400
	StatusUnauthorized        = 401
	StatusForbidden           = 403
	StatusNotFound            = 404
	StatusInternalServerError = 500
)

const (
	ErrRecordNotFound = "record not found"
)
