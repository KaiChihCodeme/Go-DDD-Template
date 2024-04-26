package errors

type ApiServiceException struct {
	err    error
	source string
}

func (e *ApiServiceException) Error() string {
	return e.err.Error()
}

func NewApiException(err error, source string) *ApiServiceException {
	return &ApiServiceException{
		err:    err,
		source: source,
	}
}
