package handlers

import (
	"fmt"
	"net/http"

	"github.com/Bevs-n-Devs/dearmatrongo/database"
	"github.com/Bevs-n-Devs/dearmatrongo/encrypt"
	"github.com/Bevs-n-Devs/dearmatrongo/logs"
)

func ViewUK(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		logs.Logs(2, fmt.Sprintf("Invalid request method: %s. Redirecting back to home page.", r.Method))
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	getData, err := database.GetAllReportsUK()
	if err != nil {
		logs.Logs(3, fmt.Sprintf("Could not retrieve data from database: %s", err.Error()))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	showData := []database.ShowDearMatronReport{}

	// loop through getData and decrypt data
	logs.Logs(1, "Decrypting data...")
	for index := range getData {
		var convertedData database.ShowDearMatronReport
		convertedData.FacilityType = getData[index].FacilityType

		getData[index].FacilityName, err = encrypt.Decrypt(getData[index].FacilityName) // get the binary number from the database
		if err != nil {
			logs.Logs(3, fmt.Sprintf("Could not decrypt data: %s", err.Error()))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		convertedData.FacilityName = string(getData[index].FacilityName) // convert the binary number to a string

		getData[index].IncidentDate, err = encrypt.Decrypt(getData[index].IncidentDate)
		if err != nil {
			logs.Logs(3, fmt.Sprintf("Could not decrypt data: %s", err.Error()))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		convertedData.IncidentDate = string(getData[index].IncidentDate)

		getData[index].IncidentLocation, err = encrypt.Decrypt(getData[index].IncidentLocation)
		if err != nil {
			logs.Logs(3, fmt.Sprintf("Could not decrypt data: %s", err.Error()))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		convertedData.IncidentLocation = string(getData[index].IncidentLocation)

		getData[index].IncidentDescription, err = encrypt.Decrypt(getData[index].IncidentDescription)
		if err != nil {
			logs.Logs(3, fmt.Sprintf("Could not decrypt data: %s", err.Error()))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		convertedData.IncidentDescription = string(getData[index].IncidentDescription)

		// append the converted data struct to the showData slice
		showData = append(showData, convertedData)
	}
	logs.Logs(1, "Decryption complete.")

	// send data to template
	err = tmplUK.Execute(w, showData)
	if err != nil {
		logs.Logs(3, fmt.Sprintf("Could not execute HTML template: %s", err.Error()))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
