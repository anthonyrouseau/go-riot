package lol

import "time"

//MatchID is the id for a game
type MatchID int64

//Match represents a MatchDTO
type Match struct {
	seasonID              int32
	queueID               int32
	gameID                MatchID
	participantIdentities []*ParticipantIdentity
	gameVersion           string
	platformID            string
	gameMode              string
	mapID                 int32
	gameType              string
	teams                 []*TeamStats
	participants          []*Participant
	gameDuration          time.Duration
	gameCreation          time.Time
}

//MatchList represents a MatchListDTO
type MatchList struct {
	matches    []*MatchReference
	totalGames int32
	startIndex int32
	endIndex   int32
}

//MatchReference represents a MatchReferenceDTO
type MatchReference struct {
	lane       string
	gameID     MatchID
	champion   ChampionID
	platformID int32
	season     int32
	queue      int32
	role       string
	timestamp  time.Time
}

//MatchTimeline represents a MatchTimelineDTO
type MatchTimeline struct {
	frames        []*MatchFrame
	frameInterval int64
}

//MatchFrame represents a MatchFrameDTO
type MatchFrame struct {
	timestamp         time.Time
	participantFrames map[string]*MatchParticipantFrame
	events            []*MatchEvent
}

//MatchParticipantFrame represents a MatchParticipantFrameDTO
type MatchParticipantFrame struct {
	totalGold           int32
	teamScore           int32
	participantID       int32
	level               int32
	currentGold         int32
	minionsKilled       int32
	dominionScore       int32
	position            *MatchPosition
	xp                  int32
	jungleMinionsKilled int32
}

//MatchEvent represents a MatchEventDTO
type MatchEvent struct {
	eventType               string
	towerType               string
	teamID                  int32
	ascendedType            string
	killerID                int32
	levelUpType             string
	pointCaptured           string
	assistingParticipantIDs []int32
	wardType                string
	monsterType             string
	matchEventType          string
	skillSlot               int32
	victimID                int32
	timestamp               time.Time
	afterID                 int32
	monsterSubType          string
	laneType                string
	itemID                  int32
	participantID           int32
	buildingType            string
	creatorID               int32
	position                *MatchPosition
	beforeID                int32
}

//MatchPosition represents a MatchPositionDTO
type MatchPosition struct {
	y int32
	x int32
}
