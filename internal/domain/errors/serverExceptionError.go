package errors

type ServerException struct {
	err error
}

func (e *ServerException) Error() string {
	return e.err.Error()
}

func NewServerException(err error) *ServerException {
	return &ServerException{
		err: err,
	}
}
