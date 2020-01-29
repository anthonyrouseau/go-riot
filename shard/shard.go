package shard

//Status represents a ShardStatus
type Status struct {
	Name      string
	RegionTag string
	Hostname  string
	Services  []*Service
	Slug      string
	Locales   []string
}

//Service gives information about a service and its incidents
type Service struct {
	Status    string
	Incidents []*Incident
	Name      string
	Slug      string
}

//Incident gives information about a services incident and updates
type Incident struct {
	Active    bool
	CreatedAt string
	ID        int64
	Updates   []*Message
}

//Message gives information about an incident
type Message struct {
	Severity     string
	Author       string
	CreatedAt    string
	Translations []*Translation
	UpdatedAt    string
	Content      string
	ID           string
}

//Translation is a translation for a message
type Translation struct {
	Locale    string
	Content   string
	UpdatedAt string
}
