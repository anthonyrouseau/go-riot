package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/anthonyrouseau/go-riot/account"
	"github.com/anthonyrouseau/go-riot/summoner"
	"github.com/anthonyrouseau/go-riot/tft"
)

func (c *client) TFTChallenger(ctx context.Context) (*tft.LeagueList, error) {
	url := fmt.Sprintf("https://%s.%s/tft/league/v1/challenger", c.region, c.host)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	leagueList, err := c.getValue(ctx, req, routeTftChallenger, &tft.LeagueList{})
	if err != nil {
		return nil, err
	}
	return leagueList.(*tft.LeagueList), nil
}

func (c *client) TFTSummonerLeagueEntries(ctx context.Context, sumID summoner.ID) ([]*tft.LeagueEntry, error) {
	url := fmt.Sprintf("https://%s.%s/tft/league/v1/entries/by-summoner/%s", c.region, c.host, sumID)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	rec := []*tft.LeagueEntry{}
	leagueEntry, err := c.getValue(ctx, req, routeTftSummonerLeagueEntries, &rec)
	if err != nil {
		return nil, err
	}
	return *leagueEntry.(*[]*tft.LeagueEntry), nil
}

func (c *client) TFTAllLeagueEntries(ctx context.Context, tier tft.Tier, division tft.Division) ([]*tft.LeagueEntry, error) {
	url := fmt.Sprintf("https://%s.%s/tft/league/v1/entries/%s/%s", c.region, c.host, tier, division)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	rec := []*tft.LeagueEntry{}
	leagueEntry, err := c.getValue(ctx, req, routeTftAllLeagueEntries, &rec)
	if err != nil {
		return nil, err
	}
	return *leagueEntry.(*[]*tft.LeagueEntry), nil
}

func (c *client) TFTGrandmaster(ctx context.Context) (*tft.LeagueList, error) {
	url := fmt.Sprintf("https://%s.%s/tft/league/v1/grandmaster", c.region, c.host)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	leagueList, err := c.getValue(ctx, req, routeTftGrandmaster, &tft.LeagueList{})
	if err != nil {
		return nil, err
	}
	return leagueList.(*tft.LeagueList), nil
}

func (c *client) TFTLeague(ctx context.Context, leagueID tft.LeagueID) (*tft.LeagueList, error) {
	url := fmt.Sprintf("https://%s.%s/tft/league/v1/leagues/%s", c.region, c.host, leagueID)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	leagueList, err := c.getValue(ctx, req, routeTftLeague, &tft.LeagueList{})
	if err != nil {
		return nil, err
	}
	return leagueList.(*tft.LeagueList), nil
}

func (c *client) TFTMaster(ctx context.Context) (*tft.LeagueList, error) {
	url := fmt.Sprintf("https://%s.%s/tft/league/v1/master", c.region, c.host)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	leagueList, err := c.getValue(ctx, req, routeTftMaster, &tft.LeagueList{})
	if err != nil {
		return nil, err
	}
	return leagueList.(*tft.LeagueList), nil
}

func (c *client) TFTMatchIDs(ctx context.Context, sumID summoner.PUUID) ([]tft.MatchID, error) {
	regRoute := regionalRouting(c.region)
	url := fmt.Sprintf("https://%s.%s/tft/match/v1/matches/by-puuid/%s/ids", regRoute, c.host, sumID)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	var rec []tft.MatchID
	matchList, err := c.getValue(ctx, req, routeTftMatchIDs, &rec)
	if err != nil {
		return nil, err
	}
	return *matchList.(*[]tft.MatchID), nil
}

func (c *client) TFTMatch(ctx context.Context, matchID tft.MatchID) (*tft.Match, error) {
	regRoute := regionalRouting(c.region)
	url := fmt.Sprintf("https://%s.%s/tft/match/v1/matches/%s", regRoute, c.host, matchID)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	match, err := c.getValue(ctx, req, routeTftMatch, &tft.Match{})
	if err != nil {
		return nil, err
	}
	return match.(*tft.Match), nil
}

func (c *client) TFTSummonerByAccount(ctx context.Context, accountID account.ID) (*summoner.Info, error) {
	url := fmt.Sprintf("https://%s.%s/tft/summoner/v1/summoners/by-account/%s", c.region, c.host, accountID)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	summonerInfo, err := c.getValue(ctx, req, routeTftSummonerByAccount, &summoner.Info{})
	if err != nil {
		return nil, err
	}
	return summonerInfo.(*summoner.Info), nil
}

func (c *client) TFTSummonerByName(ctx context.Context, sumName summoner.Name) (*summoner.Info, error) {
	url := fmt.Sprintf("https://%s.%s/tft/summoner/v1/summoners/by-name/%s", c.region, c.host, sumName)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	summonerInfo, err := c.getValue(ctx, req, routeTftSummonerByName, &summoner.Info{})
	if err != nil {
		return nil, err
	}
	return summonerInfo.(*summoner.Info), nil
}

func (c *client) TFTSummonerByPUUID(ctx context.Context, sumID summoner.PUUID) (*summoner.Info, error) {
	url := fmt.Sprintf("https://%s.%s/tft/summoner/v1/summoners/by-puuid/%s", c.region, c.host, sumID)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	summonerInfo, err := c.getValue(ctx, req, routeTftSummonerByPUUID, &summoner.Info{})
	if err != nil {
		return nil, err
	}
	return summonerInfo.(*summoner.Info), nil
}

func (c *client) TFTSummonerBySummonerID(ctx context.Context, sumID summoner.ID) (*summoner.Info, error) {
	url := fmt.Sprintf("https://%s.%s/tft/summoner/v1/summoners/%s", c.region, c.host, sumID)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	summonerInfo, err := c.getValue(ctx, req, routeTftSummonerBySummonerID, &summoner.Info{})
	if err != nil {
		return nil, err
	}
	return summonerInfo.(*summoner.Info), nil
}

//regionalRouting takes a platform routing value (e.g. na1) and returns a regional routing value (e.g. americas)
//some routes use regional routing rather than platform routing (e.g. tft match routes)
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
