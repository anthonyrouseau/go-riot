package lol

import "time"

//ParticipantIdentity represetns a ParticipantIdentityDTO
type ParticipantIdentity struct {
	Player *Player
	ID     int32
}

//Participant represents a ParticipantDTO
type Participant struct {
	Stats                     *ParticipantStats
	ParticipantID             int32
	Runes                     []*Rune
	Timeline                  *ParticipantTimeline
	TeamID                    int32
	Spell2ID                  int32
	Masteries                 []*Mastery
	HighestAcheivedSeasonTier string
	Spell1ID                  int32
	ChampionID                ChampionID
}

//ParticipantStats represents a ParticipantStatsDTO
type ParticipantStats struct {
	FirstBloodAssist                bool
	VisionScore                     int64
	MagicDamageDealtToChampions     int64
	DamageDealtToObjectives         int64
	TotalTimeCrowdControlDealt      int32
	LongestTimeSpentLiving          int32
	Perk0Var1                       int32
	Perk0Var2                       int32
	Perk0Var3                       int32
	Perk1Var1, Perk1Var2, Perk1Var3 int32
	Perk2Var1, Perk2Var2, Perk2Var3 int32
	Perk3Var1, Perk3Var2, Perk3Var3 int32
	Perk4Var1, Perk4Var2, Perk4Var3 int32
	Perk5Var1, Perk5Var2, Perk5Var3 int32
	TripleKills                     int32
	NodeNeutralizeAssist            int32
	PlayerScore0, PlayerScore1, PlayerScore2, PlayerScore3, PlayerScore4,
	PlayerScore5, PlayerScore6, PlayerScore7, PlayerScore8, PlayerScore9 int32
	Kills                                           int32
	TotalScoreRank                                  int32
	NeutralMinionsKilled                            int32
	DamageDealtToTurrets                            int64
	PhysicalDamageDealtToChampions                  int64
	NodeCapture                                     int32
	LargetMultiKill                                 int32
	TotalUnitsHealed                                int32
	WardsKilled                                     int32
	LargestCriticalStrike                           int32
	LargestKillingSpree                             int32
	QuadraKills                                     int32
	TeamObjective                                   int32
	MagicDamageDealt                                int64
	Item0, Item1, Item2, Item3, Item4, Item5, Item6 int32
	NeutralMinionsKilledTeamJungle                  int32
	Perk0, Perk1, Perk2, Perk3, Perk4, Perk5        int32
	DamageSelfMitigated                             int64
	MagicalDamageTaken                              int64
	FirstInhibitorKill                              bool
	TrueDamgeTaken                                  int64
	NodeNeutralize                                  int32
	Assists                                         int32
	CombatPlayerScore                               int32
	PerkPrimaryStyle                                int32
	GoldSpent                                       int32
	TrueDamageDealt                                 int64
	ParticipantID                                   int32
	TotalDamageTake                                 int64
	PhysicalDamageDealt                             int64
	SightWardsBoughtInGame                          int32
	TotalDamageDealtToChampions                     int64
	PhysicalDamageTaken                             int64
	TotalPlayerScore                                int32
	Win                                             bool
	ObjectivePlayerScore                            int32
	TotalDamageDealt                                int64
	NeutralMinionsKilledEnemyJungle                 int32
	Deaths                                          int32
	WardsPlaced                                     int32
	PerkSubStyle                                    int32
	TurretKills                                     int32
	FirstBloodKill                                  bool
	TrueDamageDealtToChampions                      int64
	GoldEarned                                      int32
	KillingSprees                                   int32
	UnrealKills                                     int32
	AltarsCaptured                                  int32
	FirstTowerAssist                                bool
	FirstTowerKill                                  bool
	ChampLevel                                      int32
	DoubleKills                                     int32
	NodeCaptureAssist                               int32
	InhibitorKills                                  int32
	FirstInhibitorAssist                            bool
	VisionWardsBoughtInGame                         int32
	AltarsNeutralized                               int32
	PentaKills                                      int32
	TotalHeal                                       int32
	TotalMinionsKilled                              int32
	TimeCCingOthers                                 time.Duration
}

//ParticipantTimeline represents a ParticipantTimelineDTO
type ParticipantTimeline struct {
	Lane                        string
	ParticipantID               int32
	CSDiffPerMinDeltas          map[string]float32
	GoldPerMinDeltas            map[string]float32
	XPDiffPerMinDeltas          map[string]float32
	CreepsPerMinDeltas          map[string]float32
	XPPerMinDeltas              map[string]float32
	Role                        string
	DamageTakenDiffPerMinDeltas map[string]float32
	DamageTakePerMinDeltas      map[string]float32
}
