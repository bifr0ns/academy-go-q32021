# Golang Bootcamp course application

Instructions regarding the challenge on GOLANGBOOTCAMP.md file. This file is regarding the application itself (how to run it, tests, etc.)

## Technologies

This project runs with Go. In order to run the project we need to install [Go](https://golang.org/doc/install)

## Run the project

We need to go into the root of the project and run

    go run main.go

## Tests

We can run tests with

    go test -v ./...

or

    go test ./...

## Functionality

The project contains two endpoints as of now (GET) _/pokemons/{pokemon_id}_ and (POST) _/api/pokemons/{pokemon_id}_

The first endpoint will let us get a pokemon from our CSV file. And the second endpoint will get a pokemon from an external API ([PokeAPI](https://pokeapi.co/)) and if the pokemon does not exists in our CSV it will save it.

We can hit the first endpoint with the following link http://localhost:8000/pokemons/155 and the second with http://localhost:8000/api/pokemons/888

### Pokemons endpoint

Example when pinging the previous link

    {
        "id": 155,
        "name": "Cyndaquil",
        "type_1": "Fire",
        "type_2": "",
        "total_points": 309,
        "hp": 39,
        "attack": 52,
        "defense": 43,
        "speed_attack": 60,
        "speed_defense": 50,
        "speed": 65,
        "generation": 2,
        "legendary": "False"
    }

### Api external pokemon

Example when pinging the second link

    {
        "id": 888,
        "name": "Zacian-Hero",
        "type_1": "Fairy",
        "type_2": "",
        "total_points": 670,
        "hp": 92,
        "attack": 130,
        "defense": 115,
        "speed_attack": 80,
        "speed_defense": 115,
        "speed": 138,
        "generation": 8,
        "legendary": "False"
    }
