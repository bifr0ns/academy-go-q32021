package client

import "github.com/bifr0ns/academy-go-q32021/model"

//GetExternalPokemon recieves a uri and an id, will return a model of PokemonExternal
//after processing some info from a rest client.
type PokemonClient interface {
	GetExternalPokemon(uri string, id string) model.PokemonExternal
}
