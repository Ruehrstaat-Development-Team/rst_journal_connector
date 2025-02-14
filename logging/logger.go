package logging

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Logger struct {
	Package string
}

func InitLogSys() *os.File {
	logFile, err := os.OpenFile("connector.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0600)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	return logFile
}

func (l Logger) Print(v ...any) {
	log.Printf("[ %s ] %s", l.Package, fmt.Sprint(v...))
}

func (l Logger) Println(v ...any) {
	log.Printf("[ %s ] %s", l.Package, fmt.Sprintln(v...))
}

func (l Logger) Printf(format string, v ...any) {
	log.Printf("[ %s ] %s", l.Package, fmt.Sprintf(format, v...))
}

func (l Logger) Fatal(v ...any) {
	log.Fatalf("[ %s ] %s", l.Package, fmt.Sprintln(v...))
}

func (l Logger) Fatalln(v ...any) {
	log.Fatalf("[ %s ] %s", l.Package, fmt.Sprintln(v...))
}

func (l Logger) Fatalf(format string, v ...any) {
	log.Fatalf("[ %s ] %s", l.Package, fmt.Sprintf(format, v...))
}

func (l Logger) PrintWithLevel(level DebugLevel, v ...any) {
	if IsBelowDebugLevel(level) {
		l.Print(v...)
	}
}

func (l Logger) PrintlnWithLevel(level DebugLevel, v ...any) {
	if IsBelowDebugLevel(level) {
		l.Println(v...)
	}
}

func (l Logger) PrintfWithLevel(level DebugLevel, format string, v ...any) {
	if IsBelowDebugLevel(level) {
		l.Printf(format, v...)
	}
}
