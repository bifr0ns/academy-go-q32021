package controller

import (
	"encoding/json"
	"net/http"

	"github.com/bifr0ns/academy-go-q32021/common"
	"github.com/bifr0ns/academy-go-q32021/service"

	"github.com/gorilla/mux"
)

type PokemonController interface {
	GetPokemonById(http.ResponseWriter, *http.Request)
}

type pokemonController struct{}

func NewPokemonController(service service.PokemonService) PokemonController {
	pokemonService = service
	return &pokemonController{}
}

var (
	pokemonService service.PokemonService
)

func (*pokemonController) GetPokemonById(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pokemonId := vars["pokemon_id"]

	pokemon, err := pokemonService.FindById(pokemonId)

	rw.Header().Add("Content-Type", "application/json")

	if err != nil {
		if err.Error() == common.PokemonNotFound {
			rw.WriteHeader(http.StatusNotFound)
		} else {
			rw.WriteHeader(http.StatusInternalServerError)
		}
		json.NewEncoder(rw).Encode(err.Error())
		return
	}

	json.NewEncoder(rw).Encode(pokemon)
}
