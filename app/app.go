package app

import (
	"fmt"
	"net/http"

	"github.com/bifr0ns/academy-go-q32021/common"
	"github.com/bifr0ns/academy-go-q32021/controller"
	"github.com/bifr0ns/academy-go-q32021/repository"
	"github.com/bifr0ns/academy-go-q32021/router"
	"github.com/bifr0ns/academy-go-q32021/service"
)

var (
	pokemonRepository repository.PokemonRepository = repository.NewPokemonRepository()
	pokemonService    service.PokemonService       = service.NewPokemonService(pokemonRepository)
	pokemonController controller.PokemonController = controller.NewPokemonController(pokemonService)
	httpRouter        router.Router                = router.NewMuxRouter()
)

func Start() {

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})

	httpRouter.GET("/pokemons/{pokemon_id:[0-9]+}", pokemonController.GetPokemonById)

	httpRouter.SERVE(common.LocalHost + ":" + common.LocalPort)
}
