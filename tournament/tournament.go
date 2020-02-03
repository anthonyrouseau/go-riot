package tournament

import "time"

import "github.com/anthonyrouseau/go-riot/summoner"

//Code is a tournament code
type Code string

//ID is the identifier for a tournament
type ID int32

//CodeInfo reperesents a TournamnetCodeDTO
type CodeInfo struct {
	CodeMap      string `json:"map"`
	Code         `json:"code"`
	Spectators   string   `json:"spectators"`
	Region       string   `json:"region"`
	ProviderID   int32    `json:"providerId"`
	TeamSize     int32    `json:"teamSize"`
	Participants []string `json:"participants"`
	PickType     string   `json:"pickType"`
	ID           `json:"id"`
	MetaData     string `json:"metaData"`
}

//LobbyEvents represents a LobbyEventDTOWrapper
type LobbyEvents struct {
	EventList []*LobbyEvent `json:"eventList"`
}

//LobbyEvent represents a LobbyEventDTO
type LobbyEvent struct {
	Timestamp  time.Duration `json:"timestamp"`
	SummonerID summoner.ID   `json:"summonerId"`
	EventType  string        `json:"eventType"`
}
