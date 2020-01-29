package lol

import "github.com/anthonyrouseau/go-riot/summoner"

import "github.com/anthonyrouseau/go-riot/account"

//Player represents a PlayerDTO
type Player struct {
	CurrentPlatformID string
	SummonerName      summoner.Name
	MatchHistoryURI   string
	PlatformID        string
	CurrentAccountID  string
	ProfileIcon       int32
	SummonerID        summoner.ID
	AccountID         account.ID
}
