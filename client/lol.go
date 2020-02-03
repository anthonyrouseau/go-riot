package client

import (
	"context"
	"fmt"
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
	leagueInfo, err := c.getValue(ctx, req, routeLolChallenger, &lol.LeagueInfo{})
	if err != nil {
		return nil, err
	}
	return leagueInfo.(*lol.LeagueInfo), nil
}

func (c *client) LOLSummonerLeagueEntries(ctx context.Context, sumID summoner.ID) ([]*lol.LeagueEntry, error) {
	url := fmt.Sprintf("https://%s.%s/lol/league/v4/entries/by-summoner/%s", c.region, c.host, sumID)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	rec := []*lol.LeagueEntry{}
	leagueEntries, err := c.getValue(ctx, req, routeLolSummonerLeagueEntries, &rec)
	if err != nil {
		return nil, err
	}
	return *leagueEntries.(*[]*lol.LeagueEntry), nil
}

func (c *client) LOLAllLeagueEntries(ctx context.Context, queueName queue.Name, tier lol.Tier, division lol.Division) ([]*lol.LeagueEntry, error) {
	url := fmt.Sprintf("https://%s.%s/lol/league/v4/entries/%s/%s/%s", c.region, c.host, queueName, tier, division)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	rec := []*lol.LeagueEntry{}
	leagueEntries, err := c.getValue(ctx, req, routeLoLAllLeagueEntries, &rec)
	if err != nil {
		return nil, err
	}
	return *leagueEntries.(*[]*lol.LeagueEntry), nil
}

func (c *client) LOLGrandmaster(ctx context.Context, queueName queue.Name) (*lol.LeagueInfo, error) {
	url := fmt.Sprintf("https://%s.%s/lol/league/v4/grandmasterleagues/by-queue/%s", c.region, c.host, queueName)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	leagueInfo, err := c.getValue(ctx, req, routeLolGrandmaster, &lol.LeagueInfo{})
	if err != nil {
		return nil, err
	}
	return leagueInfo.(*lol.LeagueInfo), nil
}

func (c *client) LOLLeague(ctx context.Context, leagueID lol.LeagueID) (*lol.LeagueInfo, error) {
	url := fmt.Sprintf("https://%s.%s/lol/league/v4/leagues/%s", c.region, c.host, leagueID)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	leagueInfo, err := c.getValue(ctx, req, routeLolLeague, &lol.LeagueInfo{})
	if err != nil {
		return nil, err
	}
	return leagueInfo.(*lol.LeagueInfo), nil
}

func (c *client) LOLMaster(ctx context.Context, queueName queue.Name) (*lol.LeagueInfo, error) {
	url := fmt.Sprintf("https://%s.%s/lol/league/v4/masterleagues/by-queue/%s", c.region, c.host, queueName)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	leagueInfo, err := c.getValue(ctx, req, routeLolChallenger, &lol.LeagueInfo{})
	if err != nil {
		return nil, err
	}
	return leagueInfo.(*lol.LeagueInfo), nil
}

func (c *client) LOLMatch(ctx context.Context, matchID lol.MatchID) (*lol.Match, error) {
	url := fmt.Sprintf("https://%s.%s/lol/match/v4/matches/%d", c.region, c.host, matchID)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	match, err := c.getValue(ctx, req, routeLolMatch, &lol.Match{})
	if err != nil {
		return nil, err
	}
	return match.(*lol.Match), nil
}

func (c *client) LOLAccountMatches(ctx context.Context, accountID account.ID) (*lol.MatchList, error) {
	url := fmt.Sprintf("https://%s.%s/lol/match/v4/matchlists/by-account/%s", c.region, c.host, accountID)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	matchList, err := c.getValue(ctx, req, routeLolAccountMatches, &lol.MatchList{})
	if err != nil {
		return nil, err
	}
	return matchList.(*lol.MatchList), nil
}

func (c *client) LOLMatchTimeline(ctx context.Context, matchID lol.MatchID) (*lol.MatchTimeline, error) {
	url := fmt.Sprintf("https://%s.%s/lol/match/v4/timelines/by-match/%d", c.region, c.host, matchID)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	matchTimeline, err := c.getValue(ctx, req, routeLolMatchTimeline, &lol.MatchTimeline{})
	if err != nil {
		return nil, err
	}
	return matchTimeline.(*lol.MatchTimeline), nil
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
