package common

const (
	LocalHost           = "localhost"
	Port                = "PORT"
	PokemonNotFound     = "pokemon not found"
	PokemonAlreadyExist = "pokemon is already on the CSV File"
	InvalidParameters   = "invalid parameters"
	InvalidWorkers      = "cannot have more than 8 workers"

	CsvPokemonName              = "data/pokemon.csv"
	CsvPokemonName_TestOriginal = "./testbackup/pokemon_original_test.csv"
	CsvPokemonName_Test         = "../data/testdata/pokemon_test.csv"

	GetPokemonUri = "https://pokeapi.co/api/v2/pokemon/{pokemonId}"
)
