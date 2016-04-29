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
	//dedicated sub route for /api which will use the PathPrefix
	apiRouter := mux.NewRouter().PathPrefix("/api").Subrouter().StrictSlash(true)

	//recovery middleware for any panics in the handlers
	recovery := negroni.NewRecovery()
	recovery.PrintStack = false
	//add middleware for all routes
	n := negroni.New(recovery)
	//add some top level routes
<% if("yes" === metrics) { %>
	r.Handle("/metrics",prometheus.Handler())
<% } %>
	r.HandleFunc("/sys/info/health",RouteErrorHandler(HealthHandler))
	r.HandleFunc("/sys/info/ping",RouteErrorHandler(Ping))


	r.PathPrefix("/api").Handler(negroni.New(
		negroni.HandlerFunc(middleware.ExampleMiddleware),
		negroni.Wrap(apiRouter),
	))

	<% if("yes" === metrics) { %>
        apiRouter.HandleFunc("/", prometheus.InstrumentHandlerFunc("/api/",RouteErrorHandler(IndexHandler))).Methods("GET")
	<% if("mongo" === database) { %>
	apiRouter.HandleFunc("/mongo", prometheus.InstrumentHandlerFunc("/api/mongo",RouteErrorHandler(IndexMongo))).Methods("POST")
	<% } %>
	<% if("yes" == messaging) { %>
	apiRouter.HandleFunc("/stomp",prometheus.InstrumentHandlerFunc("/api/stomp",RouteErrorHandler(IndexStomp))).Methods("GET","POST")
	<% } %>
        r.Handle("/metrics",prometheus.Handler())
	<% }else{ %>

        apiRouter.HandleFunc("/", RouteErrorHandler(IndexHandler)).Methods("GET")
	<% if("mongo" === database) { %>
	apiRouter.HandleFunc("/mongo", RouteErrorHandler(IndexMongo)).Methods("POST")
	<% } %>
	<% if("yes" == messaging) { %>
	apiRouter.HandleFunc("/stomp",RouteErrorHandler(IndexStomp)).Methods("GET","POST")
	<% } %>

        <% } %>
	//wire up middleware and router
	n.UseHandler(r)

	return n  //negroni implements the http.Handler interface
}


