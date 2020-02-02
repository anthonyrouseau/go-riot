package client

import (
	"context"

	"github.com/anthonyrouseau/go-riot/account"
	"github.com/anthonyrouseau/go-riot/summoner"
	"github.com/anthonyrouseau/go-riot/tft"
)

func (c *client) TFTChallenger(ctx context.Context) (*tft.LeagueList, error) {
	return nil, errUnimplemented
}

func (c *client) TFTSummonerLeagueEntries(ctx context.Context, sumID summoner.ID) ([]*tft.LeagueEntry, error) {
	return nil, errUnimplemented
}

func (c *client) TFTAllLeagueEntries(ctx context.Context, tier tft.Tier, division tft.Division) ([]*tft.LeagueEntry, error) {
	return nil, errUnimplemented
}

func (c *client) TFTGrandmaster(ctx context.Context) (*tft.LeagueList, error) {
	return nil, errUnimplemented
}

func (c *client) TFTLeague(ctx context.Context, leagueID tft.LeagueID) (*tft.LeagueList, error) {
	return nil, errUnimplemented
}

func (c *client) TFTMaster(ctx context.Context) (*tft.LeagueList, error) {
	return nil, errUnimplemented
}

func (c *client) TFTMatchIDs(ctx context.Context, sumID summoner.PUUID) ([]tft.MatchID, error) {
	return nil, errUnimplemented
}

func (c *client) TFTMatch(ctx context.Context, matchID tft.MatchID) (*tft.Match, error) {
	return nil, errUnimplemented
}

func (c *client) TFTSummonerByAccount(ctx context.Context, accountID account.ID) (*summoner.Info, error) {
	return nil, errUnimplemented
}

func (c *client) TFTSummonerByName(ctx context.Context, sumName summoner.Name) (*summoner.Info, error) {
	return nil, errUnimplemented
}

func (c *client) TFTSummonerByPUUID(ctx context.Context, sumID summoner.PUUID) (*summoner.Info, error) {
	return nil, errUnimplemented
}

func (c *client) TFTSummonerBySummonerID(ctx context.Context, sumID summoner.ID) (*summoner.Info, error) {
	return nil, errUnimplemented
}
