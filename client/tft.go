package client

import (
	"github.com/anthonyrouseau/go-riot/account"
	"github.com/anthonyrouseau/go-riot/summoner"
	"github.com/anthonyrouseau/go-riot/tft"
)

func (c *client) TFTChallenger() (*tft.LeagueList, error) {

}

func (c *client) TFTSummonerLeagueEntries(summoner.ID) ([]*tft.LeagueEntry, error) {

}

func (c *client) TFTAllLeagueEntries(tft.Tier, tft.Division) ([]*tft.LeagueEntry, error) {

}

func (c *client) TFTGrandmaster() (*tft.LeagueList, error) {

}

func (c *client) TFTLeague(tft.LeagueID) (*tft.LeagueList, error) {

}

func (c *client) TFTMaster() (*tft.LeagueList, error) {

}

func (c *client) TFTMatchIDs(summoner.PUUID) ([]tft.MatchID, error) {

}

func (c *client) TFTMatch(tft.MatchID) (*tft.Match, error) {

}

func (c *client) TFTSummonerByAccount(account.ID) (*summoner.Info, error) {

}

func (c *client) TFTSummonerByName(summoner.Name) (*summoner.Info, error) {

}

func (c *client) TFTSummonerByPUUID(summoner.PUUID) (*summoner.Info, error) {

}

func (c *client) TFTSummonerBySummonerID(summoner.ID) (*summoner.Info, error) {

}
