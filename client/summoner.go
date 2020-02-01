package client

import (
	"github.com/anthonyrouseau/go-riot/account"
	"github.com/anthonyrouseau/go-riot/lol"
	"github.com/anthonyrouseau/go-riot/summoner"
)

func (c *client) SummonerByAccount(account.ID) (*summoner.Info, error) {

}

func (c *client) SummonerByName(summoner.Name) (*summoner.Info, error) {

}

func (c *client) SummonerByPUUID(summoner.PUUID) (*summoner.Info, error) {

}

func (c *client) SummonerBySummonerID(summoner.ID) (*summoner.Info, error) {

}

func (c *client) AllSummonerChampionMastery(summoner.ID) ([]*lol.ChampionMastery, error) {

}

func (c *client) SummonerChampionMastery(summoner.ID, lol.ChampionID) (*lol.ChampionMastery, error) {

}

func (c *client) SummonerChampionMasteryScore(summoner.ID) (int32, error) {

}

func (c *client) ActiveGame(summoner.ID) (*lol.CurrentGame, error) {

}
