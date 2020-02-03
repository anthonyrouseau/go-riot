package lol

import "github.com/anthonyrouseau/go-riot/summoner"

//ChampionID represents a champion Id
type ChampionID int64

//ChampionMastery represents a ChampionMasteryDTO containing mastery information for a single champion and summoner combination.
type ChampionMastery struct {
	ChestGranted         bool        `json:"chestGranted"`
	Level                int32       `json:"championLevel"`
	Points               int32       `json:"championPoints"`
	ChampionID           ChampionID  `json:"championId"`
	PointsUntilNextLevel int64       `json:"championPointsUntilNextLevel"`
	LastPlayTime         int64       `json:"lastPlayTime"`
	TokensEarned         int32       `json:"tokensEarned"`
	PointsSinceLastLevel int64       `json:"championPointsSinceLastLevel"`
	SummonerID           summoner.ID `json:"summonerId"`
}

//Rotation represents a ChampionInfo
type Rotation struct {
	FreeChampions           []ChampionID `json:"freeChampionIds"`
	FreeChampionsNewPlayers []ChampionID `json:"freeChampionIdsForNewPlayers"`
	MaxNewPlayerLevel       int32        `json:"maxNewPlayerLevel"`
}

//BannedChampion represents a BannedChampionDTO
type BannedChampion struct {
	PickTurn   int32      `json:"pickTurn"`
	ChampionID ChampionID `json:"championId"`
	TeamID     int64      `json:"teamId"`
}
