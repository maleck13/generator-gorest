package api

//generic http error wrapper that implements the error interface
type HttpError interface {
	HttpErrorCode() int
	ErrorContext()string
}

type HttpHandlerError struct {
	Code    int `json:"code"`
	Message string `json:"message"`
	Context string `json:"context"`
}

func NewHttpError(err error, code int) HttpError {
	return &HttpHandlerError{Message: err.Error(), Code: code}
}

func NewHttpErrorWithContext(err error, code int, context string) HttpError{
	return &HttpHandlerError{Message:err.Error(),Code:code, Context:context }
}

func (he *HttpHandlerError) Error() string {
	return he.Message
}

func (he *HttpHandlerError) ErrorContext()string  {
	return he.Context
}

func (he *HttpHandlerError) HttpErrorCode() int {
	return he.Code
}
