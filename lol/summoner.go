package lol

import "github.com/anthonyrouseau/go-riot/account"

//SummonerID represents an encryptedSummonerId from Riot
type SummonerID string

//SummonerName is a league of legends summoner name
type SummonerName string

//PUUID is the encrypted PUUID for a summoner
type PUUID string

//Summoner represents a SummonerDTO
type Summoner struct {
	profileIconID int32
	name          SummonerName
	puuid         string
	summonerLevel int64
	revisionDate  int64
	ID            SummonerID
	accountID     account.ID
}
