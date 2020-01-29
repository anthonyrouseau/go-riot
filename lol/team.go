package lol

//TeamStats represents a TeamStatsDTO
type TeamStats struct {
	firstDragon          bool
	firstInhibitor       bool
	bans                 []*TeamBans
	baronKills           int32
	firstRiftHerald      bool
	firstBaron           bool
	riftHeraldKills      int32
	firstBlood           bool
	teamID               int32
	firstTower           bool
	vilemawKills         int32
	inhibitorKills       int32
	towerKills           int32
	dominionVictoryScore int32
	win                  string
	dragonKills          int32
}

//TeamBans represents a TeamBansDTO
type TeamBans struct {
	pickTurn   int32
	championID ChampionID
}
