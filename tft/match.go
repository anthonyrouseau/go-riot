package tft

import "github.com/anthonyrouseau/go-riot/summoner"

//MatchID is the id for a match
type MatchID string

//Match represents a MatchDTO
type Match struct {
	Info     *MatchInfo     `json:"info"`
	Metadata *MatchMetaData `json:"metadata"`
}

//MatchInfo represents an InfoDTO
type MatchInfo struct {
	GameDatetime int64          `json:"game_datetime"`
	Participants []*Participant `json:"participants"`
	SetNumber    int32          `json:"tft_set_number"`
	GameLength   float64        `json:"game_length"`
	QueueID      int32          `json:"queue_id"`
	GameVersion  string         `json:"game_version"`
}

//MatchMetaData represents a MetaDataDTO
type MatchMetaData struct {
	DataVersion  string           `json:"data_version"`
	Participants []summoner.PUUID `json:"participants"`
	MatchID      MatchID          `json:"match_id"`
}
