package logging

type DebugLevel string

const (
	DebugLevelDebug DebugLevel = "DEBUG"
	DebugLevelInfo  DebugLevel = "INFO"
	DebugLevelWarn  DebugLevel = "WARN"
	DebugLevelError DebugLevel = "ERROR"
)

var debugLevelOrder = map[DebugLevel]int{
	DebugLevelDebug: 0,
	DebugLevelInfo:  10,
	DebugLevelWarn:  20,
	DebugLevelError: 30,
}

var debug = DebugLevelInfo

func SetDebug(value DebugLevel) {
	debug = value
}

func Debug() DebugLevel {
	return debug
}

func IsBelowDebugLevel(level DebugLevel) bool {
	return debugLevelOrder[level] >= debugLevelOrder[debug]
}
