package handlers

import (
	"fmt"
	"net/http"

	"github.com/Bevs-n-Devs/dearmatrongo/logs"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		logs.Logs(3, fmt.Sprintf("Unable to load home page: %s", err.Error()))
		http.Error(w, "Unable to load home page: "+err.Error(), http.StatusInternalServerError)
	}
}
