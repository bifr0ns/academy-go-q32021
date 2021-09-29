package app

import (
	"log"
	"net/http"

	"github.com/bifr0ns/academy-go-q32021/common"
	"github.com/bifr0ns/academy-go-q32021/routes"

	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()

	routes.Setup(router)

	log.Fatal(http.ListenAndServe(common.LocalHost+":"+common.LocalPort, router))
}
