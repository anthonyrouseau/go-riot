package lol

import (
	"github.com/anthonyrouseau/go-riot/summoner"
)

//CurrentGame represents a CurrentGameInfo
type CurrentGame struct {
	GameID            MatchID                   `json:"gameId"`
	GameStartTime     int64                     `json:"gameStartTime"`
	PlatformID        string                    `json:"platformId"`
	GameMode          string                    `json:"gameMode"`
	MapID             string                    `json:"mapId"`
	GameType          string                    `json:"gameType"`
	BannedChampions   []*BannedChampion         `json:"bannedChampions"`
	Observers         *Observer                 `json:"observers"`
	Participants      []*CurrentGameParticipant `json:"participants"`
	GameLength        int64                     `json:"gameLength"`
	GameQueueConfigID int64                     `json:"gameQueueConfigId"`
}

//CurrentGameParticipant represents a CurrentGameParticipantDTO
type CurrentGameParticipant struct {
	ProfileIconID            int64                      `json:"profileIconId"`
	ChampionID               ChampionID                 `json:"championId"`
	SummonerName             summoner.Name              `json:"summonerName"`
	GameCustomizationObjects []*GameCustomizationObject `json:"gameCustomizationObjects"`
	Bot                      bool                       `json:"bot"`
	Perks                    *Perks                     `json:"perks"`
	Spell1ID                 int64                      `json:"spell1Id"`
	Spell2ID                 int64                      `json:"spell2Id"`
	TeamID                   int64                      `json:"teamId"`
	SummonerID               summoner.ID                `json:"summonerId"`
}

//GameCustomizationObject represents a CurrentGameParticipant customization object
type GameCustomizationObject struct {
	Category string `json:"category"`
	Content  string `json:"content"`
}

//FeaturedGames represents FeaturedGames
type FeaturedGames struct {
	ClientRefreshInterval int64           `json:"clientRefreshInterval"`
	GameList              []*FeaturedGame `json:"gameList"`
}

//FeaturedGame represents FeaturedGameInfo
type FeaturedGame struct {
	GameID            int64             `json:"gameId"`
	GameStartTime     int64             `json:"gameStartTime"`
	PlatformID        string            `json:"platformId"`
	GameMode          string            `json:"gameMode"`
	MapID             int64             `json:"mapId"`
	GameType          string            `json:"gameType"`
	BannedChampions   []*BannedChampion `json:"bannedChampions"`
	Observers         *Observer         `json:"observers"`
	Participants      []*Participant    `json:"participants"`
	GameLength        int64             `json:"gameLength"`
	GameQueueConfigID int64             `json:"gameQueueConfigId"`
}
