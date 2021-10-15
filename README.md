# Golang Bootcamp course application

Instructions regarding the challenge on GOLANGBOOTCAMP.md file. This file is regarding the application itself (how to run it, tests, etc.)

## Technologies

This project runs with Go. In order to run the project we need to install [Go](https://golang.org/doc/install)

## Run the project

We need to add an environment variable called PORT in our computer where we will run our local project based on os.Getenv("PORT")

    export PORT=8000

We need to go into the root of the project and run

    go run main.go

### Run the project using a Docker image

_Note: The image right now has the version of the previous PR #1 (Second Deliverable) and I will wait for PR #2 to merge to build a push the new one._

    docker run --rm -p 8080:8000 bifr0ns/academy-go-q32021

And you will use the following link for endpoints

    http://localhost:8080/pokemons ...

## Tests

We can run tests with

    go test -v ./...

or

    go test ./...

## Functionality

The project contains two endpoints as of now (GET) _/pokemons/{pokemon_id}_ , (POST) _/api/pokemons/{pokemon_id}_ and _/pokemons?type={type}&items_per_workers={items_per_workers}&items={items}_

The first endpoint will let us get a pokemon from our CSV file. And the second endpoint will get a pokemon from an external API ([PokeAPI](https://pokeapi.co/)) and if the pokemon does not exists in our CSV it will save it. The third endpoint will get us an array of pokemons, based on our query parameters (items, type, items_per_workers) all of them optional.

We can hit the first endpoint with the following link http://localhost:8000/pokemons/155, the second with http://localhost:8000/api/pokemons/888 and the third with http://localhost:8000/pokemons?type=even&items_per_workers=10&items=21

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

### Query pokemons

Example when pinging the third link. All of the parameters are optional.
Type can be _all/even/odd_

    [
        {
            "id": 1,
            "name": "Bulbasaur",
            "type_1": "Grass",
            "type_2": "Poison",
            "total_points": 318,
            "hp": 45,
            "attack": 49,
            "defense": 49,
            "speed_attack": 65,
            "speed_defense": 65,
            "speed": 45,
            "generation": 1,
            "legendary": "False"
        },
        {
            "id": 2,
            "name": "Ivysaur",
            "type_1": "Grass",
            "type_2": "Poison",
            "total_points": 405,
            "hp": 60,
            "attack": 62,
            "defense": 63,
            "speed_attack": 80,
            "speed_defense": 80,
            "speed": 60,
            "generation": 1,
            "legendary": "False"
        },
        {
            "id": 3,
            "name": "Venusaur",
            "type_1": "Grass",
            "type_2": "Poison",
            "total_points": 525,
            "hp": 80,
            "attack": 82,
            "defense": 83,
            "speed_attack": 100,
            "speed_defense": 100,
            "speed": 80,
            "generation": 1,
            "legendary": "False"
        }
    ]

## Build and run a docker image

Build the docker image

    docker build -t academy-go-q32021 .

Run the docker image locally

    docker run -p 8080:8000 academy-go-q32021
