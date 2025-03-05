package handlers

import (
	"fmt"
	"net/http"

	"github.com/Bevs-n-Devs/dearmatrongo/database"
	"github.com/Bevs-n-Devs/dearmatrongo/logs"
	"github.com/Bevs-n-Devs/dearmatrongo/sendemail"
	"github.com/Bevs-n-Devs/dearmatrongo/utils"
)

func SubmitUK(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		logs.Logs(logWarning, fmt.Sprintf("Invalid request method: %s. Redirecting back to home page.", r.Method))
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// parse from data
	err := r.ParseForm()
	if err != nil {
		logs.Logs(logError, fmt.Sprintf("Could not extract data from form: %s", err.Error()))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// extract from fields
	const country = UK
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
	makePublic := r.FormValue("make_public")

	err = database.CreateDearMatronReport(
		name,
		email,
		phone,
		date,
		facilityType,
		facilityName,
		incidentLocation,
		severity,
		affiliation,
		incidentDescription,
		makeClaim,
		makePublic,
		country,
	)

	if err != nil {
		logs.Logs(logError, fmt.Sprintf("Unable to save report to database: %s", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logs.Logs(logDb, "Ecrypted data saved to database")

	// check if makeClaim == "yes"
	checkClaim := utils.MakeCklaimCheck(makeClaim)
	if checkClaim {
		err := sendemail.SendEmailUK(name, email, phone, date, facilityType, facilityName, incidentLocation, severity, affiliation, incidentDescription)
		if err != nil {
			logs.Logs(logWarning, fmt.Sprintf("Unable to send email: %s", err.Error()))
		}
	}
	// redirect to home page
	logs.Logs(logInfo, "Redirecting to Dear Matron UK view page...")
	http.Redirect(w, r, "/uk/view", http.StatusSeeOther)
}
