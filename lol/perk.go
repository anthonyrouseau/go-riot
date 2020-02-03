package lol

//Perks is the perks/runes reforged info
type Perks struct {
	PerkStyle    int64   `json:"perkStyle"`
	PerkIDs      []int64 `json:"perkIds"`
	PerkSubstyle int64   `json:"perkSubstyle"`
}
