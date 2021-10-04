package service

import (
	"github.com/bifr0ns/academy-go-q32021/model"
	"github.com/bifr0ns/academy-go-q32021/repository"
)

type PokemonService interface {
	FindById(pokemonId string) (*model.Pokemon, error)
}

type service struct{}

func NewPokemonService(repository repository.PokemonRepository) PokemonService {
	repo = repository
	return &service{}
}

var (
	repo repository.PokemonRepository
)

func (*service) FindById(pokemonId string) (*model.Pokemon, error) {
	return repo.GetPokemon(pokemonId)
}
