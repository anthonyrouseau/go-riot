package lol

import "time"

//MatchID is the id for a game
type MatchID int64

//Match represents a MatchDTO
type Match struct {
	SeasonID              int32
	QueueID               int32
	GameID                MatchID
	ParticipantIdentities []*ParticipantIdentity
	GameVersion           string
	PlatformID            string
	GameMode              string
	MapID                 int32
	GameType              string
	Teams                 []*TeamStats
	Participants          []*Participant
	GameDuration          time.Duration
	GameCreation          time.Time
}

//MatchList represents a MatchListDTO
type MatchList struct {
	Matches    []*MatchReference
	TotalGames int32
	StartIndex int32
	EndIndex   int32
}

//MatchReference represents a MatchReferenceDTO
type MatchReference struct {
	Lane       string
	GameID     MatchID
	Champion   ChampionID
	PlatformID int32
	Season     int32
	Queue      int32
	Role       string
	Timestamp  time.Time
}

//MatchTimeline represents a MatchTimelineDTO
type MatchTimeline struct {
	Frames        []*MatchFrame
	FrameInterval int64
}

//MatchFrame represents a MatchFrameDTO
type MatchFrame struct {
	Timestamp         time.Time
	ParticipantFrames map[string]*MatchParticipantFrame
	Events            []*MatchEvent
}

//MatchParticipantFrame represents a MatchParticipantFrameDTO
type MatchParticipantFrame struct {
	TotalGold           int32
	TeamScore           int32
	ParticipantID       int32
	Level               int32
	CurrentGold         int32
	MinionsKilled       int32
	DominionScore       int32
	Position            *MatchPosition
	XP                  int32
	JungleMinionsKilled int32
}

//MatchEvent represents a MatchEventDTO
type MatchEvent struct {
	EventType               string
	TowerType               string
	TeamID                  int32
	AscendedType            string
	KillerID                int32
	LevelUpType             string
	PointCaptured           string
	AssistingParticipantIDs []int32
	WardType                string
	MonsterType             string
	MatchEventType          string
	SkillSlot               int32
	VictimID                int32
	Timestamp               time.Time
	AfterID                 int32
	MonsterSubType          string
	LaneType                string
	ItemID                  int32
	ParticipantID           int32
	BuildingType            string
	CreatorID               int32
	Position                *MatchPosition
	BeforeID                int32
}

//MatchPosition represents a MatchPositionDTO
type MatchPosition struct {
	Y int32
	X int32
}
