package tft

import (
	"time"

	"github.com/anthonyrouseau/go-riot/summoner"
)

//Participant represents a ParticipantDTO
type Participant struct {
	Placement      int32
	Level          int32
	LastRound      int32
	TimeEliminated time.Duration
	*Companion
	traits               []*Trait
	PlayersEliminated    int32
	PUUID                summoner.PUUID
	TotalDamageToPlayers int32
	Units                []*Unit
	GoldLeft             int32
}
