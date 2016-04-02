package middleware

import (
	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"net/http"
)

const (
	_HEADER_ALLOW_HEADERS     = "Access-Control-Allow-Headers"
	_HEADER_ALLOW_ORIGIN      = "Access-Control-Allow-Origin"
	_HEADER_ALLOW_CREDENTIALS = "Access-Control-Allow-Credentials"
	_HEADER_ALLOW_METHODS     = "Access-Control-Allow-Methods"
)

type CorsOpts struct {
	AllowedHeaders   []string
	AllowedOrigins   []string
	AllowedMethods   []string
	AllowCredentials bool
	Debug            bool
}

func (opts CorsOpts) addAllowedHeaders(headers http.Header) {
	if nil == headers {
		headers = http.Header{}
	}
	for _, v := range opts.AllowedHeaders {
		headers.Add(_HEADER_ALLOW_HEADERS, v)
	}
}

func (opts CorsOpts) addAllowOrigin(headers http.Header) {

	for _, v := range opts.AllowedOrigins {
		headers.Add(_HEADER_ALLOW_ORIGIN, v)
	}

}

func (opts CorsOpts) addAllowCredentials(headers http.Header) {
	if opts.AllowCredentials {
		headers.Add(_HEADER_ALLOW_CREDENTIALS, "true")
	}
}

func (opts CorsOpts) addAllowMethods(headers http.Header) {

	for _, v := range opts.AllowedMethods {
		headers.Add(_HEADER_ALLOW_METHODS, v)
	}
}

func DefaultCorsOpts(debug bool) CorsOpts {
	opts := CorsOpts{
		AllowedHeaders:   []string{"User-Agent", "Origin", "Accept", "Content-Type"},
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	}
	if debug {
		opts.Debug = true
	}
	return opts
}

func Cors(opts CorsOpts) negroni.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
		if req.Method == "OPTIONS" {
			responseHeaders := rw.Header()
			opts.addAllowOrigin(responseHeaders)
			opts.addAllowCredentials(responseHeaders)
			opts.addAllowedHeaders(responseHeaders)
			opts.addAllowMethods(responseHeaders)
			return
		}
		if !isOriginAllowed(opts, req.Header.Get("Origin")) {
			if opts.Debug {
				logrus.Info("cors: ", "failed origin ", req.Header.Get("Origin"))
			}
			return
		}
		if !isMethodAllowed(opts, req.Method) {
			if opts.Debug {
				logrus.Info("cors: ", "failed method allowed ", req.Method)
			}
			return
		}
		if !areHeadersAllowed(opts, req.Header) {
			if opts.Debug {
				logrus.Info("cors: ", "failed headers allowed ", req.Header)
			}
			return
		}
		next(rw, req)

	}
}

func isOriginAllowed(opts CorsOpts, origin string) bool {

	for _, opt := range opts.AllowedOrigins {
		if opt == "*" {
			return true
		}
		if opt == origin {
			return true
		}
	}
	return false
}

func isMethodAllowed(opts CorsOpts, method string) bool {
	for _, m := range opts.AllowedMethods {
		if m == method {
			return true
		}
	}
	return false
}

func areHeadersAllowed(opts CorsOpts, headers http.Header) bool {
	if opts.Debug {
		logrus.Info("allowed headers ", opts.AllowedHeaders, " passed headers ", headers)
	}
	for k, _ := range headers {
		header := http.CanonicalHeaderKey(k)
		found := false
		if opts.Debug {
			logrus.Info("looking for header ", header)
		}
		for _, allowed := range opts.AllowedHeaders {
			allowed = http.CanonicalHeaderKey(allowed)
			if opts.Debug {
				logrus.Info("allowed header ", allowed, " checking header ", header, allowed == header)
			}

			found = allowed == header
			if found {
				break
			}
		}
		if !found {
			return found
		}
	}
	return true

}
