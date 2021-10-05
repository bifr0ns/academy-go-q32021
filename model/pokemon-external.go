package model

//PokemonExternal is used when extracting data from external API.
type PokemonExternal struct {
	Id           int            `json:"id"`
	Name         string         `json:"name"`
	NotLegendary bool           `json:"is_default"`
	Types        []pokemonTypes `json:"types"`
	Stats        []pokemonStats `json:"stats"`
}

type pokemonTypes struct {
	Slot int               `json:"slot"`
	Type map[string]string `json:"type"`
}

type pokemonStats struct {
	BaseStat int               `json:"base_stat"`
	Stat     map[string]string `json:"stat"`
}
