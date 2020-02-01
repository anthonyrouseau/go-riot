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
	AllSummonerChampionMastery(summoner.ID) ([]*lol.ChampionMastery, error)
	SummonerChampionMastery(summoner.ID, lol.ChampionID) (*lol.ChampionMastery, error)
	SummonerChampionMasteryScore(summoner.ID) (int32, error)
	ChampionRotations() (*lol.Rotation, error)
	LOLChallenger(queue.Name) (*lol.LeagueInfo, error)
	LOLSummonerLeagueEntries(summoner.ID) ([]*lol.LeagueEntry, error)
	LOLAllLeagueEntries(queue.Name, lol.Tier, lol.Division) ([]*lol.LeagueEntry, error)
	LOLGrandmaster(queue.Name) (*lol.LeagueInfo, error)
	LOLLeague(lol.LeagueID) (*lol.LeagueInfo, error)
	LOLMaster(queue.Name) (*lol.LeagueInfo, error)
	LOLMatch(lol.MatchID) (*lol.Match, error)
	Status() (*shard.Status, error)
	LOLAccountMatches(account.ID) (*lol.MatchList, error)
	LOLMatchTimeline(lol.MatchID) (*lol.MatchTimeline, error)
	ActiveGame(summoner.ID) (*lol.CurrentGame, error)
	FeaturedGames() (*lol.FeaturedGames, error)
	SummonerByAccount(account.ID) (*summoner.Info, error)
	SummonerByName(summoner.Name) (*summoner.Info, error)
	SummonerByPUUID(summoner.PUUID) (*summoner.Info, error)
	SummonerBySummonerID(summoner.ID) (*summoner.Info, error)
	TFTChallenger() (*tft.LeagueList, error)
	TFTSummonerLeagueEntries(summoner.ID) ([]*tft.LeagueEntry, error)
	TFTAllLeagueEntries(tft.Tier, tft.Division) ([]*tft.LeagueEntry, error)
	TFTGrandmaster() (*tft.LeagueList, error)
	TFTLeague(tft.LeagueID) (*tft.LeagueList, error)
	TFTMaster() (*tft.LeagueList, error)
	TFTMatchIDs(summoner.PUUID) ([]tft.MatchID, error)
	TFTMatch(tft.MatchID) (*tft.Match, error)
	TFTSummonerByAccount(account.ID) (*summoner.Info, error)
	TFTSummonerByName(summoner.Name) (*summoner.Info, error)
	TFTSummonerByPUUID(summoner.PUUID) (*summoner.Info, error)
	TFTSummonerBySummonerID(summoner.ID) (*summoner.Info, error)
	ThirdPartyCode(summoner.ID) (string, error)
	TournamentMatchIDs(tournament.Code) ([]lol.MatchID, error)
	TournamentMatch(lol.MatchID, tournament.Code) (*lol.Match, error)
	TournamentCodes() ([]tournament.Code, error)
	TournamanetCodeInfo(tournament.Code) (*tournament.CodeInfo, error)
	UpdateTournamentCode(tournament.Code) error
	LobbyEvents(tournament.Code) ([]*tournament.LobbyEvents, error)
	TournamentProvider() (int32, error)
	Tournament() (tournament.ID, error)
}

//client is the internal implementation of the Client interface
//the client is has a limiter that will rate limit according to the Riot API
type client struct {
	limiter *rate.Limiter
}

//NewClient returns a basic Client for use with the Riot API
func NewClient() Client {
	return &client{}
}
