package handlers

import "html/template"

const (
	UK  = "UK"  // UK report type
	USA = "USA" // USA report type
)

var (
	Templates *template.Template // Templates variable to hold all templates

	tmplUK  = template.Must(template.ParseFiles("./handlers/templates/viewUK.html"))  // Parse the UK view template
	tmplUSA = template.Must(template.ParseFiles("./handlers/templates/viewUSA.html")) // Parse the USA view template
)
