package client

import (
	"github.com/anthonyrouseau/go-riot/account"
	"github.com/anthonyrouseau/go-riot/lol"
	"github.com/anthonyrouseau/go-riot/queue"
	"github.com/anthonyrouseau/go-riot/shard"
)

//Client is the interface whose methods will correspond to the Riot api routes
type Client interface {
	AllSummonerChampionMastery(lol.SummonerID) []*lol.ChampionMastery
	SummonerChampionMastery(lol.SummonerID, lol.ChampionID) *lol.ChampionMastery
	SummonerChampionMasteryScore(lol.SummonerID) int32
	ChampionRotations() lol.Rotation
	ChallengerLeague(queue.Name) lol.LeagueInfo
	SummonerLeagueEntries(lol.SummonerID) []*lol.LeagueEntry
	AllLeagueEntries(queue.Name, lol.Tier, lol.Division) []*lol.LeagueEntry
	GrandmasterLeague(queue.Name) lol.LeagueInfo
	League(lol.LeagueID) lol.LeagueInfo
	MasterLeague(queue.Name) lol.LeagueInfo
	Status() shard.Status
	LolMatch(lol.MatchID) *lol.Match
	AccountMatches(account.ID) *lol.MatchList
	MatchTimeline(lol.MatchID) *lol.MatchTimeline
	TournamentMatchIDs(lol.TournamentCode) []lol.MatchID
	TournamentMatch(lol.MatchID, lol.TournamentCode) *lol.Match
	ActiveGame(lol.SummonerID) *lol.CurrentGame
	FeaturedGames() *lol.FeaturedGames
	SummonerByAccount(account.ID) *lol.Summoner
	SummonerByName(lol.SummonerName) *lol.Summoner
	SummonerByPUUID(lol.PUUID) *lol.Summoner
	SummonerBySummonerID(lol.SummonerID) *lol.Summoner
}
