package errors

type DomainError struct {
	Message string
}

func (e *DomainError) Error() string {
	return e.Message
}

func NewDomainError(message string) *DomainError {
	return &DomainError{
		Message: message,
	}
}

// calling: return NewDomainError("error messages")
// when return, we can call err.Error() to call Error() function (will get e.Message)
