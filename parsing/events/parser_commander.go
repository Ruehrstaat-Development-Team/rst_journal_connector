package events

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

func (p CommanderParser) ParseEvent(eventData []byte) (Event, error) {
	var event CommanderEvent
	err := json.Unmarshal(eventData, &event)
	if err != nil {
		return nil, err
	}

	return event, nil
}
