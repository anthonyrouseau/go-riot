package lol

import "time"

import "net/url"

import "strings"

import "strconv"

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
	GameCreation          int64                  `json:"gameCreation"`
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
	PlatformID string     `json:"platformId"`
	Season     int32      `json:"season"`
	Queue      int32      `json:"queue"`
	Role       string     `json:"role"`
	Timestamp  int64      `json:"timestamp"`
}

//MatchTimeline represents a MatchTimelineDTO
type MatchTimeline struct {
	Frames        []*MatchFrame `json:"frames"`
	FrameInterval int64         `json:"frameInterval"`
}

//MatchFrame represents a MatchFrameDTO
type MatchFrame struct {
	Timestamp         int64                             `json:"timestamp"`
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
	Timestamp               int64          `json:"timestamp"`
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

//MatchOption is a function that updates the query parameters for a matchlist query
type MatchOption func(*url.Values)

//SetChampion returns a MatchOption to set the Champion parameter
func SetChampion(champs []int) MatchOption {
	var stringChamps []string
	for _, v := range champs {
		stringChamps = append(stringChamps, strconv.Itoa(v))
	}
	return func(v *url.Values) {
		v.Set("champion", strings.Join(stringChamps, ","))
	}
}

//SetQueue returns a MatchOption to set the Queue parameter
func SetQueue(queue []int) MatchOption {
	var queueString []string
	for _, v := range queue {
		queueString = append(queueString, strconv.Itoa(v))
	}
	return func(v *url.Values) {
		v.Set("queue", strings.Join(queueString, ","))
	}
}

//SetSeason returns a MatchOption to set the Season parameter
func SetSeason(season []int) MatchOption {
	var seasonString []string
	for _, v := range season {
		seasonString = append(seasonString, strconv.Itoa(v))
	}
	return func(v *url.Values) {
		v.Set("season", strings.Join(seasonString, ","))
	}
}

//SetEndTime returns a MatchOption to set the EndTime parameter
func SetEndTime(end int64) MatchOption {
	return func(v *url.Values) {
		v.Set("endTime", strconv.FormatInt(end, 10))
	}
}

//SetBeginTime returns a MatchOption to set the BeginTime parameter
func SetBeginTime(begin int64) MatchOption {
	return func(v *url.Values) {
		v.Set("beginTime", strconv.FormatInt(begin, 10))
	}
}

//SetEndIndex returns a MatchOption to set the EndIndex parameter
func SetEndIndex(end int64) MatchOption {
	return func(v *url.Values) {
		v.Set("endIndex", strconv.FormatInt(end, 10))
	}
}

//SetBeginIndex returns a MatchOption to set the BeginIndex parameter
func SetBeginIndex(begin int64) MatchOption {
	return func(v *url.Values) {
		v.Set("beginIndex", strconv.FormatInt(begin, 10))
	}
}
