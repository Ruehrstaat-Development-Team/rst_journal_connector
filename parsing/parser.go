package parsing

import (
	"github.com/Ruehrstaat-Development-Team/rst_journal_connector/logging"
	"github.com/Ruehrstaat-Development-Team/rst_journal_connector/parsing/events"
)

var Events = make(map[events.EventMetadata]events.Event)
var UnknownEvents = make(map[events.EventMetadata]string)

type Parser[T events.Event] interface {
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

var parsers = map[string]Parser[events.Event]{
	"Fileheader":  events.FileheaderParser{},
	"Commander":   events.CommanderParser{},
	"Docked":      events.DockedParser{},
	"Undocked":    events.UndockedParser{},
	"CarrierJump": events.CarrierJumpParser{},
	// more parsers...
}
