package parsing

import (
	"Journal-Connector/logging"
)

var Events = make(map[EventMetadata]Event)
var UnknownEvents = make(map[EventMetadata]string)

type Event interface {
	GetType() string
	GetTimestamp() string
}

type EventMetadata struct {
	Event     string `json:"event"`
	Timestamp string `json:"timestamp"`
}

type Parser[T Event] interface {
	ParseEvent(eventData []byte) (T, error)
}

var log = logging.Logger{Package: "parser"}

type DebugLevel string

const (
	DebugLevelDebug DebugLevel = "DEBUG"
	DebugLevelInfo  DebugLevel = "INFO"
	DebugLevelWarn  DebugLevel = "WARN"
	DebugLevelError DebugLevel = "ERROR"
)

var debug = DebugLevelInfo

func SetDebug(value DebugLevel) {
	debug = value
}

var parsers = map[string]Parser[Event]{
	"Fileheader":  FileheaderParser{},
	"Commander":   CommanderParser{},
	"Docked":      DockedParser{},
	"Undocked":    UndockedParser{},
	"CarrierJump": CarrierJumpParser{},
	// more parsers...
}
