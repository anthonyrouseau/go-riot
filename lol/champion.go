package lol

import "github.com/anthonyrouseau/go-riot/summoner"

//ChampionID represents a champion Id
type ChampionID int64

//ChampionMastery represents a ChampionMasteryDTO containing mastery information for a single champion and summoner combination.
type ChampionMastery struct {
	ChestGranted         bool
	Level                int32
	Points               int32
	ChampionID           ChampionID
	PointsUntilNextLevel int64
	LastPlayTime         int64
	TokensEarned         int32
	PointsSinceLastLevel int64
	SummonerID           summoner.ID
}

//Rotation represents a ChampionInfo
type Rotation struct {
	FreeChampions           []ChampionID
	FreeChampionsNewPlayers []ChampionID
	MaxNewPlayerLevel       int32
}

//BannedChampion represents a BannedChampionDTO
type BannedChampion struct {
	PickTurn   int32
	ChampionID ChampionID
	TeamID     int64
}
