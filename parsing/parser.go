package parsing

import (
	"Journal-Connector/logging"
)

type Event interface {
	GetType() string
	GetTimestamp() string
}

type Parser[T Event] interface {
	ParseEvent(eventData string) (T, error)
}

var log = logging.Logger{Package: "parser"}
