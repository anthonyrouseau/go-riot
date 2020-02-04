package tournament

import "time"

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
	Timestamp  time.Duration `json:"timestamp"`
	SummonerID summoner.ID   `json:"summonerId"`
	EventType  string        `json:"eventType"`
}

//Registration are the parameters passed in the body when creating a tournament
type Registration struct {
	Name       string
	ProviderID int32
}

//RegistrationOption is a function that allows you to change the default registration parameters
type RegistrationOption func(*Registration)

//RegistrationName is a RegistrationOption that updates the name parameter
func (r *Registration) RegistrationName(n string) func(*Registration) {
	return func(r *Registration) {
		r.Name = n
	}
}

//CodeRequest are the parameters passed in the body when requesting tournament codes
type CodeRequest struct {
	Count        int32
	TournamentID ID
}

//CodeRequestOption is a function that allows you to change the default code request parameters
type CodeRequestOption func(*CodeRequest)

//CodeRequestCount is a CodeRequestOption that updates the count parameter
func (c *CodeRequest) CodeRequestCount(n int32) func(*CodeRequest) {
	return func(c *CodeRequest) {
		c.Count = n
	}
}
