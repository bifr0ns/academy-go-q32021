package service

import (
	"github.com/bifr0ns/academy-go-q32021/model"
)

type pokemonRepository interface {
	GetPokemon(pokemonId string) (*model.Pokemon, error)
	SaveExternalPokemon(pokemon model.PokemonExternal) (*model.Pokemon, error)
}

//PokemonService cointains an interface of pokemon repository which contains two methods.
//
//GetPokemon recieves a string, and returns a model of Pokemon and error if any.
//
//SaveExternalPokemon recieves a model of PokemonExternal, and returns a model of Pokemon and error if any.
type PokemonService struct {
	repo pokemonRepository
}

//NewPokemonController expects to recieve PokemonRepository, returns an struct of PokemonService.
func NewPokemonService(repository pokemonRepository) PokemonService {
	return PokemonService{repository}
}

//FindById of type service recieves an id of type string to call a repository.
//Will return the response of the repository, a model of Pokemon and error if any.
func (ps *PokemonService) FindById(pokemonId string) (*model.Pokemon, error) {
	return ps.repo.GetPokemon(pokemonId)
}

//SaveFromExternal of type service recieves a model of PokemonExternal to call a repository.
//Will return the repsonse of the repository, a model of Pokemon and error if any.
func (ps *PokemonService) SaveFromExternal(externalPokemon model.PokemonExternal) (*model.Pokemon, error) {
	return ps.repo.SaveExternalPokemon(externalPokemon)
}
