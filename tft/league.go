package tft

//Tier is the tier within the league
type Tier string

//Division is the division within the league
type Division string

//LeagueID is the id of the league
type LeagueID string

//LeagueList represents a LeagueListDTO
type LeagueList struct {
	LeagueID LeagueID
	Tier     Tier
	Entries  []*LeagueItem
}

//LeagueItem represents a LeagueItemDTO
type LeagueItem struct {
	SummonerName string
	HotStreak    bool
	MiniSeries   *MiniSeries
	Wins         int32
	Veteran      bool
	Losses       int32
	FreshBlood   bool
	Inactive     bool
	Rank         string
	SummonerID   string
	LeaguePoints int32
}

//MiniSeries represents a MiniSeriesDTO
type MiniSeries struct {
	Progress string
	Losses   int32
	Target   int32
	Wins     int32
}

//LeagueEntry represents a LeagueEntryDTO
type LeagueEntry struct {
	QueueType string
	SummonerName string
	HotStreak bool
	MiniSeries *MiniSeries
	Wins int32
	Veteran bool
	Losses int32
	Rank string
	LeagueID LeagueID
	Inactive bool
	FreshBlood bool
	Tier Tier
	SummonerID string
	LeaguePoints int32
}
