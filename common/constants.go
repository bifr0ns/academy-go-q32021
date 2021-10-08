package common

const (
	LocalHost           = "localhost"
	LocalPort           = "8000"
	PokemonNotFound     = "pokemon not found"
	PokemonAlreadyExist = "pokemon is already on the CSV File"
	InvalidParameters   = "invalid parameters"

	CsvPokemonName              = "data/pokemon.csv"
	CsvPokemonName_TestOriginal = "../data/testdata/pokemon_original_test.csv"
	CsvPokemonName_Test         = "../data/testdata/pokemon_test.csv"

	GetPokemonUri = "https://pokeapi.co/api/v2/pokemon/{pokemonId}"
)
