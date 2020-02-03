package lol

import "github.com/anthonyrouseau/go-riot/summoner"

import "github.com/anthonyrouseau/go-riot/account"

//Player represents a PlayerDTO
type Player struct {
	CurrentPlatformID string        `json:"currentPlatformId"`
	SummonerName      summoner.Name `json:"summonerName"`
	MatchHistoryURI   string        `json:"matchHistoryUri"`
	PlatformID        string        `json:"platformId"`
	CurrentAccountID  string        `json:"currentAccountId"`
	ProfileIcon       int32         `json:"profileIcon"`
	SummonerID        summoner.ID   `json:"summonerId"`
	AccountID         account.ID    `json:"accountId"`
}
