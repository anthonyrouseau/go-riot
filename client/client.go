package client

import (
	"github.com/anthonyrouseau/go-riot/account"
	"github.com/anthonyrouseau/go-riot/lol"
	"github.com/anthonyrouseau/go-riot/queue"
	"github.com/anthonyrouseau/go-riot/shard"
	"github.com/anthonyrouseau/go-riot/summoner"
	"github.com/anthonyrouseau/go-riot/tft"
	"github.com/anthonyrouseau/go-riot/tournament"
	"golang.org/x/time/rate"
)

//Client is the interface whose methods will correspond to the Riot api routes
type Client interface {
	AllSummonerChampionMastery(summoner.ID) []*lol.ChampionMastery
	SummonerChampionMastery(summoner.ID, lol.ChampionID) *lol.ChampionMastery
	SummonerChampionMasteryScore(summoner.ID) int32
	ChampionRotations() *lol.Rotation
	LOLChallenger(queue.Name) *lol.LeagueInfo
	LOLSummonerLeagueEntries(summoner.ID) []*lol.LeagueEntry
	LOLAllLeagueEntries(queue.Name, lol.Tier, lol.Division) []*lol.LeagueEntry
	LOLGrandmaster(queue.Name) *lol.LeagueInfo
	LOLLeague(lol.LeagueID) *lol.LeagueInfo
	LOLMaster(queue.Name) *lol.LeagueInfo
	LOLMatch(lol.MatchID) *lol.Match
	Status() *shard.Status
	LOLAccountMatches(account.ID) *lol.MatchList
	LOLMatchTimeline(lol.MatchID) *lol.MatchTimeline
	ActiveGame(summoner.ID) *lol.CurrentGame
	FeaturedGames() *lol.FeaturedGames
	SummonerByAccount(account.ID) *summoner.Info
	SummonerByName(summoner.Name) *summoner.Info
	SummonerByPUUID(summoner.PUUID) *summoner.Info
	SummonerBySummonerID(summoner.ID) *summoner.Info
	TFTChallenger() *tft.LeagueList
	TFTSummonerLeagueEntries(summoner.ID) []*tft.LeagueEntry
	TFTAllLeagueEntries(tft.Tier, tft.Division) []*tft.LeagueEntry
	TFTGrandmaster() *tft.LeagueList
	TFTLeague(tft.LeagueID) *tft.LeagueList
	TFTMaster() *tft.LeagueList
	TFTMatchIDs(summoner.PUUID) []tft.MatchID
	TFTMatch(tft.MatchID) *tft.Match
	TFTSummonerByAccount(account.ID) *summoner.Info
	TFTSummonerByName(summoner.Name) *summoner.Info
	TFTSummonerByPUUID(summoner.PUUID) *summoner.Info
	TFTSummonerBySummonerID(summoner.ID) *summoner.Info
	ThirdPartyCode(summoner.ID) string
	TournamentMatchIDs(tournament.Code) []lol.MatchID
	TournamentMatch(lol.MatchID, tournament.Code) *lol.Match
	TournamentCodes() []tournament.Code
	TournamanetCodeInfo(tournament.Code) *tournament.CodeInfo
	UpdateTournamentCode(tournament.Code)
	LobbyEvents(tournament.Code) []*tournament.LobbyEvents
	TournamentProvider() int32
	Tournament() tournament.ID
}

//client is the internal implementation of the Client interface
//the client is has a limiter that will rate limit according to the Riot API
type client struct {
	limiter *rate.Limiter
}

func NewClient() *Client {
	return &client{}
}
