package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/anthonyrouseau/go-riot/account"
	"github.com/anthonyrouseau/go-riot/lol"
	"github.com/anthonyrouseau/go-riot/queue"
	"github.com/anthonyrouseau/go-riot/shard"
	"github.com/anthonyrouseau/go-riot/summoner"
	"github.com/anthonyrouseau/go-riot/tft"
	"github.com/anthonyrouseau/go-riot/tournament"
	"github.com/pkg/errors"
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
	LOLAllLeagueEntries(ctx context.Context, queue queue.Name, tier lol.Tier, division lol.Division, options ...lol.LeagueQueryOption) ([]*lol.LeagueEntry, error)
	LOLGrandmaster(context.Context, queue.Name) (*lol.LeagueInfo, error)
	LOLLeague(context.Context, lol.LeagueID) (*lol.LeagueInfo, error)
	LOLMaster(context.Context, queue.Name) (*lol.LeagueInfo, error)
	LOLMatch(context.Context, lol.MatchID) (*lol.Match, error)
	Status(context.Context) (*shard.Status, error)
	LOLAccountMatches(ctx context.Context, acc account.ID, options ...lol.MatchOption) (*lol.MatchList, error)
	LOLMatchTimeline(context.Context, lol.MatchID) (*lol.MatchTimeline, error)
	LOLActiveGame(context.Context, summoner.ID) (*lol.CurrentGame, error)
	LOLFeaturedGames(context.Context) (*lol.FeaturedGames, error)
	LOLSummonerByAccount(context.Context, account.ID) (*summoner.Info, error)
	LOLSummonerByName(context.Context, summoner.Name) (*summoner.Info, error)
	LOLSummonerByPUUID(context.Context, summoner.PUUID) (*summoner.Info, error)
	LOLSummonerBySummonerID(context.Context, summoner.ID) (*summoner.Info, error)
	TFTChallenger(context.Context) (*tft.LeagueList, error)
	TFTSummonerLeagueEntries(context.Context, summoner.ID) ([]*tft.LeagueEntry, error)
	TFTAllLeagueEntries(ctx context.Context, tier tft.Tier, division tft.Division, options ...tft.LeagueQueryOption) ([]*tft.LeagueEntry, error)
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
	TournamentCodes(ctx context.Context, tournamentID tournament.ID, body *tournament.CodeRequestBody, options ...tournament.CodeRequestOption) ([]tournament.Code, error)
	TournamentCodeInfo(context.Context, tournament.Code) (*tournament.CodeInfo, error)
	UpdateTournamentCode(context.Context, tournament.Code) error
	LobbyEvents(context.Context, tournament.Code) (*tournament.LobbyEvents, error)
	TournamentProvider(ctx context.Context, region tournament.Region, url string) (int32, error)
	Tournament(ctx context.Context, providerID int32, options ...tournament.RegistrationOption) (tournament.ID, error)
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
	host     string
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
		host:     "api.riotgames.com",
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

//checkResponse checks the response from the riot api and returns the body
func (c *client) checkResponse(resp *http.Response) error {
	if code := resp.StatusCode; code != 200 {
		if code == 400 {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			newError := &riotError{}
			err = json.Unmarshal(body, newError)
			if err != nil {
				return err
			}
			return newError.Format()
		}
		if err, ok := httpCodes[code]; ok {
			return err
		}
		return errors.Errorf("An unexpected error response from Riot with code: %d", code)
	}
	return nil
}

//handleResponse takes a response, checks if ok, then unmarshals into the rec
//(rec must be a non-nil pointer since it is being provided to json.Unmarshal)
func (c *client) handleResponse(resp *http.Response, rec interface{}) (interface{}, error) {
	rv := reflect.ValueOf(rec)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return nil, errors.New("Value passed to handleResponse is not valid. Must be non-nil pointer")
	}
	err := c.checkResponse(resp)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		fmt.Println(body)
	}
	err = json.Unmarshal(body, rec)
	if err != nil {
		return nil, err
	}
	return rec, nil
}

//getValue wraps the client methods do and handleResponse
//(rec must be a non-nil pointer since it is being provided to json.Unmarshal)
func (c *client) getValue(ctx context.Context, req *http.Request, routeKey routeKey, rec interface{}) (interface{}, error) {
	resp, err := c.do(ctx, req, routeKey)
	if err != nil {
		return nil, err
	}
	value, err := c.handleResponse(resp, rec)
	if err != nil {
		return nil, err
	}
	return value, nil
}

//regionalRouting takes a platform routing value (e.g. na1) and returns a regional routing value (e.g. americas)
//some routes use regional routing rather than platform routing (e.g. tft match routes, and tournament routes)
func regionalRouting(r string) string {
	switch r {
	case "na1":
		fallthrough
	case "la1":
		fallthrough
	case "la2":
		fallthrough
	case "br1":
		fallthrough
	case "oc1":
		return "americas"
	case "jp1":
		fallthrough
	case "kr":
		return "asia"
	case "eun1":
		fallthrough
	case "euw1":
		fallthrough
	case "tr1":
		fallthrough
	case "ru":
		return "europe"
	}
	return "americas"
}
