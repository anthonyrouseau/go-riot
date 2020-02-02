package client

import (
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
