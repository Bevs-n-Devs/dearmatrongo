package handlers

import (
	"html/template"
	"net/http"
)

var homepage *template.Template

func init() {
	var err error
	homepage, err = template.ParseGlob("handlers/templates/*.html")
	if err != nil {
		panic(err) // or handle the error in a way that makes sense for your application
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	if homepage == nil {
		http.Error(w, "template not initialized", http.StatusInternalServerError)
		return
	}
	err := homepage.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, "unable to load index page: "+err.Error(), http.StatusInternalServerError)
	}
}
