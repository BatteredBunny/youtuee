package main

import (
	"fmt"
	"log"

	"github.com/BatteredBunny/youtuee/cmd"
)

func main() {
	app := cmd.NewApplication()
	app.ParseConfig()
	app.SetupYtApi()

	app.SetupRatelimiter()

	app.SetupRouter()

	log.Printf("Starting api on port %d\n", app.Conf.Port)
	log.Fatal(app.Router.Run(fmt.Sprintf(":%d", app.Conf.Port)))
}
