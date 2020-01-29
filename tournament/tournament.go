package tournament

import "time"

import "github.com/anthonyrouseau/go-riot/summoner"

//Code is a tournament code
type Code string

//ID is the identifier for a tournament
type ID int32

//CodeInfo reperesents a TournamnetCodeDTO
type CodeInfo struct {
	CodeMap string
	Code
	Spectators   string
	Region       string
	ProviderID   int32
	TeamSize     int32
	Participants []string
	PickType     string
	ID
	MetaData string
}

//LobbyEvents represents a LobbyEventDTOWrapper
type LobbyEvents struct {
	EventList []*LobbyEvent
}

//LobbyEvent represents a LobbyEventDTO
type LobbyEvent struct {
	Timestamp  time.Duration
	SummonerID summoner.ID
	EventType  string
}
