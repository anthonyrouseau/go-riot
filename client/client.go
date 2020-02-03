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
	Region() string
}

//client wraps an http client and implements the Client interface.
//The client may have a limiters map that will rate limit according to the Riot API guidelines for that client variant
//A client can only have one region associated with it
type client struct {
	limiters map[routeKey][]*rate.Limiter
	variant  variant
	apiKey   string
	client   *http.Client
	region   string
}

//Option is a function used to alter the default client during creation
type Option func(*client)

//SetRegion returns a ClientOption which sets the clients reigon to the given string
func SetRegion(region string) Option {
	return func(c *client) {
		c.region = region
	}
}

//SetVariant returns a ClientOption which sets the client variant and makes any other necessary changes for the type of client
func SetVariant(v variant) Option {
	return func(c *client) {
		c.variant = v
	}
}

//Variant returns the variant of the client
func (c *client) Variant() variant {
	return c.variant
}

//APIKey returns the api key associated with the client
func (c *client) APIKey() string {
	return c.apiKey
}

//Region returns the region of the client
func (c *client) Region() string {
	return c.region
}

//NewClient returns a default Client unless options are specified
func NewClient(key string, options ...Option) (Client, error) {
	c := &http.Client{}
	newClient := &client{
		variant:  unspecifiedClient,
		apiKey:   key,
		client:   c,
		region:   "na1",
		limiters: make(map[routeKey][]*rate.Limiter),
	}
	for _, opt := range options {
		opt(newClient)
	}
	return newClient, nil
}

//do wraps the http.Do method and adds the APIKey to the request header
func (c *client) do(ctx context.Context, req *http.Request, key routeKey) (*http.Response, error) {
	err := c.rateLimit(ctx, key)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Riot-Token", c.APIKey())
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
