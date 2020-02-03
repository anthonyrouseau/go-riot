package lol

//TeamStats represents a TeamStatsDTO
type TeamStats struct {
	FirstDragon          bool        `json:"firstDragon"`
	FirstInhibitor       bool        `json:"firstInhibitor"`
	Bans                 []*TeamBans `json:"bans"`
	BaronKills           int32       `json:"baronKills"`
	FirstRiftHerald      bool        `json:"firstRiftHerald"`
	FirstBaron           bool        `json:"firstBaron"`
	RiftHeraldKills      int32       `json:"riftHeraldKills"`
	FirstBlood           bool        `json:"firstBLood"`
	TeamID               int32       `json:"teamId"`
	FirstTower           bool        `json:"firstTower"`
	VilemawKills         int32       `json:"vilemawKills"`
	InhibitorKills       int32       `json:"inhibitorKills"`
	TowerKills           int32       `json:"towerKills"`
	DominionVictoryScore int32       `json:"dominionVicotryScore"`
	Win                  string      `json:"win"`
	DragonKills          int32       `json:"dragonKills"`
}

//TeamBans represents a TeamBansDTO
type TeamBans struct {
	PickTurn   int32      `json:"pickTurn"`
	ChampionID ChampionID `json:"championId"`
}
