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
	url := fmt.Sprintf("https://%s.%s/lol/summoner/v4/summoners/by-account/%s", c.region, c.host, accountID)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	summonerInfo, err := c.getValue(ctx, req, routeLolSummonerByAccount, &summoner.Info{})
	if err != nil {
		return nil, err
	}
	return summonerInfo.(*summoner.Info), nil
}

func (c *client) LOLSummonerByName(ctx context.Context, sumName summoner.Name) (*summoner.Info, error) {
	url := fmt.Sprintf("https://%s.%s/lol/summoner/v4/summoners/by-name/%s", c.region, c.host, sumName)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	summonerInfo, err := c.getValue(ctx, req, routeLolSummonerByName, &summoner.Info{})
	if err != nil {
		return nil, err
	}
	return summonerInfo.(*summoner.Info), nil
}

func (c *client) LOLSummonerByPUUID(ctx context.Context, sumPUUID summoner.PUUID) (*summoner.Info, error) {
	url := fmt.Sprintf("https://%s.%s/lol/summoner/v4/summoners/by-puuid/%s", c.region, c.host, sumPUUID)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	summonerInfo, err := c.getValue(ctx, req, routeLolSummonerByPUUID, &summoner.Info{})
	if err != nil {
		return nil, err
	}
	return summonerInfo.(*summoner.Info), nil
}

func (c *client) LOLSummonerBySummonerID(ctx context.Context, sumID summoner.ID) (*summoner.Info, error) {
	url := fmt.Sprintf("https://%s.%s/lol/summoner/v4/summoners/%s", c.region, c.host, sumID)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	summonerInfo, err := c.getValue(ctx, req, routeLolSummonerBySummonerID, &summoner.Info{})
	if err != nil {
		return nil, err
	}
	return summonerInfo.(*summoner.Info), nil
}

func (c *client) LOLAllSummonerChampionMastery(ctx context.Context, sumID summoner.ID) ([]*lol.ChampionMastery, error) {
	url := fmt.Sprintf("https://%s.%s/lol/champion-mastery/v4/champion-masteries/by-summoner/%s", c.region, c.host, sumID)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	rec := []*lol.ChampionMastery{}
	masteries, err := c.getValue(ctx, req, routeLolAllSummonerChampionMastery, &rec)
	if err != nil {
		return nil, err
	}
	return *masteries.(*[]*lol.ChampionMastery), nil
}

func (c *client) LOLSummonerChampionMastery(ctx context.Context, sumID summoner.ID, champID lol.ChampionID) (*lol.ChampionMastery, error) {
	url := fmt.Sprintf("https://%s.%s/lol/champion-mastery/v4/champion-masteries/by-summoner/%s/by-champion/%d", c.region, c.host, sumID, champID)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	mastery, err := c.getValue(ctx, req, routeLolSummonerChampionMastery, &lol.ChampionMastery{})
	if err != nil {
		return nil, err
	}
	return mastery.(*lol.ChampionMastery), nil
}

func (c *client) LOLSummonerChampionMasteryScore(ctx context.Context, sumID summoner.ID) (int32, error) {
	url := fmt.Sprintf("https://%s.%s/lol/champion-mastery/v4/scores/by-summoner/%s", c.region, c.host, sumID)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return -1, err
	}
	var rec int32
	score, err := c.getValue(ctx, req, routeLolSummonerChampionMasteryScore, &rec)
	if err != nil {
		return -1, err
	}
	return *score.(*int32), nil
}

func (c *client) LOLChampionRotations(ctx context.Context) (*lol.Rotation, error) {
	url := fmt.Sprintf("https://%s.%s/lol/platform/v3/champion-rotations", c.region, c.host)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	rotations, err := c.getValue(ctx, req, routeLolChampionRotations, &lol.Rotation{})
	if err != nil {
		return nil, err
	}
	return rotations.(*lol.Rotation), nil
}

func (c *client) LOLFeaturedGames(ctx context.Context) (*lol.FeaturedGames, error) {
	url := fmt.Sprintf("https://%s.%s/lol/spectator/v4/featured-games", c.region, c.host)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	featured, err := c.getValue(ctx, req, routeLolFeaturedGames, &lol.FeaturedGames{})
	if err != nil {
		return nil, err
	}
	return featured.(*lol.FeaturedGames), nil
}

func (c *client) LOLActiveGame(ctx context.Context, sumID summoner.ID) (*lol.CurrentGame, error) {
	url := fmt.Sprintf("https://%s.%s/lol/spectator/v4/active-games/by-summoner/%s", c.region, c.host, sumID)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	active, err := c.getValue(ctx, req, routeLolActiveGame, &lol.CurrentGame{})
	if err != nil {
		return nil, err
	}
	return active.(*lol.CurrentGame), nil
}
