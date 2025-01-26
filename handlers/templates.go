package handlers

import (
	"html/template"
	"log"
)

var Templates *template.Template

func InitTemplates() {
	var err error
	Templates, err = template.ParseGlob("handlers/templates/*.html")
	if err != nil {
		log.Fatalf("Failed to parse templates: %v", err.Error())
	}
}
