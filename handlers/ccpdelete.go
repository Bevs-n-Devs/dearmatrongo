package handlers

import (
	"fmt"
	"net/http"

	"github.com/Bevs-n-Devs/dearmatrongo/database"
	"github.com/Bevs-n-Devs/dearmatrongo/encrypt"
	"github.com/Bevs-n-Devs/dearmatrongo/logs"
	"github.com/Bevs-n-Devs/dearmatrongo/sendemail"
)

func DeleteCCPAData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		logs.Logs(2, fmt.Sprintf("Invalid request method: %s. Redirecting back to GDPR page.", r.Method))
		http.Redirect(w, r, "/usa/data-policy", http.StatusSeeOther)
		return
	}

	// parse form data
	err := r.ParseForm()
	if err != nil {
		logs.Logs(3, fmt.Sprintf("Could not extract data from form: %s", err.Error()))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// extract from fields
	const country = USA
	name := r.FormValue("name")
	email := r.FormValue("email")
	incidentDate := r.FormValue("incident_date")
	facilityType := r.FormValue("facility_type")
	facilityName := r.FormValue("facility_name")
	sendDataEmail := r.FormValue("send_data_email")

	if name == "" || email == "" || incidentDate == "" || facilityType == "" || sendDataEmail == "" || facilityName == "" {
		logs.Logs(3, "Missing required form fields")
		http.Error(w, "Missing required form fields", http.StatusBadRequest)
		return
	}

	// Check if GDPR data exists in the database
	match, err := database.CheckGDPRData(name, email, incidentDate, facilityType, facilityName)
	if err != nil {
		logs.Logs(3, fmt.Sprintf("Error occured checking GDPR data: %s", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// if there is no matching data, redirect back to GDPR page
	if !match {
		// send email to user to confirm no match
		err := sendemail.SearchGDPRDataFailed(name, sendDataEmail)
		if err != nil {
			logs.Logs(3, fmt.Sprintf("Could not send email: %s", err.Error()))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		logs.Logs(2, "GDPR data not found in database. Redirecting back to GDPR page.")
		http.Redirect(w, r, "/usa/data-policy", http.StatusSeeOther)
		return
	}

	// Retrieve and decrypt GDPR data from the database if a match is found
	getData, err := database.DearMatronFullReport(name, email, incidentDate, facilityType, facilityName, country)
	if err != nil {
		logs.Logs(3, fmt.Sprintf("Could not retrieve data from database: %s", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// decrypt user data
	var convertedData database.ShowDearMatronFullReport

	convertedData.Name = name
	convertedData.FacilityType = getData.FacilityType
	convertedData.MakePublic = getData.MakePublic
	convertedData.Country = getData.Country

	convertedData.FacilityName = string(getData.FacilityName)

	getData.Email, err = encrypt.Decrypt(getData.Email)
	if err != nil {
		logs.Logs(3, fmt.Sprintf("Could not decrypt data: %s", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	convertedData.Email = string(getData.Email)

	getData.PhoneNumber, err = encrypt.Decrypt(getData.PhoneNumber)
	if err != nil {
		logs.Logs(3, fmt.Sprintf("Could not decrypt data: %s", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	convertedData.PhoneNumber = string(getData.PhoneNumber)

	getData.IncidentDate, err = encrypt.Decrypt(getData.IncidentDate)
	if err != nil {
		logs.Logs(3, fmt.Sprintf("Could not decrypt data: %s", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	convertedData.IncidentDate = string(getData.IncidentDate)

	getData.FacilityName, err = encrypt.Decrypt(getData.FacilityName)
	if err != nil {
		logs.Logs(3, fmt.Sprintf("Could not decrypt data: %s", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	convertedData.FacilityName = string(getData.FacilityName)

	getData.IncidentLocation, err = encrypt.Decrypt(getData.IncidentLocation)
	if err != nil {
		logs.Logs(3, fmt.Sprintf("Could not decrypt data: %s", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	convertedData.IncidentLocation = string(getData.IncidentLocation)

	getData.Severity, err = encrypt.Decrypt(getData.Severity)
	if err != nil {
		logs.Logs(3, fmt.Sprintf("Could not decrypt data: %s", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	convertedData.Severity = string(getData.Severity)

	getData.Affiliation, err = encrypt.Decrypt(getData.Affiliation)
	if err != nil {
		logs.Logs(3, fmt.Sprintf("Could not decrypt data: %s", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	convertedData.Affiliation = string(getData.Affiliation)

	getData.IncidentDescription, err = encrypt.Decrypt(getData.IncidentDescription)
	if err != nil {
		logs.Logs(3, fmt.Sprintf("Could not decrypt data: %s", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	convertedData.IncidentDescription = string(getData.IncidentDescription)

	getData.MakeClaim, err = encrypt.Decrypt(getData.MakeClaim)
	if err != nil {
		logs.Logs(3, fmt.Sprintf("Could not decrypt data: %s", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	convertedData.MakeClaim = string(getData.MakeClaim)

	// Send confirmation email to the user with the decrypted GDPR data
	err = sendemail.SearchGDPRDataFound(
		convertedData.Name,
		convertedData.Email,
		convertedData.PhoneNumber,
		convertedData.IncidentDate,
		convertedData.FacilityType,
		convertedData.FacilityName,
		convertedData.IncidentLocation,
		convertedData.Severity,
		convertedData.Affiliation,
		convertedData.IncidentDescription,
		convertedData.MakeClaim,
		convertedData.MakePublic,
		convertedData.Country,
		sendDataEmail,
	)
	if err != nil {
		logs.Logs(3, fmt.Sprintf("Could not send email: %s", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Delete the GDPR data from the database after sending confirmation
	err = database.DeleteGDPRData(name, email, incidentDate, facilityType, facilityName, country)
	if err != nil {
		logs.Logs(3, fmt.Sprintf("Could not delete data: %s", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logs.Logs(1, "GDPR data deleted successfully")
	http.Redirect(w, r, "/usa", http.StatusSeeOther)
}
