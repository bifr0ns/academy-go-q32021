package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bifr0ns/academy-go-q32021/error"

	"github.com/gorilla/mux"
)

type muxRouter struct{}

var (
	muxDispatcher = mux.NewRouter()
)

//NewMuxRouter returns an interface of Router. Which contains the methods:
//
//GET(uri string, f func(w http.ResponseWriter, r *http.Request)).
//POST(uri string, f func(w http.ResponseWriter, r *http.Request)).
//SERVE(port string).
func NewMuxRouter() Router {
	return &muxRouter{}
}

//GET of type muxRouter is expected to recieve a uri and a function with ResponseWrite and Request.
//Registers a new route with a matcher for the URL path.
func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

//GET of type muxRouter is expected to recieve a uri and a function with ResponseWrite and Request.
//Registers a new route with a matcher for the URL path.
func (*muxRouter) GETWITHQUERY(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
	muxDispatcher.HandleFunc(uri, f).Queries("type", "{type}", "items", "{items}", "items_per_workers", "{items_per_workers}").Methods("GET")
	muxDispatcher.HandleFunc(uri, f).Queries("items", "{items}", "items_per_workers", "{items_per_workers}").Methods("GET")
	muxDispatcher.HandleFunc(uri, f).Queries("items_per_workers", "{items_per_workers}").Methods("GET")
	muxDispatcher.HandleFunc(uri, f).Queries("type", "{type}", "items", "{items}").Methods("GET")
	muxDispatcher.HandleFunc(uri, f).Queries("type", "{type}").Methods("GET")
	muxDispatcher.HandleFunc(uri, f).Queries("type", "{type}", "items_per_workers", "{items_per_workers}").Methods("GET")
}

//POST of type muxRouter is expected to recieve a uri and a function with ResponseWrite and Request.
//Registers a new route with a matcher for the URL path.
func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

//SERVE of type muxRouter is expected to recieve a port. Will load the server and run it.
//Also assigns a function to the NotFoundHandler for every route that does not exists.
func (*muxRouter) SERVE(port string) {
	fmt.Printf("Mux HTTP server running on port %v\n", port)
	muxDispatcher.NotFoundHandler = http.HandlerFunc(error.NotFoundHandler)

	log.Fatal(http.ListenAndServe(":"+port, muxDispatcher))
}
