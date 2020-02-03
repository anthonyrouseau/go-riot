package shard

//Status represents a ShardStatus
type Status struct {
	Name      string     `json:"name"`
	RegionTag string     `json:"region_tag"`
	Hostname  string     `json:"hostname"`
	Services  []*Service `json:"services"`
	Slug      string     `json:"slug"`
	Locales   []string   `json:"locales"`
}

//Service gives information about a service and its incidents
type Service struct {
	Status    string      `json:"status"`
	Incidents []*Incident `json:"incidents"`
	Name      string      `json:"name"`
	Slug      string      `json:"slug"`
}

//Incident gives information about a services incident and updates
type Incident struct {
	Active    bool       `json:"active"`
	CreatedAt string     `json:"created_at"`
	ID        int64      `json:"id"`
	Updates   []*Message `json:"updates"`
}

//Message gives information about an incident
type Message struct {
	Severity     string         `json:"severity"`
	Author       string         `json:"author"`
	CreatedAt    string         `json:"created_at"`
	Translations []*Translation `json:"translations"`
	UpdatedAt    string         `json:"updated_at"`
	Content      string         `json:"content"`
	ID           string         `json:"id"`
}

//Translation is a translation for a message
type Translation struct {
	Locale    string `json:"locale"`
	Content   string `json:"content"`
	UpdatedAt string `json:"updated_at"`
}
