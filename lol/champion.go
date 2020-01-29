package lol

import "github.com/anthonyrouseau/go-riot/summoner"

//ChampionID represents a champion Id
type ChampionID int64

//ChampionMastery represents a ChampionMasteryDTO containing mastery information for a single champion and summoner combination.
type ChampionMastery struct {
	chestGranted         bool
	level                int32
	points               int32
	championID           ChampionID
	pointsUntilNextLevel int64
	lastPlayTime         int64
	tokensEarned         int32
	pointsSinceLastLevel int64
	summonerID           summoner.ID
}

//Rotation represents a ChampionInfo
type Rotation struct {
	freeChampions           []ChampionID
	freeChampionsNewPlayers []ChampionID
	maxNewPlayerLevel       int32
}

//BannedChampion represents a BannedChampionDTO
type BannedChampion struct {
	pickTurn   int32
	championID ChampionID
	teamID     int64
}
