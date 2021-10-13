package router

import "net/http"

//GET a uri string and a function with ResponseWriter and Request, will dispatch Get request.
//
//POST a uri string and a function with ResponseWriter and Request, will dispatch Post request.
//
//SERVE recieves a port and expects to load and run the server.
type Router interface {
	GET(uri string, f func(w http.ResponseWriter, r *http.Request))
	POST(uri string, f func(w http.ResponseWriter, r *http.Request))
	SERVE(port string)
}
