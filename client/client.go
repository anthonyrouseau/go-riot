package client

import (
	"context"
	"net/http"

	"github.com/anthonyrouseau/go-riot/account"
	"github.com/anthonyrouseau/go-riot/lol"
	"github.com/anthonyrouseau/go-riot/queue"
	"github.com/anthonyrouseau/go-riot/shard"
	"github.com/anthonyrouseau/go-riot/summoner"
	"github.com/anthonyrouseau/go-riot/tft"
	"github.com/anthonyrouseau/go-riot/tournament"
	"golang.org/x/time/rate"
)

//variant is the type of client e.g. dev, production, etc. used for rate limiting
type variant int

const (
	unspecifiedClient variant = iota
	devClient
	personalClient
	productionClient
)

//Client is an interface with methods corresponding to the Riot api routes
type Client interface {
	LOLAllSummonerChampionMastery(context.Context, summoner.ID) ([]*lol.ChampionMastery, error)
	LOLSummonerChampionMastery(context.Context, summoner.ID, lol.ChampionID) (*lol.ChampionMastery, error)
	LOLSummonerChampionMasteryScore(context.Context, summoner.ID) (int32, error)
	LOLChampionRotations(context.Context) (*lol.Rotation, error)
	LOLChallenger(context.Context, queue.Name) (*lol.LeagueInfo, error)
	LOLSummonerLeagueEntries(context.Context, summoner.ID) ([]*lol.LeagueEntry, error)
	LOLAllLeagueEntries(context.Context, queue.Name, lol.Tier, lol.Division) ([]*lol.LeagueEntry, error)
	LOLGrandmaster(context.Context, queue.Name) (*lol.LeagueInfo, error)
	LOLLeague(context.Context, lol.LeagueID) (*lol.LeagueInfo, error)
	LOLMaster(context.Context, queue.Name) (*lol.LeagueInfo, error)
	LOLMatch(context.Context, lol.MatchID) (*lol.Match, error)
	Status(context.Context) (*shard.Status, error)
	LOLAccountMatches(context.Context, account.ID) (*lol.MatchList, error)
	LOLMatchTimeline(context.Context, lol.MatchID) (*lol.MatchTimeline, error)
	LOLActiveGame(context.Context, summoner.ID) (*lol.CurrentGame, error)
	LOLFeaturedGames(context.Context) (*lol.FeaturedGames, error)
	LOLSummonerByAccount(context.Context, account.ID) (*summoner.Info, error)
	LOLSummonerByName(context.Context, summoner.Name) (*summoner.Info, error)
	LOLSummonerByPUUID(context.Context, summoner.PUUID) (*summoner.Info, error)
	LOLSummonerBySummonerID(context.Context, summoner.ID) (*summoner.Info, error)
	TFTChallenger(context.Context) (*tft.LeagueList, error)
	TFTSummonerLeagueEntries(context.Context, summoner.ID) ([]*tft.LeagueEntry, error)
	TFTAllLeagueEntries(context.Context, tft.Tier, tft.Division) ([]*tft.LeagueEntry, error)
	TFTGrandmaster(context.Context) (*tft.LeagueList, error)
	TFTLeague(context.Context, tft.LeagueID) (*tft.LeagueList, error)
	TFTMaster(context.Context) (*tft.LeagueList, error)
	TFTMatchIDs(context.Context, summoner.PUUID) ([]tft.MatchID, error)
	TFTMatch(context.Context, tft.MatchID) (*tft.Match, error)
	TFTSummonerByAccount(context.Context, account.ID) (*summoner.Info, error)
	TFTSummonerByName(context.Context, summoner.Name) (*summoner.Info, error)
	TFTSummonerByPUUID(context.Context, summoner.PUUID) (*summoner.Info, error)
	TFTSummonerBySummonerID(context.Context, summoner.ID) (*summoner.Info, error)
	ThirdPartyCode(context.Context, summoner.ID) (string, error)
	TournamentMatchIDs(context.Context, tournament.Code) ([]lol.MatchID, error)
	TournamentMatch(context.Context, lol.MatchID, tournament.Code) (*lol.Match, error)
	TournamentCodes(context.Context) ([]tournament.Code, error)
	TournamanetCodeInfo(context.Context, tournament.Code) (*tournament.CodeInfo, error)
	UpdateTournamentCode(context.Context, tournament.Code) error
	LobbyEvents(context.Context, tournament.Code) ([]*tournament.LobbyEvents, error)
	TournamentProvider(context.Context) (int32, error)
	Tournament(context.Context) (tournament.ID, error)
	Variant() variant
	APIKey() string
}

//client wraps an http client and implements the Client interface.
//The client may have a limiters map that will rate limit according to the Riot API guidelines for that client variant
type client struct {
	limiters map[routeKey][]*rate.Limiter
	variant  variant
	apiKey   string
	client   *http.Client
}

//Variant returns the variant of the client
func (c *client) Variant() variant {
	return c.variant
}

//APIKey returns the api key associated with the client
func (c *client) APIKey() string {
	return c.apiKey
}

//NewClient returns a Client with variant as unspecified
func NewClient(key string) (Client, error) {
	c := &http.Client{}
	return &client{variant: unspecifiedClient, apiKey: key, client: c, limiters: make(map[routeKey][]*rate.Limiter)}, nil
}

//NewDevClient returns a Client with variant dev
func NewDevClient(key string) (Client, error) {
	c := &http.Client{}
	return &client{variant: devClient, apiKey: key, client: c, limiters: make(map[routeKey][]*rate.Limiter)}, nil
}

//NewPersonalClient returns a Client with variant personal
func NewPersonalClient(key string) (Client, error) {
	c := &http.Client{}
	return &client{variant: personalClient, apiKey: key, client: c, limiters: make(map[routeKey][]*rate.Limiter)}, nil
}

//NewProductionClient returns a Client with variant production
func NewProductionClient(key string) (Client, error) {
	c := &http.Client{}
	return &client{variant: productionClient, apiKey: key, client: c, limiters: make(map[routeKey][]*rate.Limiter)}, nil
}

//do wraps the http.Do method and adds the APIKey to the request header
func (c *client) do(req *http.Request) (*http.Response, error) {
	req.Header.Add("X-Riot-Token", c.APIKey())
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
