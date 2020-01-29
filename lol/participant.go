package lol

import "time"

//ParticipantIdentity represetns a ParticipantIdentityDTO
type ParticipantIdentity struct {
	player *Player
	id     int32
}

//Participant represents a ParticipantDTO
type Participant struct {
	stats                     *ParticipantStats
	participantID             int32
	runes                     []*Rune
	timeline                  *ParticipantTimeline
	teamID                    int32
	spell2ID                  int32
	masteries                 []*Mastery
	highestAcheivedSeasonTier string
	spell1ID                  int32
	championID                ChampionID
}

//ParticipantStats represents a ParticipantStatsDTO
type ParticipantStats struct {
	firstBloodAssist                bool
	visionScore                     int64
	magicDamageDealtToChampions     int64
	damageDealtToObjectives         int64
	totalTimeCrowdControlDealt      int32
	longestTimeSpentLiving          int32
	perk0Var1                       int32
	perk0Var2                       int32
	perk0Var3                       int32
	perk1Var1, perk1Var2, perk1Var3 int32
	perk2Var1, perk2Var2, perk2Var3 int32
	perk3Var1, perk3Var2, perk3Var3 int32
	perk4Var1, perk4Var2, perk4Var3 int32
	perk5Var1, perk5Var2, perk5Var3 int32
	tripleKills                     int32
	nodeNeutralizeAssist            int32
	playerScore0, playerScore1, playerScore2, playerScore3, playerScore4,
	playerScore5, playerScore6, playerScore7, playerScore8, playerScore9 int32
	kills                                           int32
	totalScoreRank                                  int32
	neutralMinionsKilled                            int32
	damageDealtToTurrets                            int64
	physicalDamageDealtToChampions                  int64
	nodeCapture                                     int32
	largetMultiKill                                 int32
	totalUnitsHealed                                int32
	wardsKilled                                     int32
	largestCriticalStrike                           int32
	largestKillingSpree                             int32
	quadraKills                                     int32
	teamObjective                                   int32
	magicDamageDealt                                int64
	item0, item1, item2, item3, item4, item5, item6 int32
	neutralMinionsKilledTeamJungle                  int32
	perk0, perk1, perk2, perk3, perk4, perk5        int32
	damageSelfMitigated                             int64
	magicalDamageTaken                              int64
	firstInhibitorKill                              bool
	trueDamgeTaken                                  int64
	nodeNeutralize                                  int32
	assists                                         int32
	combatPlayerScore                               int32
	perkPrimaryStyle                                int32
	goldSpent                                       int32
	trueDamageDealt                                 int64
	participantID                                   int32
	totalDamageTake                                 int64
	physicalDamageDealt                             int64
	sightWardsBoughtInGame                          int32
	totalDamageDealtToChampions                     int64
	physicalDamageTaken                             int64
	totalPlayerScore                                int32
	win                                             bool
	objectivePlayerScore                            int32
	totalDamageDealt                                int64
	neutralMinionsKilledEnemyJungle                 int32
	deaths                                          int32
	wardsPlaced                                     int32
	perkSubStyle                                    int32
	turretKills                                     int32
	firstBloodKill                                  bool
	trueDamageDealtToChampions                      int64
	goldEarned                                      int32
	killingSprees                                   int32
	unrealKills                                     int32
	altarsCaptured                                  int32
	firstTowerAssist                                bool
	firstTowerKill                                  bool
	champLevel                                      int32
	doubleKills                                     int32
	nodeCaptureAssist                               int32
	inhibitorKills                                  int32
	firstInhibitorAssist                            bool
	visionWardsBoughtInGame                         int32
	altarsNeutralized                               int32
	pentaKills                                      int32
	totalHeal                                       int32
	totalMinionsKilled                              int32
	timeCCingOthers                                 time.Duration
}

//ParticipantTimeline represents a ParticipantTimelineDTO
type ParticipantTimeline struct {
	lane                        string
	participantID               int32
	csDiffPerMinDeltas          map[string]float32
	goldPerMinDeltas            map[string]float32
	xpDiffPerMinDeltas          map[string]float32
	creepsPerMinDeltas          map[string]float32
	xpPerMinDeltas              map[string]float32
	role                        string
	damageTakenDiffPerMinDeltas map[string]float32
	damageTakePerMinDeltas      map[string]float32
}
