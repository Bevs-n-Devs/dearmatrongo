package handlers

import (
	"fmt"
	"net/http"

	"github.com/Bevs-n-Devs/dearmatrongo/database"
	"github.com/Bevs-n-Devs/dearmatrongo/encrypt"
	"github.com/Bevs-n-Devs/dearmatrongo/logs"
)

func ViewUSA(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		logs.Logs(logWarning, fmt.Sprintf("Invalid request method: %s. Redirecting back to Dear Matron USA page.", r.Method))
		http.Redirect(w, r, "/usa", http.StatusSeeOther)
		return
	}

	getData, err := database.GetAllReportsUSA()
	if err != nil {
		logs.Logs(logError, fmt.Sprintf("Could not retrieve data from database: %s", err.Error()))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	showData := []database.ShowDearMatronReport{}

	// loop through getData and decrypt data
	logs.Logs(logInfo, "Decrypting data...")
	for index := range getData {
		var convertedData database.ShowDearMatronReport
		convertedData.FacilityType = getData[index].FacilityType

		getData[index].FacilityName, err = encrypt.Decrypt(getData[index].FacilityName)
		if err != nil {
			logs.Logs(logError, fmt.Sprintf("Could not decrypt data: %s", err.Error()))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		convertedData.FacilityName = string(getData[index].FacilityName)

		getData[index].IncidentDate, err = encrypt.Decrypt(getData[index].IncidentDate)
		if err != nil {
			logs.Logs(logError, fmt.Sprintf("Could not decrypt data: %s", err.Error()))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		convertedData.IncidentDate = string(getData[index].IncidentDate)

		getData[index].IncidentLocation, err = encrypt.Decrypt(getData[index].IncidentLocation)
		if err != nil {
			logs.Logs(logError, fmt.Sprintf("Could not decrypt data: %s", err.Error()))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		convertedData.IncidentLocation = string(getData[index].IncidentLocation)

		getData[index].IncidentDescription, err = encrypt.Decrypt(getData[index].IncidentDescription)
		if err != nil {
			logs.Logs(logError, fmt.Sprintf("Could not decrypt data: %s", err.Error()))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		convertedData.IncidentDescription = string(getData[index].IncidentDescription)

		showData = append(showData, convertedData)
	}
	logs.Logs(logInfo, "Decryption complete.")

	// send data to template
	err = tmplUSA.Execute(w, showData)
	if err != nil {
		logs.Logs(logError, fmt.Sprintf("Could not execute HTML template: %s", err.Error()))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
