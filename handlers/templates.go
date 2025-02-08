package handlers

import (
	"fmt"
	"html/template"
	"os"

	"github.com/Bevs-n-Devs/dearmatrongo/logs"
)

var Templates *template.Template

// initialise HTML templates
func InitTemplates() {
	var err error
	Templates, err = template.ParseGlob("handlers/templates/*.html")
	if err != nil {
		logs.Logs(3, fmt.Sprintf("Failed to parse templates: %s", err.Error()))
		os.Exit(1)
	}
}
