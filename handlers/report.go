package handlers

import (
	"net/http"
)

func ReportPage(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "report.html", nil)
	if err != nil {
		http.Error(w, "Unable to load report page: "+err.Error(), http.StatusInternalServerError)
	}
}
