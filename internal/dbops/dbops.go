package dbops

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/colbysprague/SleeperBoard/internal/sleeper"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

var playerPointsCache = make(map[string]float64)

func checkIfPlayerScoreExist(app core.App, player_id string, nflWeek int) (bool, error) {
	_, err := app.Dao().FindFirstRecordByFilter("players_scores", fmt.Sprintf("player_id = %s && nfl_week = %d", player_id, nflWeek))

	if err == sql.ErrNoRows {
		// no record
		return false, nil
	} else if err != nil {
		log.Fatal(err)
	} else {
		return true, nil
	}

	return false, nil
}

func updateExistingPlayerScoreRecord(app core.App, player_id string, points float64, nflWeek int, collection *models.Collection) error {
	// does player exist already for given week?
	record, err := app.Dao().FindFirstRecordByFilter("players_scores", fmt.Sprintf("player_id = %s && nfl_week = %d", player_id, nflWeek))

	if err != nil {
		log.Fatal(err)
	}

	record.Set("player_id", player_id)
	record.Set("points", points)
	record.Set("nfl_week", nflWeek)

	if err := app.Dao().SaveRecord(record); err != nil {
		return err
	}
	return nil
}

func createNewPlayerScoreRecord(app core.App, player_id string, points float64, nflWeek int, collection *models.Collection) error {

	record := models.NewRecord(collection)
	record.Set("player_id", player_id)
	record.Set("points", points)
	record.Set("nfl_week", nflWeek)

	if err := app.Dao().SaveRecord(record); err != nil {
		return err
	}
	return nil
}

func BulkUpdatePlayerScoresInDB(app core.App) error {
	// retrieve matchups,
	leagueMatchups, _ := sleeper.GetSleeperLeagueMatchups()
	nflWeek, _ := sleeper.GetNFLWeek()

	// get collection
	collection, err := app.Dao().FindCollectionByNameOrId("players_scores")

	if err != nil {
		return err
	}

	// loop over and update matchups
	for _, matchup := range leagueMatchups {
		for player_id, points := range matchup.PlayersPoints {
			fmt.Println("Updating score for player: ", player_id, points)

			// does player exist?
			playerExists, _ := checkIfPlayerScoreExist(app, player_id, nflWeek)

			if playerExists {
				// is player in cache
				if cachedPoints, exists := playerPointsCache[player_id]; exists {
					if cachedPoints != points {
						// update record and cache
						updateExistingPlayerScoreRecord(app, player_id, points, nflWeek, collection)
						playerPointsCache[player_id] = points
						continue
					}
				} else {
					fmt.Println("For some reason player not in cache", player_id)
				}
			} else {
				createNewPlayerScoreRecord(app, player_id, points, nflWeek, collection)
				playerPointsCache[player_id] = points
				continue
			}
		}
	}

	return nil
}
