package main

import (
	"fmt"
	"strings"

	"github.com/Bevs-n-Devs/dearmatrongo/database"
	"github.com/Bevs-n-Devs/dearmatrongo/handlers"
	"github.com/Bevs-n-Devs/dearmatrongo/logs"
)

func main() {
	go logs.ProcessLogs()
	err := database.ConnectDB()
	if err != nil {
		logs.Logs(3, fmt.Sprintf("Failed to initialize database: %s", err.Error()))
	}
	defer database.CloseDB()

	go handlers.StartTCPServer()
	go func() {
		handlers.StartHTTPServer()

		var templateNames []string
		for _, tmpl := range handlers.Templates.Templates() {
			templateNames = append(templateNames, tmpl.Name())
		}
		logs.Logs(1, "Parsed templates: "+strings.Join(templateNames, ", "))
	}()

	select {}
}
