package model

//PokemonExternal is used when extracting data from external API.
type PokemonExternal struct {
	Id           int            `json:"id"`
	Name         string         `json:"name"`
	NotLegendary bool           `json:"is_default"`
	Types        []PokemonTypes `json:"types"`
	Stats        []PokemonStats `json:"stats"`
}

type PokemonTypes struct {
	Slot int               `json:"slot"`
	Type map[string]string `json:"type"`
}

type PokemonStats struct {
	BaseStat int               `json:"base_stat"`
	Stat     map[string]string `json:"stat"`
}
