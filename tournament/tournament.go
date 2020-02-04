package tournament

import "github.com/anthonyrouseau/go-riot/summoner"

//Code is a tournament code
type Code string

//ID is the identifier for a tournament
type ID int32

//Region is the region in which the provider will be running tournaments
type Region string

//Possible values for tournament provider region
const (
	BR   Region = "BR"
	EUNE Region = "EUNE"
	EUW  Region = "EUW"
	JP   Region = "JP"
	LAN  Region = "LAN"
	LAS  Region = "LAS"
	NA   Region = "NA"
	OCE  Region = "OCE"
	PBE  Region = "PBE"
	RU   Region = "RU"
	TR   Region = "TR"
)

//CodeInfo reperesents a TournamnetCodeDTO
type CodeInfo struct {
	CodeMap      string `json:"map"`
	Code         `json:"code"`
	Spectators   string   `json:"spectators"`
	Region       string   `json:"region"`
	ProviderID   int32    `json:"providerId"`
	TeamSize     int32    `json:"teamSize"`
	Participants []string `json:"participants"`
	PickType     string   `json:"pickType"`
	ID           `json:"id"`
	MetaData     string `json:"metaData"`
}

//LobbyEvents represents a LobbyEventDTOWrapper
type LobbyEvents struct {
	EventList []*LobbyEvent `json:"eventList"`
}

//LobbyEvent represents a LobbyEventDTO
type LobbyEvent struct {
	Timestamp  string      `json:"timestamp"`
	SummonerID summoner.ID `json:"summonerId"`
	EventType  string      `json:"eventType"`
}

//Registration are the parameters passed in the body when creating a tournament
type Registration struct {
	Name       *string `json:"name"`
	ProviderID int32   `json:"providerId"`
}

//RegistrationOption is a function that allows you to change the default registration parameters
type RegistrationOption func(*Registration)

//RegistrationName is a RegistrationOption that updates the name parameter
func (r *Registration) RegistrationName(n string) func(*Registration) {
	return func(r *Registration) {
		r.Name = &n
	}
}

//CodeRequest are the optional parameters for requesting tournament codes
type CodeRequest struct {
	Count int32 `json:"count"`
}

//CodeRequestBody are the required body parameters for requesting tournament codes
type CodeRequestBody struct {
	AllowedSummonerIDs []summoner.ID `json:"allowedSummonerIds,omitempty"`
	MapType            MapType       `json:"mapType"`
	Metadata           string        `json:"metadata"`
	PickType           PickType      `json:"pickType"`
	SpectatorType      SpectatorType `json:"spectatorType"`
	TeamSize           int32         `json:"teamSize"`
}

//CodeUpdateRequest is the body for a put request to tournament code
type CodeUpdateRequest struct {
	AllowedSummonerIDs []summoner.ID  `json:"allowedSummonerIds,omitempty"`
	MapType            *MapType       `json:"mapType,omitempty"`
	PickType           *PickType      `json:"pickType,omitempty"`
	SpectatorType      *SpectatorType `json:"spectatorType"`
}

//MapType is the type for a code request map type
type MapType string

//Possible values for MapType
const (
	SummonersRift   MapType = "SUMMONERS_RIFT"
	TwistedTreeline MapType = "TWISTED_TREELINE"
	HowlingAbyss    MapType = "HOWLING_ABYSS"
)

//PickType is the type for a code request pick type
type PickType string

//Possible values for PickType
const (
	BlindPick       PickType = "BLIND_PICK"
	DraftMode       PickType = "DRAFT_MODE"
	AllRandom       PickType = "ALL_RANDOM"
	TournamentDraft PickType = "TOURNAMENT_DRAFT"
)

//SpectatorType is the type for a code request spectator type
type SpectatorType string

//Possible values for SpectatorType
const (
	None      SpectatorType = "NONE"
	LobbyOnly SpectatorType = "LOBBYONLY"
	All       SpectatorType = "ALL"
)

//CodeRequestOption is a function that allows you to change the default code request parameters
type CodeRequestOption func(*CodeRequest)

//CodeRequestCount is a CodeRequestOption that updates the count parameter
func (c *CodeRequest) CodeRequestCount(n int32) func(*CodeRequest) {
	return func(c *CodeRequest) {
		c.Count = n
	}
}

//ProviderRegistration is the post body when creating a tournament provider
type ProviderRegistration struct {
	URL    string `json:"url"`
	Region Region `json:"region"`
}
