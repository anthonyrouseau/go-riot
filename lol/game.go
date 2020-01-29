package lol

import (
	"time"

	"github.com/anthonyrouseau/go-riot/summoner"
)

//CurrentGame represents a CurrentGameInfo
type CurrentGame struct {
	GameID            MatchID
	GameStartTime     time.Time
	PlatformID        string
	GameMode          string
	MapID             string
	GameType          string
	BannedChampions   []*BannedChampion
	Observers         *Observer
	Participants      []*CurrentGameParticipant
	GameLength        int64
	GameQueueConfigID int64
}

//CurrentGameParticipant represents a CurrentGameParticipantDTO
type CurrentGameParticipant struct {
	ProfileIconID            int64
	ChampionID               ChampionID
	SummonerName             summoner.Name
	GameCustomizationObjects []*GameCustomizationObject
	Bot                      bool
	Perks                    *Perks
	Spell1ID                 int64
	Spell2ID                 int64
	TeamID                   int64
	SummonerID               summoner.ID
}

//GameCustomizationObject represents a CurrentGameParticipant customization object
type GameCustomizationObject struct {
	Category string
	Content  string
}

//FeaturedGames represents FeaturedGames
type FeaturedGames struct {
	ClientRefreshInterval int64
	GameList              []*FeaturedGame
}

//FeaturedGame represents FeaturedGameInfo
type FeaturedGame struct {
	GameID            int64
	GameStartTime     time.Time
	PlatformID        string
	GameMode          string
	MapID             int64
	GameType          string
	BannedChampions   []*BannedChampion
	Observers         *Observer
	Participants      []*Participant
	GameLength        int64
	GameQueueConfigID int64
}
