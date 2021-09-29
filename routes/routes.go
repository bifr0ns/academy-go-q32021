package routes

import (
	"net/http"

	"github.com/bifr0ns/academy-go-q32021/common"
	"github.com/bifr0ns/academy-go-q32021/controllers"

	"github.com/gorilla/mux"
)

func Setup(router *mux.Router) {

	router.HandleFunc("/pokemons/{pokemon_id:[0-9]+}", controllers.GetPokemonById).Methods(http.MethodGet)
	router.NotFoundHandler = http.HandlerFunc(common.NotFoundHandler)
}
