package api

//generic http error wrapper
type HttpError interface {
	HttpErrorCode() int
}

type HttpHandlerError struct {
	Code    int
	Message string
}

func NewHttpError(err error, code int) *HttpHandlerError {
	return &HttpHandlerError{Message: err.Error(), Code: code}
}

func (he *HttpHandlerError) Error() string {
	return he.Message
}

func (he *HttpHandlerError) HttpErrorCode() int {
	return he.Code
}
