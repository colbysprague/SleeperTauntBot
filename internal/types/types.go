package types

type NFLState struct {
	Week               int    `json:"week"`
	SeasonType         string `json:"season_type"`
	SeasonStartDate    string `json:"season_start_date"`
	Season             string `json:"season"`
	PreviousSeason     string `json:"previous_season"`
	Leg                int    `json:"leg"`
	LeagueSeason       string `json:"league_season"`
	LeagueCreateSeason string `json:"league_create_season"`
	DisplayWeek        int    `json:"display_week"`
}

type SleeperLeagueMatchup struct {
	Points         float64            `json:"points"`
	Players        []string           `json:"players"`
	RosterID       int                `json:"roster_id"`
	CustomPoints   *float64           `json:"custom_points"` // Null values are handled as pointers
	MatchupID      int                `json:"matchup_id"`
	Starters       []string           `json:"starters"`
	StartersPoints []float64          `json:"starters_points"`
	PlayersPoints  map[string]float64 `json:"players_points"` // Pointer to handle null values
}

// TODO:
// 		[ ] Make matchup object to return as []matchup and populate DB with new data
