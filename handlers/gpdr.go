package handlers

import (
	"fmt"
	"net/http"

	"github.com/Bevs-n-Devs/dearmatrongo/logs"
)

func GDPR(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "gdpr.html", nil)
	if err != nil {
		logs.Logs(3, fmt.Sprintf("Unable to load home page: %s", err.Error()))
		http.Error(w, "Unable to load home page: "+err.Error(), http.StatusInternalServerError)
	}
}
