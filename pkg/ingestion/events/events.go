package events

import "time"

type Change struct {
	Source      string    `json:"source"`
	ChangeID    string    `json:"change_id"`
	TimeCreated time.Time `json:"time_created"`
	EventType   string    `json:"event_type"`
	Type        string    `json:"$type"`
}

type Deployment struct {
	Source      string    `json:"source"`
	DeployID    string    `json:"deploy_id"`
	TimeCreated time.Time `json:"time_created"`
	Changes     []string  `json:"changes"`
	Type        string    `json:"$type"`
}

type Incident struct {
	Source       string    `json:"source"`
	IncidentID   string    `json:"incident_id"`
	TimeCreated  time.Time `json:"time_created"`
	TimeResolved time.Time `json:"time_resolved"`
	Type         string    `json:"$type"`
}
