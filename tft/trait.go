package tft

//Trait represents a TraitDTO
type Trait struct {
	TierTotal   int32  `json:"tier_total"`
	Name        string `json:"name"`
	TierCurrent int32  `json:"tier_current"`
	NumUnits    int32  `json:"num_units"`
}
