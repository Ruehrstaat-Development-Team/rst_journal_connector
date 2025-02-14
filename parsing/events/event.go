package events

type Event interface {
	GetType() string
	GetTimestamp() string
}

type EventMetadata struct {
	Event     string `json:"event"`
	Timestamp string `json:"timestamp"`
}
