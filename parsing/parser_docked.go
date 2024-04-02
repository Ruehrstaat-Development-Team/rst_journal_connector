package parsing

import "encoding/json"

type DockedEvent struct {
	Event                       string             `json:"event"`
	Timestamp                   string             `json:"timestamp"`
	StationName                 string             `json:"StationName"`
	StationType                 string             `json:"StationType"`
	StarSystem                  string             `json:"StarSystem"`
	SystemAddress               int                `json:"SystemAddress"`
	MarketID                    int                `json:"MarketID"`
	StationFaction              StationFaction     `json:"StationFaction"`
	StationGovernment           string             `json:"StationGovernment"`
	StationGovernment_Localised string             `json:"StationGovernment_Localised"`
	StationAllegiance           string             `json:"StationAllegiance"`
	StationServices             []string           `json:"StationServices"`
	StationEconomy              string             `json:"StationEconomy"`
	StationEconomy_Localised    string             `json:"StationEconomy_Localised"`
	StationEconomies            []StationEconomies `json:"StationEconomies"`
	DistFromStarLS              float64            `json:"DistFromStarLS"`
}

type StationFaction struct {
	Name         string  `json:"Name"`
	FactionState *string `json:"FactionState"`
}

type StationEconomies struct {
	Name           string  `json:"Name"`
	Name_Localised string  `json:"Name_Localised"`
	Proportion     float64 `json:"Proportion"`
}

func (e DockedEvent) GetType() string {
	return e.Event
}

func (e DockedEvent) GetTimestamp() string {
	return e.Timestamp
}

type DockedParser struct{}

func (p DockedParser) ParseEvent(eventData []byte) (Event, error) {
	var event DockedEvent
	err := json.Unmarshal([]byte(eventData), &event)
	if err != nil {
		return nil, err
	}

	return event, nil
}
