package service

import (
	"strings"
	"testing"

	"github.com/bifr0ns/academy-go-q32021/model"

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

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) GetPokemon(pokemonId string, csvFileName string) (*model.Pokemon, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*model.Pokemon), args.Error(1)
}

func (mock *MockRepository) SaveExternalPokemon(pokemon model.PokemonExternal, csvFileName string) (*model.Pokemon, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*model.Pokemon), args.Error(1)
}

func (mock *MockRepository) GetPokemons(dataType string, items int, items_per_workers int, workers int, csvFileName string) ([]model.Pokemon, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]model.Pokemon), args.Error(1)
}

func TestFindById(t *testing.T) {
	testCases := []struct {
		name        string
		id          string
		pokemonName string
	}{
		{
			name:        "Find by id",
			id:          "025",
			pokemonName: "Cyndaquil",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			mockRepo := new(MockRepository)

			mockRepo.On("GetPokemon").Return(&pokemon, nil)

			testService := NewPokemonService(mockRepo)

			result, _ := testService.FindById(tC.id)

			mockRepo.AssertExpectations(t)

			assert.Equal(t, tC.pokemonName, result.Name)
		})
	}
}

func TestSaveFromExternal(t *testing.T) {
	mockRepo := new(MockRepository)

	mockRepo.On("SaveExternalPokemon").Return(&pokemonFromExternal, nil)

	testService := NewPokemonService(mockRepo)

	result, _ := testService.SaveFromExternal(externalPokemon)

	mockRepo.AssertExpectations(t)

	assert.Equal(t, strings.Title(externalPokemon.Name), result.Name)
}

func TestGetPokemons(t *testing.T) {
	mockRepo := new(MockRepository)
	var pokemons []model.Pokemon

	pokemons = append(pokemons, pokemon)
	pokemons = append(pokemons, pokemon)

	mockRepo.On("GetPokemons").Return(pokemons, nil)

	testService := NewPokemonService(mockRepo)

	result, _ := testService.GetPokemons("all", 2, 1, 2)

	mockRepo.AssertExpectations(t)

	assert.Equal(t, strings.Title(pokemons[0].Name), result[0].Name)
}
