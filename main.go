package main

import (
	"log"

	"github.com/Bevs-n-Devs/dearmatrongo/handlers"
)

func main() {
	log.Println("Hello world, hello Yaw!")

	handlers.StartServer()

	for _, tmpl := range handlers.Templates.Templates() {
		log.Println("Parsed template:", tmpl.Name())
	}
}
