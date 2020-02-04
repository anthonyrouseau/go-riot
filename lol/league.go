package lol

import (
	"github.com/anthonyrouseau/go-riot/queue"
	"github.com/anthonyrouseau/go-riot/summoner"
)

//LeagueID is the Id for the league for a given queue
type LeagueID string

//Tier is the tier within a league
type Tier string

//These constants are the values to be used for tiers
const (
	Iron     Tier = "IRON"
	Bronze   Tier = "BRONZE"
	Silver   Tier = "SILVER"
	Gold     Tier = "GOLD"
	Platinum Tier = "PLATINUM"
	Diamond  Tier = "DIAMOND"
)

//Division is the division within a league
type Division string

//These constants are the values to be used for divisons
const (
	I   Division = "I"
	II  Division = "II"
	III Division = "III"
	IV  Division = "IV"
)

//LeagueInfo represents a LeagueListDTO
type LeagueInfo struct {
	ID      LeagueID      `json:"leagueId"`
	Tier    Tier          `json:"tier"`
	Queue   queue.Name    `json:"queue"`
	Name    string        `json:"name"`
	Entries []*LeagueItem `json:"entries"`
}

//LeagueItem represents a LeagueItemDTO
type LeagueItem struct {
	SummonerName summoner.Name `json:"summonerName"`
	HotStreak    bool          `json:"hotStreak"`
	MiniSeries   *MiniSeries   `json:"miniSeries"`
	Wins         int32         `json:"wins"`
	Veteran      bool          `json:"veteran"`
	Losses       int32         `json:"losses"`
	FreshBlood   bool          `json:"freshBlood"`
	Inactive     bool          `json:"inactive"`
	Rank         string        `json:"rank"`
	SummonerID   summoner.ID   `json:"summonerId"`
	LeaguePoints int32         `json:"leaguePoints"`
}

//LeagueEntry represents a LeagueEntryDTO
type LeagueEntry struct {
	QueueType    string        `json:"queueType"`
	SummonerName summoner.Name `json:"summonerName"`
	HotStreak    bool          `json:"hotStreak"`
	MiniSeries   *MiniSeries   `json:"miniSeries"`
	Wins         int32         `json:"wins"`
	Veteran      bool          `json:"veteran"`
	Losses       int32         `json:"losses"`
	Rank         string        `json:"rank"`
	LeagueID     LeagueID      `json:"leagueId"`
	Inactive     bool          `json:"inactive"`
	FreshBlood   bool          `json:"freshBlood"`
	Tier         Tier          `json:"tier"`
	SummonerID   summoner.ID   `json:"summonerId"`
	LeaguePoints int32         `json:"leaguePoints"`
}

//MiniSeries represents a MiniSeriesDTO
type MiniSeries struct {
	Progress string `json:"progress"`
	Losses   int32  `json:"losses"`
	Target   int32  `json:"target"`
	Wins     int32  `json:"wins"`
}

//LeagueQueryParams are the optional query parameters
type LeagueQueryParams struct {
	Page int32
}

//LeagueQueryOption is a function that changes the default value of a LeagueQueryParam
type LeagueQueryOption func(*LeagueQueryParams)

//PageOption is a LeagueQueryOption that updates the page parameter
func PageOption(n int32) LeagueQueryOption {
	return func(l *LeagueQueryParams) {
		l.Page = n
	}
}
