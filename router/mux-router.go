package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bifr0ns/academy-go-q32021/common"

	"github.com/gorilla/mux"
)

type muxRouter struct{}

var (
	muxDispatcher = mux.NewRouter()
)

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

func (*muxRouter) SERVE(port string) {
	fmt.Printf("Mux HTTP server running on port %v\n", port)
	muxDispatcher.NotFoundHandler = http.HandlerFunc(common.NotFoundHandler)

	log.Fatal(http.ListenAndServe(port, muxDispatcher))
}
