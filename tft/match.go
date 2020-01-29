package tft

import "time"

import "github.com/anthonyrouseau/go-riot/summoner"

//MatchID is the id for a match
type MatchID string

//Match represents a MatchDTO
type Match struct {
	Info     *MatchInfo
	Metadata *MatchMetaData
}

//MatchInfo represents an InfoDTO
type MatchInfo struct {
	GameDatetime time.Time
	Participants []*Participant
	SetNumber    int32
	GameLength   time.Duration
	QueueID      int32
	GameVersion  string
}

//MatchMetaData represents a MetaDataDTO
type MatchMetaData struct {
	DataVersion  string
	Participants []summoner.PUUID
	MatchID      MatchID
}
