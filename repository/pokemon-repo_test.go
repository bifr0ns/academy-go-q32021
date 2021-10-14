package repository

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/bifr0ns/academy-go-q32021/common"
	"github.com/bifr0ns/academy-go-q32021/model"

	"github.com/stretchr/testify/assert"
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
var pokemon1 = model.Pokemon{
	Id:           1,
	Name:         "Bulbasaur",
	Type1:        "Grass",
	Type2:        "Poison",
	Total:        318,
	HP:           45,
	Attack:       49,
	Defense:      49,
	SpeedAttack:  65,
	SpeedDefense: 65,
	Speed:        45,
	Generation:   1,
	Legendary:    "False",
}
var pokemon2 = model.Pokemon{
	Id:           2,
	Name:         "Ivysaur",
	Type1:        "Grass",
	Type2:        "Poison",
	Total:        405,
	HP:           60,
	Attack:       62,
	Defense:      63,
	SpeedAttack:  80,
	SpeedDefense: 80,
	Speed:        60,
	Generation:   1,
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

func TestGetPokemon(t *testing.T) {
	testCases := []struct {
		name      string
		pokemonId string
		fixture   string
		returnErr bool
		response  model.Pokemon
	}{
		{
			name:      "ValidPokemon",
			pokemonId: "25",
			fixture:   common.CsvPokemonName_Test,
			returnErr: false,
		},
		{
			name:      "PokemonNotFound",
			pokemonId: "999",
			fixture:   common.CsvPokemonName_Test,
			returnErr: true,
		},
		{
			name:      "ValidFormat",
			pokemonId: "155",
			fixture:   common.CsvPokemonName_Test,
			returnErr: false,
			response:  pokemon,
		},
		{
			name:      "InvalidFile",
			pokemonId: "25",
			fixture:   common.CsvPokemonName,
			returnErr: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			repo := NewPokemonRepo()

			resp, err := repo.GetPokemon(tC.pokemonId, tC.fixture)
			returnedErr := err != nil

			if tC.pokemonId == "155" {
				assert.ObjectsAreEqual(tC.response, resp)
			}

			if returnedErr != tC.returnErr {
				t.Fatalf("Expected returnErr: %v, got: %v with %v", tC.returnErr, returnedErr, err)
			}
		})
	}
}

func TestSaveExternalPokemon(t *testing.T) {
	testCases := []struct {
		name      string
		fixture   string
		returnErr bool
		response  model.Pokemon
		request   model.PokemonExternal
	}{
		{
			name:      "Valid save of external pokemon",
			fixture:   common.CsvPokemonName_Test,
			returnErr: false,
			response:  pokemonFromExternal,
			request:   externalPokemon,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			repo := NewPokemonRepo()

			resp, err := repo.SaveExternalPokemon(externalPokemon, tC.fixture)

			if err != nil {
				t.Fatalf(err.Error())
			}

			assert.ObjectsAreEqual(tC.response, resp)
			teardown()
		})
	}
}

func teardown() {
	e := os.Remove(common.CsvPokemonName_Test)
	if e != nil {
		fmt.Println(e)
	}

	input, err := ioutil.ReadFile(common.CsvPokemonName_TestOriginal)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(common.CsvPokemonName_Test, input, 0644)
	if err != nil {
		fmt.Println("Error creating", common.CsvPokemonName_Test)
		fmt.Println(err)
		return
	}
}

func TestGetPokemons(t *testing.T) {
	var pokemonsByWorker []model.Pokemon

	pokemonsByWorker = append(pokemonsByWorker, pokemon1)
	pokemonsByWorker = append(pokemonsByWorker, pokemon2)

	testCases := []struct {
		name              string
		fixture           string
		returnErr         bool
		response          []model.Pokemon
		dataType          string
		items             int
		items_per_workers int
	}{
		{
			name:              "Valid",
			fixture:           common.CsvPokemonName_Test,
			returnErr:         false,
			response:          pokemonsByWorker,
			dataType:          "all",
			items:             2,
			items_per_workers: 1,
		},
		{
			name:              "InvalidFile",
			fixture:           common.CsvPokemonName,
			returnErr:         true,
			dataType:          "all",
			items:             3,
			items_per_workers: 1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {

			repo := NewPokemonRepo()

			resp, err := repo.GetPokemons(tC.dataType, tC.items, tC.items_per_workers, tC.fixture)
			returnedErr := err != nil

			if tC.returnErr == false {
				assert.ObjectsAreEqual(tC.response, resp)
			}

			if returnedErr != tC.returnErr {
				t.Fatalf("Expected returnErr: %v, got: %v with %v", tC.returnErr, returnedErr, err)
			}
		})
	}
}
