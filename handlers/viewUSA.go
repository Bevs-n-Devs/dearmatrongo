package handlers

import (
	"fmt"
	"net/http"

	"github.com/Bevs-n-Devs/dearmatrongo/database"
	"github.com/Bevs-n-Devs/dearmatrongo/logs"
)

func ViewUSA(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		logs.Logs(2, fmt.Sprintf("Invalid request method: %s. Redirecting back to Dear Matron USA page.", r.Method))
		http.Redirect(w, r, "/usa", http.StatusSeeOther)
		return
	}

	getData, err := database.GetAllReportsUSA()
	if err != nil {
		logs.Logs(3, fmt.Sprintf("Could not retrieve data from database: %s", err.Error()))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// send data to template
	err = tmplUSA.Execute(w, getData)
	if err != nil {
		logs.Logs(3, fmt.Sprintf("Could not execute HTML template: %s", err.Error()))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
