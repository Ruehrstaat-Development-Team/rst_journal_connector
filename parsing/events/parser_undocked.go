package events

import "encoding/json"

type UndockedEvent struct {
	Event       string `json:"event"`
	Timestamp   string `json:"timestamp"`
	StationName string `json:"StationName"`
}

func (e UndockedEvent) GetType() string {
	return e.Event
}

func (e UndockedEvent) GetTimestamp() string {
	return e.Timestamp
}

type UndockedParser struct{}

func (p UndockedParser) ParseEvent(eventData []byte) (Event, error) {
	var event UndockedEvent
	err := json.Unmarshal([]byte(eventData), &event)
	if err != nil {
		return nil, err
	}

	return event, nil
}
