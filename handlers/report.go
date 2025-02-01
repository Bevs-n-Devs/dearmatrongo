package handlers

import (
	"net/http"

	"github.com/Bevs-n-Devs/dearmatrongo/logs"
)

func ReportPage(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "report.html", nil)
	if err != nil {
		http.Error(w, "Unable to load report page: "+err.Error(), http.StatusInternalServerError)
		logs.Log(logs.ERROR, "Unable to load report page: "+err.Error())
	}
}
