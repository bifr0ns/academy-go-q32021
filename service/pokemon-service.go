package service

import (
	"github.com/bifr0ns/academy-go-q32021/model"
	"github.com/bifr0ns/academy-go-q32021/repository"
)

//FindById recieves a string, and returns a model of Pokemon and error if any.
//
//SaveFromExternal recieves a model of PokemonExternal, and returns a model of Pokemon and error if any.
type PokemonService interface {
	FindById(pokemonId string) (*model.Pokemon, error)
	SaveFromExternal(externalPokemon model.PokemonExternal) (*model.Pokemon, error)
}

type service struct{}

//NewPokemonController expects to recieve PokemonRepository, returns an interface of PokemonService.
func NewPokemonService(repository repository.PokemonRepository) PokemonService {
	repo = repository
	return &service{}
}

var (
	repo repository.PokemonRepository
)

//FindById of type service recieves an id of type string to call a repository.
//Will return the response of the repository, a model of Pokemon and error if any.
func (*service) FindById(pokemonId string) (*model.Pokemon, error) {
	return repo.GetPokemon(pokemonId)
}

//SaveFromExternal of type service recieves a model of PokemonExternal to call a repository.
//Will return the repsonse of the repository, a model of Pokemon and error if any.
func (*service) SaveFromExternal(externalPokemon model.PokemonExternal) (*model.Pokemon, error) {
	return repo.SaveExternalPokemon(externalPokemon)
}
