package sleeper

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/colbysprague/SleeperBoard/internal/types"
)

type NFLState types.NFLState

func fetchDataFromURL(url string) ([]byte, error) {
	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %v", err)
	}
	defer resp.Body.Close()

	// Check for HTTP errors
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	return body, nil
}

func GetNFLState() (*NFLState, error) {
	body, err := fetchDataFromURL("https://api.sleeper.app/v1/state/nfl")
	if err != nil {
		return nil, err
	}

	var nflData NFLState
	err = json.Unmarshal(body, &nflData)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	return &nflData, nil
}

func GetNFLWeek() (int, error) {
	nflState, err := GetNFLState()

	if err != nil {
		return 0, err
	}

	return nflState.Week, nil
}

func GetSleeperLeagueMatchups() ([]types.SleeperLeagueMatchup, error) {
	nflWeek, _ := GetNFLWeek()

	matchupData, _ := fetchDataFromURL(fmt.Sprintf("https://api.sleeper.app/v1/league/1136032070784147456/matchups/%d", nflWeek))

	var leagueMatchups []types.SleeperLeagueMatchup

	err := json.Unmarshal(matchupData, &leagueMatchups)

	if err != nil {
		log.Fatal(err)
	}

	return leagueMatchups, nil
}
