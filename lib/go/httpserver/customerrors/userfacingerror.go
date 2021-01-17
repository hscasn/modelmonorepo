package customerrors

// UserFacingError is an error with additional information for APIs. The contents of the
// UserMessage are sent back to the caller, and not hidden
type UserFacingError interface {
	error
	StatusCode() int
	UserMessage() string
}

// GenericUserFacingError is a generic error
type GenericUserFacingError struct {
	error
	GenericErrorStatusCode  int
	GenericErrorUserMessage string
}

// StatusCode for error
func (e GenericUserFacingError) StatusCode() int {
	return e.GenericErrorStatusCode
}

// UserMessage for error
func (e GenericUserFacingError) UserMessage() string {
	return e.GenericErrorUserMessage
}
