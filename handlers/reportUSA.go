package handlers

import (
	"fmt"
	"net/http"

	"github.com/Bevs-n-Devs/dearmatrongo/logs"
)

func ReportUSA(w http.ResponseWriter, r *http.Request) {
	errorMsg := r.URL.Query().Get("error")

	// pass the error message to the template
	data := map[string]interface{}{
		"Error": errorMsg,
	}

	err := Templates.ExecuteTemplate(w, "reportUSA.html", data)
	if err != nil {
		logs.Logs(logError, fmt.Sprintf("Unable to load report page: %s", err.Error()))
		http.Error(w, fmt.Sprintf("Unable to load report page: %s", err.Error()), http.StatusInternalServerError)
	}
}
