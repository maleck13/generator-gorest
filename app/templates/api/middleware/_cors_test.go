package middleware_test

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"<%=basePackage %>/<%=baseName %>/api/middleware"
	"github.com/codegangsta/negroni"
	"github.com/stretchr/testify/assert"
	"fmt"
)

const(
	_TEST_ACCESS_ORIGIN = "Access-Control-Allow-Origin"
	_TEST_ACCESS_HEADERS = "Access-Control-Allow-Headers"
	_TEST_ACCESS_METHODS = "Access-Control-Allow-Methods"
	_TEST_ACCESS_CREDENTIALS = "Access-Control-Allow-Credentials"
)

var testHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bar"))
})

func TestDefault(t *testing.T){
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("OPTIONS", "http://example.com/foo", nil)
	req.Header.Add("Origin", "http://foobar.com")
	corsOpts := middleware.DefaultCorsOpts(true)
	//add a cors middleware
	cors := middleware.Cors(corsOpts)
	handler := negroni.HandlerFunc(cors)
	handler.ServeHTTP(res,req, testHandler)
	assert.Equal(t,res.Header().Get(_TEST_ACCESS_ORIGIN),"*","expected match headers")
	accessHeaders := []string{"User-Agent","Origin","Accept","Content-Type"}
	assert.Equal(t,accessHeaders,res.HeaderMap[http.CanonicalHeaderKey(_TEST_ACCESS_HEADERS)],"expected same set of headers")
	allowMethods := []string{"GET","POST","PUT","DELETE"}
	assert.Equal(t,allowMethods,res.HeaderMap[http.CanonicalHeaderKey(_TEST_ACCESS_METHODS)],"expected same set of headers")
	assert.Equal(t,"true",res.Header().Get(_TEST_ACCESS_CREDENTIALS),"expected credentials header to be true")
}

func TestAddCustomHeader(t *testing.T){
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("OPTIONS", "http://example.com/foo", nil)
	req.Header.Add("Origin", "http://foobar.com")
	corsOpts := middleware.DefaultCorsOpts(true)
	//add a cors middleware
	corsOpts.AllowedHeaders = append(corsOpts.AllowedHeaders,"x-test-header")
	cors := middleware.Cors(corsOpts)
	handler := negroni.HandlerFunc(cors)
	handler.ServeHTTP(res,req, testHandler)
	assert.Equal(t,res.Header().Get(_TEST_ACCESS_ORIGIN),"*","expected match headers")
	accessHeaders := []string{"User-Agent","Origin","Accept","Content-Type","x-test-header"}
	assert.Equal(t,accessHeaders,res.HeaderMap[http.CanonicalHeaderKey(_TEST_ACCESS_HEADERS)],"expected same set of headers")
	allowMethods := []string{"GET","POST","PUT","DELETE"}
	assert.Equal(t,allowMethods,res.HeaderMap[http.CanonicalHeaderKey(_TEST_ACCESS_METHODS)],"expected same set of headers")
	assert.Equal(t,"true",res.Header().Get(_TEST_ACCESS_CREDENTIALS),"expected credentials header to be true")
	fmt.Print(string(res.Body.Bytes()))
}

func TestGetCustomHeader(t *testing.T){
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://example.com/foo", nil)
	req.Header.Add("Origin", "http://foobar.com")
	req.Header.Add("x-test-header","test")
	corsOpts := middleware.DefaultCorsOpts(true)
	//add a cors middleware
	corsOpts.AllowedHeaders = append(corsOpts.AllowedHeaders,"x-test-header")
	cors := middleware.Cors(corsOpts)
	handler := negroni.HandlerFunc(cors)
	handler.ServeHTTP(res,req, testHandler)
	//should get through assert body ok
	bod := string(res.Body.Bytes())
	assert.Equal(t,bod,"bar","expected a body")
}

func TestGetFails(t *testing.T){
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://example.com/foo", nil)
	req.Header.Add("Origin", "http://foobar.com")
	req.Header.Add("x-test-header","test")
	corsOpts := middleware.DefaultCorsOpts(true)
	//add a cors middleware
	cors := middleware.Cors(corsOpts)
	handler := negroni.HandlerFunc(cors)
	handler.ServeHTTP(res,req, testHandler)
	//should get through assert body ok
	bod := string(res.Body.Bytes())
	assert.Equal(t,bod,"","expected an empty body")
}

//todo expand tests