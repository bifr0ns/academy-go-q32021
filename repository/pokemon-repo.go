package repository

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bifr0ns/academy-go-q32021/common"
	"github.com/bifr0ns/academy-go-q32021/model"
)

//PokemonRepo returns the struct to be used for this repository.
type PokemonRepo struct{}

//GetPokemon of type PokemonRepo recieves an id of type string.
//Opens the CSV file and finds the pokemon by the given id, if found creates a Pokemon model.
//Will return a model of Pokemon and error if any.
func (pr *PokemonRepo) GetPokemon(pokemonId string) (*model.Pokemon, error) {

	csvFile, err := os.Open(common.CsvPokemonName)
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
func (pr *PokemonRepo) SaveExternalPokemon(externalPokemon model.PokemonExternal) (*model.Pokemon, error) {

	//Checks if pokemon already exists in the CSV file
	csvPokemon, _ := pr.GetPokemon(strconv.Itoa(externalPokemon.Id))
	if csvPokemon != nil {
		return nil, errors.New(common.PokemonAlreadyExist)
	}

	csvFile, err := os.OpenFile(common.CsvPokemonName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	fmt.Println("Preparing to write into CSV file")

	writer := csv.NewWriter(csvFile)

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

	if err := writer.Write(pokemon); err != nil {
		return nil, err
	}

	writer.Flush()

	pokemonCreated := model.Pokemon{
		Id:           id,
		Name:         name,
		Type1:        type1,
		Type2:        type2,
		Total:        total,
		HP:           hp,
		Attack:       attack,
		Defense:      defense,
		SpeedAttack:  speedAttack,
		SpeedDefense: speedDefense,
		Speed:        speed,
		Generation:   generation,
		Legendary:    legendary,
	}

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
