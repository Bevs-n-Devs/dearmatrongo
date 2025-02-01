package handlers

import (
	"net/http"

	"github.com/Bevs-n-Devs/dearmatrongo/database"
	"github.com/Bevs-n-Devs/dearmatrongo/logs"
	"github.com/Bevs-n-Devs/dearmatrongo/sendemail"
)

func SubmitReport(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		logs.Log(logs.WARN, "Invalid request method: "+r.Method+". Redirecting back to home page.")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// parse from data
	var err = r.ParseForm()
	if err != nil {
		logs.Log(logs.ERROR, "Could not extract data from form:"+err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// extract from fields
	var logMessage string
	name := r.FormValue("name")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	date := r.FormValue("incident_date")
	facilityType := r.FormValue("facility_type")
	facilityName := r.FormValue("facility_name")
	incidentLocation := r.FormValue("incident_location")
	severity := r.FormValue("severity")
	affiliation := r.FormValue("affiliation")
	incidentDescription := r.FormValue("incident_description")
	makeClaim := r.FormValue("make_claim")

	logMessage = "Name: " + name + ", Email: " + email + ", Phone: " + phone + ", Date: " + date + "Facility Type: " + facilityType + ", Facility Name: " + facilityName + ", Incident Location: " + incidentLocation + ", Severity: " + severity + ", Affiliation: " + affiliation + ", Incident Description: " + incidentDescription + ", Make Claim: " + makeClaim
	logs.Log(logs.INFO, "New report created - "+logMessage)
	logs.Log(logs.INFO, "Saving report to database...")
	err = database.InsertDearMatron(name, email, phone, date, facilityType, facilityName, incidentLocation, severity, affiliation, incidentDescription, makeClaim)
	if err != nil {
		logs.Log(logs.ERROR, "Unable to save report to database: "+err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// check if makeClaim == "yes"
	var checkClaim = checkMakeClaim(makeClaim)
	if checkClaim {
		err := sendemail.SendEmailClaim(name, email, phone, date, facilityType, facilityName, incidentLocation, severity, affiliation, incidentDescription)
		if err != nil {
			logs.Log(logs.ERROR, "Unable to send email: "+err.Error())
			logs.Log(logs.WARN, "Report saved to database, but unable to email claim.")
		}
	}
	// redirect to home page
	logs.Log(logs.INFO, "Redirecting to home page...")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func checkMakeClaim(claim string) bool {
	return claim == "yes"
}
