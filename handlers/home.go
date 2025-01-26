package handlers

import (
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		http.Error(w, "Unable to load home page: "+err.Error(), http.StatusInternalServerError)
	}
}
