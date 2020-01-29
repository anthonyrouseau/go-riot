package lol

import (
	"github.com/anthonyrouseau/go-riot/queue"
	"github.com/anthonyrouseau/go-riot/summoner"
)

//LeagueID is the Id for the league for a given queue
type LeagueID string

//Tier is the tier within a league
type Tier string

//Division is the division within a league
type Division string

//LeagueInfo represents a LeagueListDTO
type LeagueInfo struct {
	ID      LeagueID
	Tier    Tier
	Queue   queue.Name
	Name    string
	Entries []*LeagueItem
}

//LeagueItem represents a LeagueItemDTO
type LeagueItem struct {
	SummonerName summoner.Name
	HotStreak    bool
	MiniSeries   *MiniSeries
	Wins         int32
	Veteran      bool
	Losses       int32
	FreshBlood   bool
	Inactive     bool
	Rank         string
	SummonerID   summoner.ID
	LeaguePoints int32
}

//LeagueEntry represents a LeagueEntryDTO
type LeagueEntry struct {
	QueueType    string
	SummonerName summoner.Name
	HotStreak    bool
	MiniSeries   *MiniSeries
	Wins         int32
	Veteran      bool
	Losses       int32
	Rank         string
	LeagueID     LeagueID
	Inactive     bool
	FreshBlood   bool
	Tier         Tier
	SummonerID   summoner.ID
	LeaguePoints int32
}

//MiniSeries represents a MiniSeriesDTO
type MiniSeries struct {
	Progress string
	Losses   int32
	Target   int32
	Wins     int32
}
