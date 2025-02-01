package logs

import (
	"log"
)

var logChannel = make(chan string)

const (
	INFO  = "INFO"
	WARN  = "WARNING"
	ERROR = "ERROR"
)

func ProcessLogs() {
	for logMessage := range logChannel {
		log.Println(logMessage)
	}
}

func Log(logType, logMessage string) {
	var loggedMessage string
	switch logType {
	case INFO:
		loggedMessage = INFO + ": " + logMessage
	case WARN:
		loggedMessage = WARN + "! " + logMessage
	case ERROR:
		loggedMessage = ERROR + "! " + logMessage
	}
	logChannel <- loggedMessage
}
