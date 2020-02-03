package tft

//Tier is the tier within the league
type Tier string

//Division is the division within the league
type Division string

//LeagueID is the id of the league
type LeagueID string

//LeagueList represents a LeagueListDTO
type LeagueList struct {
	LeagueID LeagueID      `json:"leagueId"`
	Tier     Tier          `json:"tier"`
	Entries  []*LeagueItem `json:"entries"`
	Queue    string        `json:"queue"`
	Name     string        `json:"name"`
}

//LeagueItem represents a LeagueItemDTO
type LeagueItem struct {
	SummonerName string      `json:"summonerName"`
	HotStreak    bool        `json:"hotStreak"`
	MiniSeries   *MiniSeries `json:"miniSeries"`
	Wins         int32       `json:"wins"`
	Veteran      bool        `json:"veteran"`
	Losses       int32       `json:"losses"`
	FreshBlood   bool        `json:"freshBlood"`
	Inactive     bool        `json:"inactive"`
	Rank         string      `json:"rank"`
	SummonerID   string      `json:"summonerId"`
	LeaguePoints int32       `json:"leaguePoints"`
}

//MiniSeries represents a MiniSeriesDTO
type MiniSeries struct {
	Progress string `json:"progress"`
	Losses   int32  `json:"losses"`
	Target   int32  `json:"target"`
	Wins     int32  `json:"wins"`
}

//LeagueEntry represents a LeagueEntryDTO
type LeagueEntry struct {
	QueueType    string      `json:"queueType"`
	SummonerName string      `json:"summonerName"`
	HotStreak    bool        `json:"hotStreak"`
	MiniSeries   *MiniSeries `json:"miniSeries"`
	Wins         int32       `json:"wins"`
	Veteran      bool        `json:"veteran"`
	Losses       int32       `json:"losses"`
	Rank         string      `json:"rank"`
	LeagueID     LeagueID    `json:"leagueId"`
	Inactive     bool        `json:"inactive"`
	FreshBlood   bool        `json:"freshBlood"`
	Tier         Tier        `json:"tier"`
	SummonerID   string      `json:"summonerId"`
	LeaguePoints int32       `json:"leaguePoints"`
}
