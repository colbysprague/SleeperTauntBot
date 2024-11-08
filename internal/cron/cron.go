package cron

import (
	"fmt"
	"log"
	"time"

	"github.com/colbysprague/SleeperBoard/internal/dbops"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/cron"
)

func PollPlayersEveryMin(app core.App) error {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {

		scheduler := cron.New()

		scheduler.MustAdd("playerPollService", "*/1 * * * *", func() {
			fmt.Println("Fetching sleeper matchups at: ", time.Now())
			err := dbops.BulkUpdatePlayerScoresInDB(app)

			if err != nil {
				log.Fatal(err)
			}
		})

		scheduler.Start()

		return nil
	})
	return nil
}
