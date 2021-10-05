package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bifr0ns/academy-go-q32021/common"
	"github.com/bifr0ns/academy-go-q32021/error"
	"github.com/bifr0ns/academy-go-q32021/model"
	"github.com/bifr0ns/academy-go-q32021/service"

	"github.com/go-resty/resty/v2"
	"github.com/gorilla/mux"
)

//GetPokemonById recieves a ResponseWriter and Request, and will encode a json as response.
//
//GetExternalPokemonById recieves a ResponseWriter and Request, and will encode a json as response.
type PokemonController interface {
	GetPokemonById(http.ResponseWriter, *http.Request)
	GetExternalPokemonById(http.ResponseWriter, *http.Request)
}

type pokemonController struct{}

//NewPokemonController expects to recieve PokemonService and a restClient, returns an interface of PokemonController.
func NewPokemonController(service service.PokemonService, restClient *resty.Client) PokemonController {
	pokemonService = service
	client = restClient
	return &pokemonController{}
}

var (
	pokemonService service.PokemonService
	client         *resty.Client
)

//GetPokemonById of type pokemonController, handles the request and calls the service to find a pokemon by the id given.
//The response is encoded in json, even if it is an error.
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
		json.NewEncoder(rw).Encode(error.FormattedError{Message: err.Error()})
		return
	}

	json.NewEncoder(rw).Encode(pokemon)
}

//GetExternalPokemonById of type pokemonController, handles the request to create a new PokemonExternal,
//calls the service to write a new pokemon in the csv file.
//The response is encoded in json. Can be successful of the pokemon can already exist in the csv file.
func (*pokemonController) GetExternalPokemonById(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pokemonId := vars["pokemon_id"]

	resp, _ := client.R().
		SetPathParams(map[string]string{
			"pokemonId": fmt.Sprint(pokemonId),
		}).
		SetResult(model.PokemonExternal{}).
		SetHeader("Accept", "application/json").
		Get("https://pokeapi.co/api/v2/pokemon/{pokemonId}")

	pokemon := *resp.Result().(*model.PokemonExternal)

	rw.Header().Add("Content-Type", "application/json")

	externalPokemon, err := pokemonService.SaveFromExternal(pokemon)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(error.FormattedError{Message: err.Error()})
		return
	}

	json.NewEncoder(rw).Encode(externalPokemon)
}
