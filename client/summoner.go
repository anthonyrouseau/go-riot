package client

import (
	"github.com/anthonyrouseau/go-riot/account"
	"github.com/anthonyrouseau/go-riot/lol"
	"github.com/anthonyrouseau/go-riot/summoner"
)

func (c *client) SummonerByAccount(account.ID) (*summoner.Info, error) {
	return nil, errUnimplemented
}

func (c *client) SummonerByName(summoner.Name) (*summoner.Info, error) {
	return nil, errUnimplemented
}

func (c *client) SummonerByPUUID(summoner.PUUID) (*summoner.Info, error) {
	return nil, errUnimplemented
}

func (c *client) SummonerBySummonerID(summoner.ID) (*summoner.Info, error) {
	return nil, errUnimplemented
}

func (c *client) AllSummonerChampionMastery(summoner.ID) ([]*lol.ChampionMastery, error) {
	return nil, errUnimplemented
}

func (c *client) SummonerChampionMastery(summoner.ID, lol.ChampionID) (*lol.ChampionMastery, error) {
	return nil, errUnimplemented
}

func (c *client) SummonerChampionMasteryScore(summoner.ID) (int32, error) {
	return -1, errUnimplemented
}

func (c *client) ActiveGame(summoner.ID) (*lol.CurrentGame, error) {
	return nil, errUnimplemented
}
