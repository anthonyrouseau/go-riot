package lol

import "time"

//MatchID is the id for a game
type MatchID int64

//Match represents a MatchDTO
type Match struct {
	SeasonID              int32                  `json:"seasonId"`
	QueueID               int32                  `json:"queueId"`
	GameID                MatchID                `json:"gameId"`
	ParticipantIdentities []*ParticipantIdentity `json:"participantIdentities"`
	GameVersion           string                 `json:"gameVersion"`
	PlatformID            string                 `json:"platformId"`
	GameMode              string                 `json:"gameMode"`
	MapID                 int32                  `json:"mapId"`
	GameType              string                 `json:"gameType"`
	Teams                 []*TeamStats           `json:"teams"`
	Participants          []*Participant         `json:"participants"`
	GameDuration          time.Duration          `json:"gameDuration"`
	GameCreation          time.Time              `json:"gameCreation"`
}

//MatchList represents a MatchListDTO
type MatchList struct {
	Matches    []*MatchReference `json:"matches"`
	TotalGames int32             `json:"totalGames"`
	StartIndex int32             `json:"startIndex"`
	EndIndex   int32             `json:"endIndex"`
}

//MatchReference represents a MatchReferenceDTO
type MatchReference struct {
	Lane       string     `json:"lane"`
	GameID     MatchID    `json:"gameId"`
	Champion   ChampionID `json:"champion"`
	PlatformID int32      `json:"platformId"`
	Season     int32      `json:"season"`
	Queue      int32      `json:"queue"`
	Role       string     `json:"role"`
	Timestamp  time.Time  `json:"timestamp"`
}

//MatchTimeline represents a MatchTimelineDTO
type MatchTimeline struct {
	Frames        []*MatchFrame `json:"frames"`
	FrameInterval int64         `json:"frameInterval"`
}

//MatchFrame represents a MatchFrameDTO
type MatchFrame struct {
	Timestamp         time.Time                         `json:"timestamp"`
	ParticipantFrames map[string]*MatchParticipantFrame `json:"participantFrames"`
	Events            []*MatchEvent                     `json:"events"`
}

//MatchParticipantFrame represents a MatchParticipantFrameDTO
type MatchParticipantFrame struct {
	TotalGold           int32          `json:"totalGold"`
	TeamScore           int32          `json:"teamScore"`
	ParticipantID       int32          `json:"participantId"`
	Level               int32          `json:"level"`
	CurrentGold         int32          `json:"currentGold"`
	MinionsKilled       int32          `json:"minionsKilled"`
	DominionScore       int32          `json:"dominionScore"`
	Position            *MatchPosition `json:"position"`
	XP                  int32          `json:"xp"`
	JungleMinionsKilled int32          `json:"jungleMinionsKilled"`
}

//MatchEvent represents a MatchEventDTO
type MatchEvent struct {
	EventType               string         `json:"eventType"`
	TowerType               string         `json:"towerType"`
	TeamID                  int32          `json:"teamId"`
	AscendedType            string         `json:"ascendedType"`
	KillerID                int32          `json:"killerId"`
	LevelUpType             string         `json:"levelUpType"`
	PointCaptured           string         `json:"pointCaptured"`
	AssistingParticipantIDs []int32        `json:"assistingParticpantIds"`
	WardType                string         `json:"wardType"`
	MonsterType             string         `json:"monsterType"`
	MatchEventType          string         `json:"type"`
	SkillSlot               int32          `json:"skillSlot"`
	VictimID                int32          `json:"victimId"`
	Timestamp               time.Time      `json:"timestamp"`
	AfterID                 int32          `json:"afterId"`
	MonsterSubType          string         `json:"monsterSubType"`
	LaneType                string         `json:"laneType"`
	ItemID                  int32          `json:"itemId"`
	ParticipantID           int32          `json:"participantId"`
	BuildingType            string         `json:"buildingType"`
	CreatorID               int32          `json:"creatorId"`
	Position                *MatchPosition `json:"position"`
	BeforeID                int32          `json:"beforeId"`
}

//MatchPosition represents a MatchPositionDTO
type MatchPosition struct {
	Y int32 `json:"y"`
	X int32 `json:"x"`
}
