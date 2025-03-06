package main

import (
	"fmt"
	"strings"

	"github.com/Bevs-n-Devs/dearmatrongo/database"
	"github.com/Bevs-n-Devs/dearmatrongo/handlers"
	"github.com/Bevs-n-Devs/dearmatrongo/logs"
)

const (
	logInfo    = 1
	logDbError = 5
)

func main() {
	go logs.ProcessLogs()
	err := database.ConnectDB()
	if err != nil {
		logs.Logs(logDbError, fmt.Sprintf("Failed to initialize database: %s", err.Error()))
	}
	// defer database.CloseDB() // this keeps the database connection open

	go func() {
		handlers.StartHTTPServer()

		var templateNames []string
		for _, tmpl := range handlers.Templates.Templates() {
			templateNames = append(templateNames, tmpl.Name())
		}
		logs.Logs(logInfo, "Parsed templates: "+strings.Join(templateNames, ", "))
	}()

	select {}
}
