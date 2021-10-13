package client

import (
	"fmt"

	"github.com/bifr0ns/academy-go-q32021/model"

	"github.com/go-resty/resty/v2"
)

type restyClient struct {
	client *resty.Client
}

//NewRestyClient returns an struct of restyClient. Which contains the method:
//
//GetExternalPokemon(uri string, id string) model.PokemonExternal.
func NewRestyClient() PokemonClient {
	return &restyClient{resty.New()}
}

//GetExternalPokemon recieves a uri and an id, will return a model of PokemonExternal
//after processing some info with Resty.
func (rc *restyClient) GetExternalPokemon(uri string, id string) model.PokemonExternal {
	resp, _ := rc.client.R().
		SetPathParams(map[string]string{
			"pokemonId": fmt.Sprint(id),
		}).
		SetResult(model.PokemonExternal{}).
		SetHeader("Accept", "application/json").
		Get(uri)

	pokemon := *resp.Result().(*model.PokemonExternal)

	return pokemon
}
