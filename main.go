package mylogger

import "log"

func LogInfo(message string) {
	log.Printf("Info - %v", message)
}
