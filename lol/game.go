package lol

import (
	"time"

	"github.com/anthonyrouseau/go-riot/summoner"
)

//CurrentGame represents a CurrentGameInfo
type CurrentGame struct {
	gameID            int64
	gameStartTime     time.Time
	platformID        string
	gameMode          string
	mapID             string
	gameType          string
	bannedChampions   []*BannedChampion
	observers         *Observer
	participants      []*CurrentGameParticipant
	gameLength        int64
	gameQueueConfigID int64
}

//CurrentGameParticipant represents a CurrentGameParticipantDTO
type CurrentGameParticipant struct {
	profileIconID            int64
	championID               ChampionID
	summonerName             summoner.Name
	gameCustomizationObjects []*GameCustomizationObject
	bot                      bool
	perks                    *Perks
	spell1ID                 int64
	spell2ID                 int64
	teamID                   int64
	summonerID               summoner.ID
}

//GameCustomizationObject represents a CurrentGameParticipant customization object
type GameCustomizationObject struct {
	category string
	content  string
}

//FeaturedGames represents FeaturedGames
type FeaturedGames struct {
	clientRefreshInterval int64
	gameList              []*FeaturedGame
}

//FeaturedGame represents FeaturedGameInfo
type FeaturedGame struct {
	gameID            int64
	gameStartTime     time.Time
	platformID        string
	gameMode          string
	mapID             int64
	gameType          string
	bannedChampions   []*BannedChampion
	observers         *Observer
	participants      []*Participant
	gameLength        int64
	gameQueueConfigID int64
}
