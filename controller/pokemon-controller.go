package controller

import (
	"encoding/json"
	"net/http"

	"github.com/bifr0ns/academy-go-q32021/common"
	fmte "github.com/bifr0ns/academy-go-q32021/error"
	"github.com/bifr0ns/academy-go-q32021/model"

	"github.com/gorilla/mux"
)

type pokemonService interface {
	FindById(pokemonId string) (*model.Pokemon, error)
	SaveFromExternal(externalPokemon model.PokemonExternal) (*model.Pokemon, error)
}

type pokemonClient interface {
	GetExternalPokemon(uri string, id string) model.PokemonExternal
}

//PokemonController contains a resty.Client and an interface of pokemonService, which contains two methods.
//
//FindById recieves a string, and returns a model of Pokemon and error if any.
//
//SaveFromExternal recieves a model of PokemonExternal, and returns a model of Pokemon and error if any.
type PokemonController struct {
	service pokemonService
	client  pokemonClient
}

//NewPokemonController expects to recieve PokemonService and a restClient, returns an struct of PokemonController.
func NewPokemonController(service pokemonService, restClient pokemonClient) PokemonController {
	return PokemonController{service, restClient}
}

//GetPokemonById of type pokemonController, handles the request and calls the service to find a pokemon by the id given.
//The response is encoded in json, even if it is an error.
func (pc *PokemonController) GetPokemonById(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pokemonId := vars["pokemon_id"]

	pokemon, err := pc.service.FindById(pokemonId)

	rw.Header().Add("Content-Type", "application/json")

	if err != nil {
		if err.Error() == common.PokemonNotFound {
			rw.WriteHeader(http.StatusNotFound)
		} else {
			rw.WriteHeader(http.StatusInternalServerError)
		}
		json.NewEncoder(rw).Encode(fmte.FormattedError{Message: err.Error()})
		return
	}

	json.NewEncoder(rw).Encode(pokemon)
}

//GetExternalPokemonById of type pokemonController, handles the request to create a new PokemonExternal,
//calls the service to write a new pokemon in the csv file.
//The response is encoded in json. Can be successful of the pokemon can already exist in the csv file.
func (pc *PokemonController) GetExternalPokemonById(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pokemonId := vars["pokemon_id"]

	externalPokemon := pc.client.GetExternalPokemon(common.GetPokemonUri, pokemonId)

	rw.Header().Add("Content-Type", "application/json")

	pokemon, err := pc.service.SaveFromExternal(externalPokemon)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(fmte.FormattedError{Message: err.Error()})
		return
	}

	json.NewEncoder(rw).Encode(pokemon)
}
