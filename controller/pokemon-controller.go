package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/bifr0ns/academy-go-q32021/common"
	fmte "github.com/bifr0ns/academy-go-q32021/error"
	"github.com/bifr0ns/academy-go-q32021/model"

	"github.com/gorilla/mux"
)

type pokemonService interface {
	FindById(pokemonId string) (*model.Pokemon, error)
	SaveFromExternal(externalPokemon model.PokemonExternal) (*model.Pokemon, error)
	GetPokemons(dataType string, items int, items_per_workers int, workers int) ([]model.Pokemon, error)
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

	rw.Header().Add("Content-Type", "application/json")

	_, err := strconv.Atoi(pokemonId)
	if err != nil && pokemonId != "" {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(fmte.FormattedError{Message: common.InvalidParameters})
		return
	}

	pokemon, err := pc.service.FindById(pokemonId)

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

	rw.Header().Add("Content-Type", "application/json")

	_, err := strconv.Atoi(pokemonId)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(fmte.FormattedError{Message: common.InvalidParameters})
		return
	}

	externalPokemon := pc.client.GetExternalPokemon(common.GetPokemonUri, pokemonId)
	if externalPokemon.Id == 0 {
		rw.WriteHeader(http.StatusNotFound)
		json.NewEncoder(rw).Encode(fmte.FormattedError{Message: common.PokemonNotFound})
		return
	}

	pokemon, err := pc.service.SaveFromExternal(externalPokemon)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(fmte.FormattedError{Message: err.Error()})
		return
	}

	json.NewEncoder(rw).Encode(pokemon)
}

//GetPokemonsByWorker of type pokemonController, handles the request to create an array of pokemons,
//calls the service to write pokemons based on the query
//The response is encoded in json.
func (pc *PokemonController) GetPokemonsByWorker(rw http.ResponseWriter, r *http.Request) {

	var types string
	var items string
	var items_per_workers string
	var workers string

	types = r.FormValue("type")
	items = r.FormValue("items")
	items_per_workers = r.FormValue("items_per_workers")
	workers = r.FormValue("workers")

	urlParams := r.URL.Query()

	rw.Header().Add("Content-Type", "application/json")

	for i := range urlParams {
		if i != "type" && i != "items" && i != "items_per_workers" && i != "workers" {
			rw.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(rw).Encode(fmte.FormattedError{Message: common.InvalidParameters})
			return
		}
	}

	if types == "" {
		types = "all"
	}

	if items == "" {
		items = "-1"
	}

	if items_per_workers == "" {
		items_per_workers = "-1"
	}

	if workers == "" {
		workers = "-1"
	}

	dataType := strings.ToLower(types)
	if dataType != "odd" && dataType != "even" && dataType != "all" {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(fmte.FormattedError{Message: common.InvalidParameters})
		return
	}

	items_int, err := strconv.Atoi(items)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(fmte.FormattedError{Message: common.InvalidParameters})
		return
	}

	items_per_workers_int, err := strconv.Atoi(items_per_workers)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(fmte.FormattedError{Message: common.InvalidParameters})
		return
	}

	workers_int, err := strconv.Atoi(workers)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(fmte.FormattedError{Message: common.InvalidParameters})
		return
	}
	if workers_int > 8 {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(fmte.FormattedError{Message: common.InvalidWorkers})
		return
	}

	pokemons, err := pc.service.GetPokemons(types, items_int, items_per_workers_int, workers_int)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(fmte.FormattedError{Message: err.Error()})
		return
	}

	json.NewEncoder(rw).Encode(pokemons)
}
