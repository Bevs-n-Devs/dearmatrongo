package handlers

import (
	"net/http"

	"github.com/Bevs-n-Devs/dearmatrongo/logs"
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
	reportingAs := r.FormValue("reporting_as")
	incidentDescription := r.FormValue("incident_description")
	makeClaim := r.FormValue("make_claim")
	logMessage = "Name: " + name + ", Email: " + email + ", Phone: " + phone + ", Date: " + date + "Facility Type: " + facilityType + ", Facility Name: " + facilityName + ", Incident Location: " + incidentLocation + ", Severity: " + severity + ", Reporting As: " + reportingAs + ", Incident Description: " + incidentDescription + ", Make Claim: " + makeClaim
	logs.Log(logs.INFO, logMessage)

	// print data to server log for debugging
	// log.Println("Name:", name)
	// log.Println("Email:", email)
	// log.Println("Phone:", phone)
	// log.Println("Date:", date)
	// log.Println("Facility Type:", facilityType)
	// log.Println("Facility Name:", facilityName)
	// log.Println("Incident Location:", incidentLocation)
	// log.Println("Severity:", severity)
	// log.Println("Reporting As:", reportingAs)
	// log.Println("Incident Description:", incidentDescription)
	// log.Println("Make Claim:", makeClaim)

	// save to database

	// check if make_claim is "yes"
	// if yes, send data via email to 3rd party service

	// redirect to home page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
