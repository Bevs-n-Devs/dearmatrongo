package main

import (
	"os"

	"github.com/Bevs-n-Devs/dearmatrongo/env"
	"github.com/Bevs-n-Devs/dearmatrongo/handlers"
	"github.com/Bevs-n-Devs/dearmatrongo/logs"
)

func main() {
	go logs.ProcessLogs()
	// load environment variables from the .env file
	err := env.LoadEnv("env/.env")
	if err != nil {
		logs.Log(logs.ERROR, "Unable to load environment variables: "+err.Error())
	}
	var message = os.Getenv("MESSAGE")
	logs.Log(logs.INFO, "Dear Matron app running...")
	logs.Log(logs.INFO, message)

	handlers.StartServer()

	for _, tmpl := range handlers.Templates.Templates() {
		logs.Log(logs.INFO, "Parsing template:"+tmpl.Name())
	}
}
