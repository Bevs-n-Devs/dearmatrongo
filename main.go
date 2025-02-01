package main

import (
	"github.com/Bevs-n-Devs/dearmatrongo/handlers"
	"github.com/Bevs-n-Devs/dearmatrongo/logs"
)

func main() {
	go logs.ProcessLogs()
	logs.Log(logs.INFO, "hello world, hello Yaw!")

	handlers.StartServer()

	for _, tmpl := range handlers.Templates.Templates() {
		logs.Log(logs.INFO, "Parsing template:"+tmpl.Name())
	}
}
