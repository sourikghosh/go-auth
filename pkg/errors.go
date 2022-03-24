package pkg

// Error represents a handler error. It provides methods for a HTTP status
// code and embeds the built-in error interface.
type Error interface {
	error
	Status() int
}

// AuthErr represents an error with an associated HTTP status code.
type AuthErr struct {
	Code int
	Err  error
}

// Error allows AuthErr to satisfy the error interface.
func (e AuthErr) Error() string {
	return e.Err.Error()
}

// Status returns the HTTP status code.
func (e AuthErr) Status() int {
	return e.Code
}
