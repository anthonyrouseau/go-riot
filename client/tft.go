package client

import (
	"github.com/anthonyrouseau/go-riot/account"
	"github.com/anthonyrouseau/go-riot/summoner"
	"github.com/anthonyrouseau/go-riot/tft"
)

func (c *client) TFTChallenger() (*tft.LeagueList, error) {
	return nil, errUnimplemented
}

func (c *client) TFTSummonerLeagueEntries(summoner.ID) ([]*tft.LeagueEntry, error) {
	return nil, errUnimplemented
}

func (c *client) TFTAllLeagueEntries(tft.Tier, tft.Division) ([]*tft.LeagueEntry, error) {
	return nil, errUnimplemented
}

func (c *client) TFTGrandmaster() (*tft.LeagueList, error) {
	return nil, errUnimplemented
}

func (c *client) TFTLeague(tft.LeagueID) (*tft.LeagueList, error) {
	return nil, errUnimplemented
}

func (c *client) TFTMaster() (*tft.LeagueList, error) {
	return nil, errUnimplemented
}

func (c *client) TFTMatchIDs(summoner.PUUID) ([]tft.MatchID, error) {
	return nil, errUnimplemented
}

func (c *client) TFTMatch(tft.MatchID) (*tft.Match, error) {
	return nil, errUnimplemented
}

func (c *client) TFTSummonerByAccount(account.ID) (*summoner.Info, error) {
	return nil, errUnimplemented
}

func (c *client) TFTSummonerByName(summoner.Name) (*summoner.Info, error) {
	return nil, errUnimplemented
}

func (c *client) TFTSummonerByPUUID(summoner.PUUID) (*summoner.Info, error) {
	return nil, errUnimplemented
}

func (c *client) TFTSummonerBySummonerID(summoner.ID) (*summoner.Info, error) {
	return nil, errUnimplemented
}
