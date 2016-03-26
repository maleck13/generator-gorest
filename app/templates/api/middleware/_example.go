package middleware

import (
	"github.com/Sirupsen/logrus"
	"net/http"
)

func ExampleMiddleware(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc){
	logrus.Info("example middleware")
	next(rw,req)
}
