package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bifr0ns/academy-go-q32021/client"
	"github.com/bifr0ns/academy-go-q32021/common"
	"github.com/bifr0ns/academy-go-q32021/controller"
	"github.com/bifr0ns/academy-go-q32021/repository"
	"github.com/bifr0ns/academy-go-q32021/router"
	"github.com/bifr0ns/academy-go-q32021/service"
)

var (
	restClient        = client.NewRestyClient()
	pokemonRepository = repository.NewPokemonRepo()
	pokemonService    = service.NewPokemonService(&pokemonRepository)
	pokemonController = controller.NewPokemonController(&pokemonService, restClient)
	httpRouter        = router.NewMuxRouter()
)

func Start() {

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})

	httpRouter.GET("/pokemons/{pokemon_id}", pokemonController.GetPokemonById)
	httpRouter.POST("/pokemons/{pokemon_id}", pokemonController.GetExternalPokemonById)
	httpRouter.GET("/pokemons", pokemonController.GetPokemonsByWorker)

	httpRouter.SERVE(os.Getenv(common.Port))
}
