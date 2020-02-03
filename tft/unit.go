package tft

//Unit represents a UnitDTO
type Unit struct {
	Tier        int32   `json:"tier"`
	Items       []int32 `json:"items"`
	CharacterID string  `json:"character_id"`
	Name        string  `json:"name"`
	Rarity      int32   `json:"rarity"`
}
