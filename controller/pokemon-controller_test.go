package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bifr0ns/academy-go-q32021/common"
	"github.com/bifr0ns/academy-go-q32021/model"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var pokemon = model.Pokemon{
	Id:           155,
	Name:         "Cyndaquil",
	Type1:        "Fire",
	Type2:        "",
	Total:        309,
	HP:           39,
	Attack:       52,
	Defense:      43,
	SpeedAttack:  60,
	SpeedDefense: 50,
	Speed:        65,
	Generation:   2,
	Legendary:    "False",
}

var pokemonFromExternal = model.Pokemon{
	Id:           888,
	Name:         "Zacian-Hero",
	Type1:        "Fairy",
	Type2:        "",
	Total:        670,
	HP:           92,
	Attack:       130,
	Defense:      115,
	SpeedAttack:  80,
	SpeedDefense: 115,
	Speed:        138,
	Generation:   8,
	Legendary:    "False",
}

var externalPokemon = model.PokemonExternal{
	Id:           888,
	Name:         "zacian-hero",
	NotLegendary: true,
	Types: []model.PokemonTypes{
		{Slot: 1,
			Type: map[string]string{"name": "psychic"}},
	},
	Stats: []model.PokemonStats{
		{BaseStat: 97,
			Stat: map[string]string{"name": "hp"}},
		{BaseStat: 107,
			Stat: map[string]string{"name": "attack"}},
		{BaseStat: 101,
			Stat: map[string]string{"name": "defense"}},
		{BaseStat: 127,
			Stat: map[string]string{"name": "special-attack"}},
		{BaseStat: 89,
			Stat: map[string]string{"name": "special-defense"}},
		{BaseStat: 97,
			Stat: map[string]string{"name": "speed"}},
	},
}

var externalInvalidPokemon = model.PokemonExternal{
	Id: 0,
}

var pokemonJsonResponse = string("{\"id\":155,\"name\":\"Cyndaquil\",\"type_1\":\"Fire\",\"type_2\":\"\",\"total_points\":309,\"hp\":39,\"attack\":52,\"defense\":43,\"speed_attack\":60,\"speed_defense\":50,\"speed\":65,\"generation\":2,\"legendary\":\"False\"}\n")

var pokemonExternalJsonResponse = string("{\"id\":888,\"name\":\"Zacian-Hero\",\"type_1\":\"Fairy\",\"type_2\":\"\",\"total_points\":670,\"hp\":92,\"attack\":130,\"defense\":115,\"speed_attack\":80,\"speed_defense\":115,\"speed\":138,\"generation\":8,\"legendary\":\"False\"}\n")

var pokemonNotFoundResponse = "{\"message\":\"pokemon not found\"}\n"

var badRequestResponse = "{\"message\":\"invalid parameters\"}\n"

var pokemonsJsonResponse = "[{\"id\":155,\"name\":\"Cyndaquil\",\"type_1\":\"Fire\",\"type_2\":\"\",\"total_points\":309,\"hp\":39,\"attack\":52,\"defense\":43,\"speed_attack\":60,\"speed_defense\":50,\"speed\":65,\"generation\":2,\"legendary\":\"False\"},{\"id\":888,\"name\":\"Zacian-Hero\",\"type_1\":\"Fairy\",\"type_2\":\"\",\"total_points\":670,\"hp\":92,\"attack\":130,\"defense\":115,\"speed_attack\":80,\"speed_defense\":115,\"speed\":138,\"generation\":8,\"legendary\":\"False\"}]\n"

var badNumbersOfWorkersResponse = "{\"message\":\"cannot have more than 8 workers\"}\n"

type MockService struct {
	mock.Mock
}

type MockClient struct {
	mock.Mock
}

func (mock *MockService) FindById(pokemonId string) (*model.Pokemon, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*model.Pokemon), args.Error(1)
}

func (mock *MockService) SaveFromExternal(externalPokemon model.PokemonExternal) (*model.Pokemon, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*model.Pokemon), args.Error(1)
}

func (mock *MockClient) GetExternalPokemon(uri string, id string) model.PokemonExternal {
	args := mock.Called()
	result := args.Get(0)
	return result.(model.PokemonExternal)
}

func (mock *MockService) GetPokemons(dataType string, items int, items_per_workers int, workers int) ([]model.Pokemon, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]model.Pokemon), args.Error(1)
}

func TestGetPokemonById(t *testing.T) {
	testCases := []struct {
		name      string
		uri       string
		parameter string
		returnErr error
		status    int
		response  string
		returned  *model.Pokemon
	}{
		{
			name:      "Valid response",
			uri:       "/pokemons",
			parameter: "155",
			returnErr: nil,
			status:    200,
			response:  pokemonJsonResponse,
			returned:  &pokemon,
		},
		{
			name:      "Pokemon not found",
			uri:       "/pokemons",
			parameter: "999",
			returnErr: errors.New(common.PokemonNotFound),
			status:    404,
			response:  pokemonNotFoundResponse,
			returned:  &pokemon,
		},
		{
			name:      "Invalid request",
			uri:       "/pokemons",
			parameter: "abc",
			returnErr: errors.New(common.InvalidParameters),
			status:    400,
			response:  badRequestResponse,
			returned:  &pokemon,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			mockService := new(MockService)
			mockClient := new(MockClient)

			//Mock methods used on controller
			mockService.On("FindById").Return(tC.returned, tC.returnErr)

			//Create new HTTP request
			req, _ := http.NewRequest("GET", tC.uri, nil)
			req = mux.SetURLVars(req, map[string]string{"pokemon_id": tC.parameter})

			//Record the HTTP response
			response := httptest.NewRecorder()

			//Assign HTTP request controller function
			pokemonController := NewPokemonController(mockService, mockClient)
			controller := http.HandlerFunc(pokemonController.GetPokemonById)
			controller.ServeHTTP(response, req)

			status := response.Code
			responseBody := response.Body.String()

			if status != 400 {
				mockService.AssertExpectations(t)
			}

			assert.Equal(t, tC.status, status)
			assert.Equal(t, tC.response, responseBody)
		})
	}

}

func TestGetExternalPokemonById(t *testing.T) {
	testCases := []struct {
		name           string
		uri            string
		parameter      string
		returnErr      error
		status         int
		clientResponse model.PokemonExternal
		response       string
		returned       *model.Pokemon
	}{
		{
			name:           "Valid response",
			uri:            "/pokemons",
			parameter:      "888",
			returnErr:      nil,
			status:         200,
			clientResponse: externalPokemon,
			response:       pokemonExternalJsonResponse,
			returned:       &pokemonFromExternal,
		},
		{
			name:           "Pokemon not found",
			uri:            "/pokemons",
			parameter:      "72819364",
			returnErr:      errors.New(common.PokemonNotFound),
			status:         404,
			clientResponse: externalInvalidPokemon,
			response:       pokemonNotFoundResponse,
			returned:       &pokemonFromExternal,
		},
		{
			name:           "Invalid request",
			uri:            "/pokemons",
			parameter:      "abc",
			returnErr:      errors.New(common.InvalidParameters),
			status:         400,
			clientResponse: externalPokemon,
			response:       badRequestResponse,
			returned:       &pokemonFromExternal,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			mockService := new(MockService)
			mockClient := new(MockClient)

			//Mock methods used on controller
			mockClient.On("GetExternalPokemon").Return(tC.clientResponse)
			mockService.On("SaveFromExternal").Return(tC.returned, nil)

			//Create new HTTP request
			req, _ := http.NewRequest("POST", tC.uri, nil)
			req = mux.SetURLVars(req, map[string]string{"pokemon_id": tC.parameter})

			//Record the HTTP response
			response := httptest.NewRecorder()

			//Assign HTTP request controller function
			pokemonController := NewPokemonController(mockService, mockClient)
			controller := http.HandlerFunc(pokemonController.GetExternalPokemonById)

			//Dispatch the HTTP request
			controller.ServeHTTP(response, req)

			status := response.Code
			responseBody := response.Body.String()

			assert.Equal(t, tC.status, status)
			assert.Equal(t, tC.response, responseBody)
		})
	}
}

func TestGetPokemonsByWorker(t *testing.T) {
	var pokemonsByWorker []model.Pokemon

	pokemonsByWorker = append(pokemonsByWorker, pokemon)
	pokemonsByWorker = append(pokemonsByWorker, pokemonFromExternal)

	testCases := []struct {
		name      string
		uri       string
		query     string
		returnErr error
		status    int
		response  string
		returned  []model.Pokemon
	}{
		{
			name:      "Valid response",
			uri:       "/pokemons",
			query:     "?type=all&items_per_workers=1&items=2",
			returnErr: nil,
			status:    200,
			response:  pokemonsJsonResponse,
			returned:  pokemonsByWorker,
		},
		{
			name:      "Invalid request",
			uri:       "/pokemons",
			query:     "?type=1",
			returnErr: errors.New(common.InvalidParameters),
			status:    400,
			response:  badRequestResponse,
			returned:  pokemonsByWorker,
		},
		{
			name:      "Invalid workers",
			uri:       "/pokemons",
			query:     "?workers=10",
			returnErr: errors.New(common.InvalidWorkers),
			status:    400,
			response:  badNumbersOfWorkersResponse,
			returned:  pokemonsByWorker,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			mockService := new(MockService)
			mockClient := new(MockClient)

			//Mock methods used on controller
			mockService.On("GetPokemons").Return(tC.returned, tC.returnErr)

			//Create new HTTP request
			req, _ := http.NewRequest("GET", tC.uri+tC.query, nil)
			// req = mux.SetURLVars(req, map[string]string{"pokemon_id": tC.parameter})

			//Record the HTTP response
			response := httptest.NewRecorder()

			//Assign HTTP request controller function
			pokemonController := NewPokemonController(mockService, mockClient)
			controller := http.HandlerFunc(pokemonController.GetPokemonsByWorker)
			controller.ServeHTTP(response, req)

			status := response.Code
			responseBody := response.Body.String()

			if status != 400 {
				mockService.AssertExpectations(t)
			}

			assert.Equal(t, tC.status, status)
			assert.Equal(t, tC.response, responseBody)
		})
	}
}
