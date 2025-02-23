package handlers

import (
	"fmt"
	"net/http"

	"github.com/Bevs-n-Devs/dearmatrongo/database"
	"github.com/Bevs-n-Devs/dearmatrongo/encrypt"
	"github.com/Bevs-n-Devs/dearmatrongo/logs"
	"github.com/Bevs-n-Devs/dearmatrongo/sendemail"
	"github.com/Bevs-n-Devs/dearmatrongo/utils"
)

func SubmitUSA(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		logs.Logs(2, fmt.Sprintf("Invalid request method: %s. Redirecting back to Dear Matron USA page.", r.Method))
		http.Redirect(w, r, "/usa", http.StatusSeeOther)
		return
	}

	// parse from data
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

	// encrypt user identifiable data
	encryptName, err := encrypt.Encrypt([]byte(name))
	if err != nil {
		logs.Logs(3, fmt.Sprintf("Unable to encrypt name: %s", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logs.Logs(1, "form name encrypted")

	encryptEmail, err := encrypt.Encrypt([]byte(email))
	if err != nil {
		logs.Logs(3, fmt.Sprintf("Unable to encrypt email: %s", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logs.Logs(1, "form email encrypted")

	encryptPhone, err := encrypt.Encrypt([]byte(phone))
	if err != nil {
		logs.Logs(3, fmt.Sprintf("Unable to encrypt phone: %s", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logs.Logs(1, "form phone encrypted")

	encryptDate, err := encrypt.Encrypt([]byte(date))
	if err != nil {
		logs.Logs(3, fmt.Sprintf("Unable to encrypt date: %s", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logs.Logs(1, "form date encrypted")

	encryptFacilityName, err := encrypt.Encrypt([]byte(facilityName))
	if err != nil {
		logs.Logs(3, fmt.Sprintf("Unable to encrypt facility name: %s", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logs.Logs(1, "form facility name encrypted")

	encryptIncidentLocation, err := encrypt.Encrypt([]byte(incidentLocation))
	if err != nil {
		logs.Logs(3, fmt.Sprintf("Unable to encrypt incident location: %s", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logs.Logs(1, "form incident location encrypted")

	encryptSeverity, err := encrypt.Encrypt([]byte(severity))
	if err != nil {
		logs.Logs(3, fmt.Sprintf("Unable to encrypt severity: %s", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logs.Logs(1, "form severity encrypted")

	encryptAffiliation, err := encrypt.Encrypt([]byte(affiliation))
	if err != nil {
		logs.Logs(3, fmt.Sprintf("Unable to encrypt affiliation: %s", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logs.Logs(1, "form affiliation encrypted")

	encryptIncidentDescription, err := encrypt.Encrypt([]byte(incidentDescription))
	if err != nil {
		logs.Logs(3, fmt.Sprintf("Unable to encrypt incident description: %s", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logs.Logs(1, "form incident description encrypted")

	encryptMakeClaim, err := encrypt.Encrypt([]byte(makeClaim))
	if err != nil {
		logs.Logs(3, fmt.Sprintf("Unable to encrypt make claim: %s", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logs.Logs(1, "form make claim encrypted")

	err = database.InsertDearMatron(
		encryptName,
		encryptEmail,
		encryptPhone,
		encryptDate,
		facilityType,
		encryptFacilityName,
		encryptIncidentLocation,
		encryptSeverity,
		encryptAffiliation,
		encryptIncidentDescription,
		encryptMakeClaim,
		makePublic,
		country,
	)
	if err != nil {
		logs.Logs(3, fmt.Sprintf("Unable to save report to database: %s", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logs.Logs(4, "Ecrypted data saved to database")

	// check if makeClaim == "yes"
	checkClaim := utils.MakeCklaimCheck(makeClaim)
	if checkClaim {
		err := sendemail.SendEmailUSA(name, email, phone, date, facilityType, facilityName, incidentLocation, severity, affiliation, incidentDescription)
		if err != nil {
			logs.Logs(2, fmt.Sprintf("Unable to send email: %s", err.Error()))
		}
	}
	// redirect to home page
	logs.Logs(1, "Redirecting to Dear Matron USA view page...")
	http.Redirect(w, r, "/usa/view", http.StatusSeeOther)
}
