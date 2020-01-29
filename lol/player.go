package lol

import "github.com/anthonyrouseau/go-riot/summoner"

//Player represents a PlayerDTO
type Player struct {
	currentPlatformID string
	summonerName      summoner.Name
	matchHistoryURI   string
	platformID        string
	currentAccountID  string
	profileIcon       int32
	summonerID        summoner.ID
	accountID         string
}
