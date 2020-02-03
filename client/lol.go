package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/anthonyrouseau/go-riot/account"
	"github.com/anthonyrouseau/go-riot/lol"
	"github.com/anthonyrouseau/go-riot/queue"
	"github.com/anthonyrouseau/go-riot/summoner"
)

func (c *client) LOLChallenger(ctx context.Context, queueName queue.Name) (*lol.LeagueInfo, error) {
	url := fmt.Sprintf("https://%s.%s/lol/league/v4/challengerleagues/by-queue/%s", c.region, c.host, queueName)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.do(ctx, req, routeLolChallenger)
	if err != nil {
		return nil, err
	}
	err = c.checkResponse(resp)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	leagueInfo := &lol.LeagueInfo{}
	err = json.Unmarshal(body, leagueInfo)
	if err != nil {
		return nil, err
	}
	return leagueInfo, nil
}

func (c *client) LOLSummonerLeagueEntries(ctx context.Context, sumID summoner.ID) ([]*lol.LeagueEntry, error) {
	return nil, errUnimplemented
}

func (c *client) LOLAllLeagueEntries(ctx context.Context, queueName queue.Name, tier lol.Tier, division lol.Division) ([]*lol.LeagueEntry, error) {
	return nil, errUnimplemented
}

func (c *client) LOLGrandmaster(ctx context.Context, queueName queue.Name) (*lol.LeagueInfo, error) {
	return nil, errUnimplemented
}

func (c *client) LOLLeague(ctx context.Context, leagueID lol.LeagueID) (*lol.LeagueInfo, error) {
	return nil, errUnimplemented
}

func (c *client) LOLMaster(ctx context.Context, queueName queue.Name) (*lol.LeagueInfo, error) {
	return nil, errUnimplemented
}

func (c *client) LOLMatch(ctx context.Context, matchID lol.MatchID) (*lol.Match, error) {
	return nil, errUnimplemented
}

func (c *client) LOLAccountMatches(ctx context.Context, accountID account.ID) (*lol.MatchList, error) {
	return nil, errUnimplemented
}

func (c *client) LOLMatchTimeline(ctx context.Context, matchID lol.MatchID) (*lol.MatchTimeline, error) {
	return nil, errUnimplemented
}

func (c *client) LOLSummonerByAccount(ctx context.Context, accountID account.ID) (*summoner.Info, error) {
	return nil, errUnimplemented
}

func (c *client) LOLSummonerByName(ctx context.Context, sumName summoner.Name) (*summoner.Info, error) {
	return nil, errUnimplemented
}

func (c *client) LOLSummonerByPUUID(ctx context.Context, sumPUUID summoner.PUUID) (*summoner.Info, error) {
	return nil, errUnimplemented
}

func (c *client) LOLSummonerBySummonerID(ctx context.Context, sumID summoner.ID) (*summoner.Info, error) {
	return nil, errUnimplemented
}

func (c *client) LOLAllSummonerChampionMastery(ctx context.Context, sumID summoner.ID) ([]*lol.ChampionMastery, error) {
	return nil, errUnimplemented
}

func (c *client) LOLSummonerChampionMastery(ctx context.Context, sumID summoner.ID, champID lol.ChampionID) (*lol.ChampionMastery, error) {
	return nil, errUnimplemented
}

func (c *client) LOLSummonerChampionMasteryScore(ctx context.Context, sumID summoner.ID) (int32, error) {
	return -1, errUnimplemented
}

func (c *client) LOLChampionRotations(ctx context.Context) (*lol.Rotation, error) {
	return nil, errUnimplemented
}

func (c *client) LOLFeaturedGames(ctx context.Context) (*lol.FeaturedGames, error) {
	return nil, errUnimplemented
}

func (c *client) LOLActiveGame(ctx context.Context, sumID summoner.ID) (*lol.CurrentGame, error) {
	return nil, errUnimplemented
}
