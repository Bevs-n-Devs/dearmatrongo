package handlers

import (
	"fmt"
	"net/http"

	"github.com/Bevs-n-Devs/dearmatrongo/database"
	"github.com/Bevs-n-Devs/dearmatrongo/logs"
)

func GetReports(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		logs.Logs(2, fmt.Sprintf("Invalid request method: %s. Redirecting back to home page.", r.Method))
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	getData, err := database.GetAllData()
	if err != nil {
		logs.Logs(3, fmt.Sprintf("Could not retrieve data from database: %s", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logs.Logs(1, "Retrieved data from database")
	logs.Logs(1, fmt.Sprintf("Data: %v", &getData))
}
