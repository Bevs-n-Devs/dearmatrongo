package handlers

import (
	"fmt"
	"net/http"

	"github.com/Bevs-n-Devs/dearmatrongo/logs"
)

func ReportUSA(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "reportUSA.html", nil)
	if err != nil {
		logs.Logs(3, fmt.Sprintf("Unable to load report page: %s", err.Error()))
		http.Error(w, "Unable to load report page: "+err.Error(), http.StatusInternalServerError)
	}
}
