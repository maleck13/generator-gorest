package middleware

import (
	"github.com/Sirupsen/logrus"
	"net/http"
)

type ReqLog struct {
	Method string
	Url string

}
//example middle ware that logs incoming requests
func ExampleMiddleware(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc){

	logrus.Info("request: ", &ReqLog{
		req.Method,
		req.URL.Path,
	})
	next(rw,req)

}
