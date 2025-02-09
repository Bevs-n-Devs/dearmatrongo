package logs

import (
	"log"
)

var logChannel = make(chan string)

const (
	info      = "INFO: "
	warn      = "WARNING! "
	logErr    = "ERROR! "
	database  = "DATABASE: "
	dbError   = "DATABASE ERROR! "
	test      = "TEST: "
	testError = "TEST ERROR! "
)

func ProcessLogs() {
	for logMessage := range logChannel {
		log.Println(logMessage)
	}
}

// logType: 1 = info, 2 = warning, 3 = error, 4 = database, 5 = database error, 6 = test, 7 = test error
func Logs(logType int, logMessage string) {
	var loggedMessage string
	switch logType {
	case 1:
		loggedMessage = info + logMessage
	case 2:
		loggedMessage = warn + logMessage
	case 3:
		loggedMessage = logErr + logMessage
	case 4:
		loggedMessage = database + logMessage
	case 5:
		loggedMessage = dbError + logMessage
	case 6:
		loggedMessage = test + logMessage
	case 7:
		loggedMessage = testError + logMessage
	}
	logChannel <- loggedMessage
}
