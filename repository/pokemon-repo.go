package repository

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/bifr0ns/academy-go-q32021/common"
	"github.com/bifr0ns/academy-go-q32021/model"
)

type PokemonRepository interface {
	GetPokemon(pokemonId string) (*model.Pokemon, error)
}

type repo struct{}

func NewPokemonRepository() PokemonRepository {
	return &repo{}
}

func (*repo) GetPokemon(pokemonId string) (*model.Pokemon, error) {

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
