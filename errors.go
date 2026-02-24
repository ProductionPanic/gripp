package gripp

// ErrMissingApiKey
var ErrMissingApiKey = &GrippError{
	Message: "API key is required",
}

// GrippError is a custom error type for the gripp package
type GrippError struct {
	Message string
}

func (e *GrippError) Error() string {
	return e.Message
}
