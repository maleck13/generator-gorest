package api

import (
	"encoding/json"
	"net/http"
	"github.com/Sirupsen/logrus"
)

//Wraps route handlers so that if there is an error returned we don't need to duplicate the error handling
func RouteErrorHandler(handler HttpHandler) http.HandlerFunc {

	return func(wr http.ResponseWriter, req *http.Request) {
		encoder := json.NewEncoder(wr)
		//may change to use a context object containing other data
		if err := handler(wr, req); err != nil {
			logrus.Error("handler error: ", err)
			wr.WriteHeader(err.HttpErrorCode())
			encoder.Encode(err)
			return
		}

	}
}

