package mylogger

import (
	"log"
)

type RomitLogger struct{}

func (l *RomitLogger) LogInfo(message string) {
	log.Printf("INFO %s", message)
}

func (l *RomitLogger) LogWarning(message string) {
	log.Printf("WARN %s", message)
}
