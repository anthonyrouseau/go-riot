package lol

import (
	"github.com/anthonyrouseau/go-riot/queue"
)

//LeagueID is the Id for the league for a given queue
type LeagueID string

//Tier is the tier within a league
type Tier string

//Division is the division within a league
type Division string

//LeagueInfo represents a LeagueListDTO
type LeagueInfo struct {
	id      LeagueID
	tier    Tier
	queue   queue.Name
	name    string
	entries []*LeagueItem
}

//LeagueItem represents a LeagueItemDTO
type LeagueItem struct {
	summonerName SummonerName
	hotStreak    bool
	miniSeries   *MiniSeries
	wins         int32
	veteran      bool
	losses       int32
	freshBlood   bool
	inactive     bool
	rank         string
	summonderID  SummonerID
	leaguePoints int32
}

//LeagueEntry represents a LeagueEntryDTO
type LeagueEntry struct {
	queueType    string
	summonerName SummonerName
	hotStreak    bool
	miniSeries   *MiniSeries
	wins         int32
	veteran      bool
	losses       int32
	rank         string
	leagueID     LeagueID
	inactive     bool
	freshBlood   bool
	tier         Tier
	summonerID   SummonerID
	leaguePoints int32
}

//MiniSeries represents a MiniSeriesDTO
type MiniSeries struct {
	progress string
	losses   int32
	target   int32
	wins     int32
}
