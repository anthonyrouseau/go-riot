package shard

//Status represents a ShardStatus
type Status struct {
	name      string
	regionTag string
	hostname  string
	services  []*Service
	slug      string
	locales   []string
}

//Service gives information about a service and its incidents
type Service struct {
	status    string
	incidents []*Incident
	name      string
	slug      string
}

//Incident gives information about a services incident and updates
type Incident struct {
	active    bool
	createdAt string
	ID        int64
	updates   []*Message
}

//Message gives information about an incident
type Message struct {
	severity     string
	author       string
	createdAt    string
	translations []*Translation
	updatedAt    string
	content      string
	ID           string
}

//Translation is a translation for a message
type Translation struct {
	locale    string
	content   string
	updatedAt string
}
