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

//Client is an interface with methods corresponding to the Riot api routes
type Client interface {
	LOLAllSummonerChampionMastery(summoner.ID) ([]*lol.ChampionMastery, error)
	LOLSummonerChampionMastery(summoner.ID, lol.ChampionID) (*lol.ChampionMastery, error)
	LOLSummonerChampionMasteryScore(summoner.ID) (int32, error)
	LOLChampionRotations() (*lol.Rotation, error)
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
	LOLActiveGame(summoner.ID) (*lol.CurrentGame, error)
	LOLFeaturedGames() (*lol.FeaturedGames, error)
	LOLSummonerByAccount(account.ID) (*summoner.Info, error)
	LOLSummonerByName(summoner.Name) (*summoner.Info, error)
	LOLSummonerByPUUID(summoner.PUUID) (*summoner.Info, error)
	LOLSummonerBySummonerID(summoner.ID) (*summoner.Info, error)
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
	limiters map[routeKey]*rate.Limiter
}

//NewClient returns a basic Client for use with the Riot API
func NewClient() (Client, error) {
	return &client{}, nil
}
