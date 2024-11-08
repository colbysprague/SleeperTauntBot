package main

import (
	"log"
	"os"

	"github.com/colbysprague/SleeperBoard/internal/collectiondefs"
	"github.com/colbysprague/SleeperBoard/internal/cron"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func main() {

	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {

		err := collectiondefs.InitAllCollectionsForPocketBaseApp(app)

		if err != nil {
			return err
		}

		return nil
	})

	cron.PollPlayersEveryMin(app)

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

// 1. create main.go application
// 2. To init the dependencies, run go mod init myapp && go mod tidy.
// 3. To start the application, run go run main.go serve.
