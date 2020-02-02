package queue

//Name is the available queue names
type Name string

//The following are the allowed value for queue names
const (
	RankedSolo5x5 Name = "RANKED_SOLO_5x5"
	RankedFlexSR  Name = "RANKED_FLEX_SR"
	RankedFlexTT  Name = "RANKED_FLEX_TT"
)
