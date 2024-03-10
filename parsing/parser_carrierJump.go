package parsing

import "encoding/json"

type CarrierJumpEvent struct {
	Event                         string             `json:"event"`
	Timestamp                     string             `json:"timestamp"`
	Docked                        bool               `json:"Docked"`
	StationName                   string             `json:"StationName"`
	StationType                   string             `json:"StationType"`
	MarketID                      int                `json:"MarketID"`
	StationFaction                StationFaction     `json:"StationFaction"`
	StationGovernment             string             `json:"StationGovernment"`
	StationGovernment_Localised   string             `json:"StationGovernment_Localised"`
	StationServices               []string           `json:"StationServices"`
	StationEconomy                string             `json:"StationEconomy"`
	StationEconomy_Localised      string             `json:"StationEconomy_Localised"`
	StationEconomies              []StationEconomies `json:"StationEconomies"`
	StarSystem                    string             `json:"StarSystem"`
	SystemAddress                 int                `json:"SystemAddress"`
	StarPos                       []float64          `json:"StarPos"`
	SystemAllegiance              string             `json:"SystemAllegiance"`
	SystemEconomy                 string             `json:"SystemEconomy"`
	SystemEconomy_Localised       string             `json:"SystemEconomy_Localised"`
	SystemSecondEconomy           string             `json:"SystemSecondEconomy"`
	SystemSecondEconomy_Localised string             `json:"SystemSecondEconomy_Localised"`
	SystemGovernment              string             `json:"SystemGovernment"`
	SystemGovernment_Localised    string             `json:"SystemGovernment_Localised"`
	SystemSecurity                string             `json:"SystemSecurity"`
	SystemSecurity_Localised      string             `json:"SystemSecurity_Localised"`
	Population                    int                `json:"Population"`
	Body                          string             `json:"Body"`
	BodyID                        int                `json:"BodyID"`
	BodyType                      string             `json:"BodyType"`
	SystemFaction                 SystemFaction      `json:"SystemFaction"`
}

type SystemFaction struct {
	Name string `json:"Name"`
}

func (e CarrierJumpEvent) GetType() string {
	return e.Event
}

func (e CarrierJumpEvent) GetTimestamp() string {
	return e.Timestamp
}

type CarrierJumpParser struct{}

func (p CarrierJumpParser) ParseEvent(eventData string) (Event, error) {
	var event CarrierJumpEvent
	err := json.Unmarshal([]byte(eventData), &event)
	if err != nil {
		return nil, err
	}

	return event, nil
}
