package errors

type BadRequestError struct {
	err error
}

func (e *BadRequestError) Error() string {
	return e.err.Error()
}

func NewBadRequestError(err error) *BadRequestError {
	return &BadRequestError{
		err: err,
	}
}
