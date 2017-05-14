package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

type ErrLevel int

const (
	ErrDebug ErrLevel = iota
	ErrError
)

func catch(err error, levels ...bool) {
	if err != nil {
		panic(err)
	}
}

func LogWarning(format string, v ...interface{}) {
	logMessage(`WARNING`, format, v...)
}

func LogError(format string, v ...interface{}) {
	logMessage(`ERROR`, format, v...)
	os.Exit(-1)
}

func logMessage(level string, format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(2)
	log.Printf(fmt.Sprintf("[%s] %s:%d ", level, file, line)+format, v...)
}

func calledFrom() string {
	_, file, line, _ := runtime.Caller(1)
	return fmt.Sprintf("%s:%d", file, line)
}
