package summoner

import "github.com/anthonyrouseau/go-riot/account"

//ID represents an encryptedSummonerId from Riot
type ID string

//Name is a league of legends summoner name
type Name string

//PUUID is the encrypted PUUID for a summoner
type PUUID string

//Info represents a SummonerDTO
type Info struct {
	ProfileIconID int32
	Name
	PUUID
	SummonerLevel int64
	RevisionDate  int64
	ID
	AccountID account.ID
}
