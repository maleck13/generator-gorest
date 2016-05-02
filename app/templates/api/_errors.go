package api

import (
	"runtime"
)

//generic http error wrapper that implements the error interface
type HttpError interface {
	HttpErrorCode() int
	ErrorContext() string
	LineNumber() int
	SourceFile() string
}

type HttpHandlerError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Context string `json:"context"`
	Line    int    `json:"-"`
	File    string `json:"-"`
}

func NewHttpError(err error, code int) HttpError {
	_, f, n, _ := runtime.Caller(1)
	return &HttpHandlerError{Message: err.Error(), Code: code, Line: n, File: f}
}

func NewHttpErrorWithContext(err error, code int, context string) HttpError {
	_, f, n, _ := runtime.Caller(1)
	return &HttpHandlerError{Message: err.Error(), Code: code, Context: context, Line: n, File: f}
}

func (he *HttpHandlerError) Error() string {
	return he.Message
}

func (he *HttpHandlerError) ErrorContext() string {
	return he.Context
}

func (he *HttpHandlerError) HttpErrorCode() int {
	return he.Code
}

func (he *HttpHandlerError) LineNumber() int {
	return he.Line
}

func (he *HttpHandlerError) SourceFile() string {
	return he.File
}
