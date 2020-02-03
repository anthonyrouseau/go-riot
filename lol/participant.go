package lol

import "time"

//ParticipantIdentity represetns a ParticipantIdentityDTO
type ParticipantIdentity struct {
	Player *Player `json:"player"`
	ID     int32   `json:"participantId"`
}

//Participant represents a ParticipantDTO
type Participant struct {
	Stats                     *ParticipantStats    `json:"stats"`
	ParticipantID             int32                `json:"participantId"`
	Runes                     []*Rune              `json:"runes"`
	Timeline                  *ParticipantTimeline `json:"timeline"`
	TeamID                    int32                `json:"teamId"`
	Spell2ID                  int32                `json:"spell2Id"`
	Masteries                 []*Mastery           `json:"masteries"`
	HighestAcheivedSeasonTier string               `json:"highestAchievedSeasonTier"`
	Spell1ID                  int32                `json:"spell1Id"`
	ChampionID                ChampionID           `json:"championId"`
}

//ParticipantStats represents a ParticipantStatsDTO
type ParticipantStats struct {
	FirstBloodAssist                bool          `json:"firstBloodAssist"`
	VisionScore                     int64         `json:"visionScore"`
	MagicDamageDealtToChampions     int64         `json:"magicDamageDealtToChampions"`
	DamageDealtToObjectives         int64         `json:"damageDealtToObjectives"`
	TotalTimeCrowdControlDealt      int32         `json:"totalTimeCrowdControlDealt"`
	LongestTimeSpentLiving          int32         `json:"longestTimeSpentLiving"`
	Perk0Var1                       int32         `json:"perk0Var1"`
	Perk0Var2                       int32         `json:"perk0Var2"`
	Perk0Var3                       int32         `json:"perk0Var3"`
	Perk1Var1                       int32         `json:"perk1Var1"`
	Perk1Var2                       int32         `json:"perk1Var2"`
	Perk1Var3                       int32         `json:"perk1Var3"`
	Perk2Var1                       int32         `json:"perk2Var1"`
	Perk2Var2                       int32         `json:"perk2Var2"`
	Perk2Var3                       int32         `json:"perk2Var3"`
	Perk3Var1                       int32         `json:"perk3Var1"`
	Perk3Var2                       int32         `json:"perk3Var2"`
	Perk3Var3                       int32         `json:"perk3Var3"`
	Perk4Var1                       int32         `json:"perk4Var1"`
	Perk4Var2                       int32         `json:"perk4Var2"`
	Perk4Var3                       int32         `json:"perk4Var3"`
	Perk5Var1                       int32         `json:"perk5Var1"`
	Perk5Var2                       int32         `json:"perk5Var2"`
	Perk5Var3                       int32         `json:"perk5Var3"`
	TripleKills                     int32         `json:"tripleKills"`
	NodeNeutralizeAssist            int32         `json:"nodeNeutralizeAssist"`
	PlayerScore0                    int32         `json:"playerScore0"`
	PlayerScore1                    int32         `json:"playerScore1"`
	PlayerScore2                    int32         `json:"playerScore2"`
	PlayerScore3                    int32         `json:"playerScore3"`
	PlayerScore4                    int32         `json:"playerScore4"`
	PlayerScore5                    int32         `json:"playerScore5"`
	PlayerScore6                    int32         `json:"playerScore6"`
	PlayerScore7                    int32         `json:"playerScore7"`
	PlayerScore8                    int32         `json:"playerScore8"`
	PlayerScore9                    int32         `json:"playerScore9"`
	Kills                           int32         `json:"kills"`
	TotalScoreRank                  int32         `json:"totalScoreRank"`
	NeutralMinionsKilled            int32         `json:"neutralMinionsKilled"`
	DamageDealtToTurrets            int64         `json:"damageDealtToTurrets"`
	PhysicalDamageDealtToChampions  int64         `json:"physicalDamgeDealtToChampions"`
	NodeCapture                     int32         `json:"nodeCapture"`
	LargestMultiKill                int32         `json:"largestMultiKill"`
	TotalUnitsHealed                int32         `json:"totalUnitsHealed"`
	WardsKilled                     int32         `json:"wardsKilled"`
	LargestCriticalStrike           int32         `json:"largestCriticalStrike"`
	LargestKillingSpree             int32         `json:"largestKIllingSpree"`
	QuadraKills                     int32         `json:"quadraKills"`
	TeamObjective                   int32         `json:"teamObjective"`
	MagicDamageDealt                int64         `json:"magicDamageDealt"`
	Item0                           int32         `json:"item0"`
	Item1                           int32         `json:"item1"`
	Item2                           int32         `json:"item2"`
	Item3                           int32         `json:"item3"`
	Item4                           int32         `json:"item4"`
	Item5                           int32         `json:"item5"`
	Item6                           int32         `json:"item6"`
	NeutralMinionsKilledTeamJungle  int32         `json:"neutralMinionsKilledTeamJungle"`
	Perk0                           int32         `json:"perk0"`
	Perk1                           int32         `json:"perk1"`
	Perk2                           int32         `json:"perk2"`
	Perk3                           int32         `json:"perk3"`
	Perk4                           int32         `json:"perk4"`
	Perk5                           int32         `json:"perk5"`
	DamageSelfMitigated             int64         `json:"damageSelfMitigated"`
	MagicalDamageTaken              int64         `json:"magicalDamgeTaken"`
	FirstInhibitorKill              bool          `json:"firstInhibitorKill"`
	TrueDamgeTaken                  int64         `json:"trueDamgeTaken"`
	NodeNeutralize                  int32         `json:"nodeNeutralize"`
	Assists                         int32         `json:"assists"`
	CombatPlayerScore               int32         `json:"combatPlayerScore"`
	PerkPrimaryStyle                int32         `json:"perkPrimaryStyle"`
	GoldSpent                       int32         `json:"goldSpent"`
	TrueDamageDealt                 int64         `json:"trueDamgeDealt"`
	ParticipantID                   int32         `json:"participantId"`
	TotalDamageTaken                int64         `json:"totalDamageTaken"`
	PhysicalDamageDealt             int64         `json:"physicalDamageDealt"`
	SightWardsBoughtInGame          int32         `json:"sightWardsBoughtInGame"`
	TotalDamageDealtToChampions     int64         `json:"totalDamgeDealtToChampions"`
	PhysicalDamageTaken             int64         `json:"physicalDamgeTaken"`
	TotalPlayerScore                int32         `json:"totalPlayerScore"`
	Win                             bool          `json:"win"`
	ObjectivePlayerScore            int32         `json:"objectivePlayerScore"`
	TotalDamageDealt                int64         `json:"totalDamgeDealt"`
	NeutralMinionsKilledEnemyJungle int32         `json:"neutralMinionsKilledEnemyJungle"`
	Deaths                          int32         `json:"deaths"`
	WardsPlaced                     int32         `json:"wardsPlaced"`
	PerkSubStyle                    int32         `json:"perkSubStyle"`
	TurretKills                     int32         `json:"turretKills"`
	FirstBloodKill                  bool          `json:"firstBloodKill"`
	TrueDamageDealtToChampions      int64         `json:"trueDamgeDealtToChampions"`
	GoldEarned                      int32         `json:"goldEarned"`
	KillingSprees                   int32         `json:"killingSprees"`
	UnrealKills                     int32         `json:"unrealKills"`
	AltarsCaptured                  int32         `json:"altarsCaptured"`
	FirstTowerAssist                bool          `json:"firstTowerAssist"`
	FirstTowerKill                  bool          `json:"firstTowerKill"`
	ChampLevel                      int32         `json:"champLevel"`
	DoubleKills                     int32         `json:"doubleKills"`
	NodeCaptureAssist               int32         `json:"nodeCaptureAssist"`
	InhibitorKills                  int32         `json:"inhibitorKills"`
	FirstInhibitorAssist            bool          `json:"firstInhibitorAssist"`
	VisionWardsBoughtInGame         int32         `json:"visionWardsBoughtInGame"`
	AltarsNeutralized               int32         `json:"altarsNeutralized"`
	PentaKills                      int32         `json:"pentaKills"`
	TotalHeal                       int32         `json:"totalHeal"`
	TotalMinionsKilled              int32         `json:"totalMinionsKilled"`
	TimeCCingOthers                 time.Duration `json:"timeCCingOthers"`
}

//ParticipantTimeline represents a ParticipantTimelineDTO
type ParticipantTimeline struct {
	Lane                        string             `json:"lane"`
	ParticipantID               int32              `json:"participantId"`
	CSDiffPerMinDeltas          map[string]float32 `json:"csDiffPerMinDeltas"`
	GoldPerMinDeltas            map[string]float32 `json:"goldPerMinDeltas"`
	XPDiffPerMinDeltas          map[string]float32 `json:"xpDiffPerMinDeltas"`
	CreepsPerMinDeltas          map[string]float32 `json:"creepsPerminDeltas"`
	XPPerMinDeltas              map[string]float32 `json:"xpPerMinDeltas"`
	Role                        string             `json:"role"`
	DamageTakenDiffPerMinDeltas map[string]float32 `json:"damageTakenDiffPerMinDeltas"`
	DamageTakePerMinDeltas      map[string]float32 `json:"damageTakenPerMinDeltas"`
}
