package events

import "encoding/json"

type FileheaderEvent struct {
	Event       string `json:"event"`
	Timestamp   string `json:"timestamp"`
	Language    string `json:"language"`
	Odyssey     bool   `json:"Odyssey"`
	Gameversion string `json:"gameversion"`
	Build       string `json:"build"`
}

func (e FileheaderEvent) GetType() string {
	return e.Event
}

func (e FileheaderEvent) GetTimestamp() string {
	return e.Timestamp
}

type FileheaderParser struct{}

func (p FileheaderParser) ParseEvent(eventData []byte) (Event, error) {
	var event FileheaderEvent
	err := json.Unmarshal([]byte(eventData), &event)
	if err != nil {
		return nil, err
	}

	return event, nil
}
