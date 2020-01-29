package lol

//TeamStats represents a TeamStatsDTO
type TeamStats struct {
	FirstDragon          bool
	FirstInhibitor       bool
	Bans                 []*TeamBans
	BaronKills           int32
	FirstRiftHerald      bool
	FirstBaron           bool
	RiftHeraldKills      int32
	FirstBlood           bool
	TeamID               int32
	FirstTower           bool
	VilemawKills         int32
	InhibitorKills       int32
	TowerKills           int32
	DominionVictoryScore int32
	Win                  string
	DragonKills          int32
}

//TeamBans represents a TeamBansDTO
type TeamBans struct {
	PickTurn   int32
	ChampionID ChampionID
}
