package tft

import (
	"time"

	"github.com/anthonyrouseau/go-riot/summoner"
)

//Participant represents a ParticipantDTO
type Participant struct {
	Placement            int32          `json:"placement"`
	Level                int32          `json:"level"`
	LastRound            int32          `json:"last_round"`
	TimeEliminated       time.Duration  `json:"time_eliminated"`
	Companion            *Companion     `json:"companion"`
	Traits               []*Trait       `json:"traits"`
	PlayersEliminated    int32          `json:"players_eliminated"`
	PUUID                summoner.PUUID `json:"puuid"`
	TotalDamageToPlayers int32          `json:"total_damage_to_players"`
	Units                []*Unit        `json:"units"`
	GoldLeft             int32          `json:"gold_left"`
}
