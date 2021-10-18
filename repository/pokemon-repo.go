package repository

import (
	"encoding/csv"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/bifr0ns/academy-go-q32021/common"
	"github.com/bifr0ns/academy-go-q32021/model"
)

//PokemonRepo returns the struct to be used for this repository.
type PokemonRepo struct{}

func NewPokemonRepo() PokemonRepo {
	return PokemonRepo{}
}

//GetPokemon of type PokemonRepo recieves an id of type string.
//Opens the CSV file and finds the pokemon by the given id, if found creates a Pokemon model.
//Will return a model of Pokemon and error if any.
func (pr *PokemonRepo) GetPokemon(pokemonId string, csvFileName string) (*model.Pokemon, error) {

	csvFile, err := os.Open(csvFileName)
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return nil, err
	}

	for _, line := range csvLines {
		id, _ := strconv.Atoi(line[0])
		total, _ := strconv.Atoi(line[4])
		hp, _ := strconv.Atoi(line[5])
		attack, _ := strconv.Atoi(line[6])
		defense, _ := strconv.Atoi(line[7])
		speedAttack, _ := strconv.Atoi(line[8])
		speedDefense, _ := strconv.Atoi(line[9])
		speed, _ := strconv.Atoi(line[10])
		generation, _ := strconv.Atoi(line[11])

		pokemon := model.Pokemon{
			Id:           id,
			Name:         line[1],
			Type1:        line[2],
			Type2:        line[3],
			Total:        total,
			HP:           hp,
			Attack:       attack,
			Defense:      defense,
			SpeedAttack:  speedAttack,
			SpeedDefense: speedDefense,
			Speed:        speed,
			Generation:   generation,
			Legendary:    line[12],
		}

		id, _ = strconv.Atoi(pokemonId)
		if pokemon.Id == id {

			return &pokemon, nil
		}
	}

	return nil, errors.New(common.PokemonNotFound)
}

//SaveExternalPokemon of type PokemonRepo recieves a model of PokemonExternal.
//Searchs if the pokemon doesnt exist already in the CSV file.
//Writes the new pokemon if its not in the CSV, and creates the Pokemon model.
//Will return a model of Pokemon and error if any.
func (pr *PokemonRepo) SaveExternalPokemon(externalPokemon model.PokemonExternal, csvFileName string) (*model.Pokemon, error) {

	//Checks if pokemon already exists in the CSV file
	csvPokemon, _ := pr.GetPokemon(strconv.Itoa(externalPokemon.Id), csvFileName)
	if csvPokemon != nil {
		return nil, errors.New(common.PokemonAlreadyExist)
	}

	csvFile, err := os.OpenFile(csvFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	fmt.Println("Preparing to write into CSV file")

	writer := csv.NewWriter(csvFile)

	pokemon := getPokemonFromExternal(externalPokemon)

	if err := writer.Write(pokemon); err != nil {
		return nil, err
	}

	writer.Flush()

	pokemonCreated := createPokemon(pokemon)

	return &pokemonCreated, nil
}

func getGeneration(id int) int {
	switch {
	case id <= 151:
		return 1
	case id <= 251:
		return 2
	case id <= 386:
		return 3
	case id <= 493:
		return 4
	case id <= 649:
		return 5
	case id <= 721:
		return 6
	case id <= 809:
		return 7
	case id <= 901:
		return 8
	}
	return 1
}

func getPokemonFromExternal(externalPokemon model.PokemonExternal) []string {
	id := externalPokemon.Id
	pokemonId := strconv.Itoa(int(id))
	name := strings.Title(externalPokemon.Name)
	type1 := strings.Title(externalPokemon.Types[0].Type["name"])
	type2 := ""
	if len(externalPokemon.Types) > 1 {
		type2 = strings.Title(externalPokemon.Types[1].Type["name"])
	}
	total := externalPokemon.Stats[0].BaseStat + externalPokemon.Stats[1].BaseStat + externalPokemon.Stats[2].BaseStat +
		externalPokemon.Stats[3].BaseStat + externalPokemon.Stats[4].BaseStat + externalPokemon.Stats[5].BaseStat
	hp := externalPokemon.Stats[0].BaseStat
	attack := externalPokemon.Stats[1].BaseStat
	defense := externalPokemon.Stats[2].BaseStat
	speedAttack := externalPokemon.Stats[3].BaseStat
	speedDefense := externalPokemon.Stats[4].BaseStat
	speed := externalPokemon.Stats[5].BaseStat
	generation := getGeneration(id)
	legendary := strings.Title(strconv.FormatBool(!externalPokemon.NotLegendary))

	pokemon := []string{
		pokemonId,
		name,
		type1,
		type2,
		strconv.Itoa(total),
		strconv.Itoa(hp),
		strconv.Itoa(attack),
		strconv.Itoa(defense),
		strconv.Itoa(speedAttack),
		strconv.Itoa(speedDefense),
		strconv.Itoa(speed),
		strconv.Itoa(generation),
		legendary,
	}

	return pokemon
}

func createPokemon(pokemon []string) model.Pokemon {

	id, _ := strconv.Atoi(pokemon[0])
	total, _ := strconv.Atoi(pokemon[4])
	hp, _ := strconv.Atoi(pokemon[5])
	attack, _ := strconv.Atoi(pokemon[6])
	defense, _ := strconv.Atoi(pokemon[7])
	speedAttack, _ := strconv.Atoi(pokemon[8])
	speedDefense, _ := strconv.Atoi(pokemon[9])
	speed, _ := strconv.Atoi(pokemon[10])
	generation, _ := strconv.Atoi(pokemon[11])

	pokemonCreated := model.Pokemon{
		Id:           id,
		Name:         pokemon[1],
		Type1:        pokemon[2],
		Type2:        pokemon[3],
		Total:        total,
		HP:           hp,
		Attack:       attack,
		Defense:      defense,
		SpeedAttack:  speedAttack,
		SpeedDefense: speedDefense,
		Speed:        speed,
		Generation:   generation,
		Legendary:    pokemon[12],
	}

	return pokemonCreated
}

var wg sync.WaitGroup

//GetPokemonsof type PokemonRepo recieves data based on a query
//Opens the csv file
//Based on workers and go routines will find pokemons filtered by the query, will return them
func (pr *PokemonRepo) GetPokemons(dataType string, items int, items_per_workers int, workers int, csvFileName string) ([]model.Pokemon, error) {
	var pokemons []model.Pokemon

	csvLines, err := openFile(csvFileName)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	dataSize := len(csvLines) - 1
	isAll := 1

	if dataType != "all" {
		isAll = 2
	}
	if items == -1 || items > dataSize {
		items = dataSize
	}
	if workers == -1 {
		workers = getWorkers(items)
	}
	fmt.Println("Working with ", workers, " workers")
	if items_per_workers == -1 {
		items_per_workers = int(math.Ceil(float64(items) / float64(workers)))
	}

	begin := 1
	tail := items_per_workers + 1
	if isAll == 2 && tail*isAll+1 < dataSize {
		tail = tail*isAll - 1
	}
	helper := tail - 1

	pokemonChannel := make(chan model.Pokemon, 300)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		if i == workers-1 && isAll == 1 {
			tail = items + 1
		}
		if tail > dataSize+1 {
			tail = items + 1
		}
		go getPokemons(dataType, items, items_per_workers, csvLines[begin:tail], pokemonChannel)
		begin = tail
		tail = tail + helper
	}

	wg.Add(1)

	go func(pokemonChannel <-chan model.Pokemon) {
		for poke := range pokemonChannel {
			pokemons = append(pokemons, poke)
			if len(pokemons) == items {
				break
			} else if len(pokemons) == items_per_workers*workers {
				break
			} else if len(pokemons) >= items/isAll && items == dataSize {
				break
			}
		}
		fmt.Println("Closing reader routine")
		wg.Done()
	}(pokemonChannel)

	wg.Wait()
	close(pokemonChannel)

	return pokemons, nil
}

func openFile(csvFileName string) ([][]string, error) {
	csvFile, err := os.Open(csvFileName)
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return nil, err
	}

	return csvLines, nil
}

func getWorkers(items int) int {
	workers := math.Cbrt(float64(items))

	if workers > 8 {
		workers = 8
	}

	return int(workers)
}

func getPokemons(dataType string, items int, items_per_worker int, csvLines [][]string, pokemonChannel chan<- model.Pokemon) {

	pokemonsAdded := 0

	for _, line := range csvLines {
		id, _ := strconv.Atoi(line[0])
		total, _ := strconv.Atoi(line[4])
		hp, _ := strconv.Atoi(line[5])
		attack, _ := strconv.Atoi(line[6])
		defense, _ := strconv.Atoi(line[7])
		speedAttack, _ := strconv.Atoi(line[8])
		speedDefense, _ := strconv.Atoi(line[9])
		speed, _ := strconv.Atoi(line[10])
		generation, _ := strconv.Atoi(line[11])

		pokemon := model.Pokemon{
			Id:           id,
			Name:         line[1],
			Type1:        line[2],
			Type2:        line[3],
			Total:        total,
			HP:           hp,
			Attack:       attack,
			Defense:      defense,
			SpeedAttack:  speedAttack,
			SpeedDefense: speedDefense,
			Speed:        speed,
			Generation:   generation,
			Legendary:    line[12],
		}

		switch {
		case dataType == "odd":
			if id%2 != 0 {
				pokemonChannel <- pokemon
				pokemonsAdded++
			}
		case dataType == "even":
			if id%2 == 0 {
				pokemonChannel <- pokemon
				pokemonsAdded++
			}
		default:
			pokemonChannel <- pokemon
			pokemonsAdded++
		}

		if pokemonsAdded == items {
			fmt.Println("=======CLOSING WORKER ITEMS")
			wg.Done()
			return
		} else if pokemonsAdded == items_per_worker {
			fmt.Println("=======CLOSING WORKER PER WORKER")
			wg.Done()
			return
		}
	}
	fmt.Println("=======CLOSING WORKER EOF")
	wg.Done()
}
