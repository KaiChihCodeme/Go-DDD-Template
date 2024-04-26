package errors

type NotFoundError struct {
	err    error
	source string
}

func (e *NotFoundError) Error() string {
	return e.err.Error()
}

func NewNotFoundError(err error, source string) *NotFoundError {
	return &NotFoundError{
		err:    err,
		source: source,
	}
}
