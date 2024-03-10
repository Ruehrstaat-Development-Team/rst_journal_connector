package parsing

import "encoding/json"

type CommanderEvent struct {
	Event     string `json:"event"`
	Timestamp string `json:"timestamp"`
	FID       string `json:"FID"`
	Name      string `json:"Name"`
}

func (e CommanderEvent) GetType() string {
	return e.Event
}

func (e CommanderEvent) GetTimestamp() string {
	return e.Timestamp
}

type CommanderParser struct{}

func (p CommanderParser) ParseEvent(eventData string) (Event, error) {
	var event CommanderEvent
	err := json.Unmarshal([]byte(eventData), &event)
	if err != nil {
		return nil, err
	}

	return event, nil
}
