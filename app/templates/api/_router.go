package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/codegangsta/negroni"
	"<%=basePackage %>/<%=baseName %>/api/middleware"
	<% if("yes" === metrics) { %>
        "github.com/prometheus/client_golang/prometheus"
	<% } %>
)

type HttpHandler func(wr http.ResponseWriter, req *http.Request) HttpError

func NewRouter() http.Handler {
	r := mux.NewRouter().StrictSlash(true)

	//dedicated to /api which will use the PathPrefix
	apiRouter := mux.NewRouter().PathPrefix("/api").Subrouter().StrictSlash(false)

	r.PathPrefix("/api").Handler(negroni.New(
		negroni.HandlerFunc(middleware.ExampleMiddleware),
		negroni.Wrap(apiRouter),
	))

	<% if("yes" === metrics) { %>
        apiRouter.HandleFunc("/", prometheus.InstrumentHandlerFunc("/api/",RouteErrorHandler(IndexHandler))).Methods("GET")
        r.Handle("/metrics",prometheus.Handler())
	<% }else{ %>

        apiRouter.HandleFunc("/", RouteErrorHandler(IndexHandler)).Methods("GET")

        <% } %>

	r.HandleFunc("/health",RouteErrorHandler(HealthHandler))
	r.HandleFunc("/ping",RouteErrorHandler(Ping))

	return r
}


