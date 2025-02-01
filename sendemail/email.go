package sendemail

import (
	"net/smtp"

	"github.com/Bevs-n-Devs/dearmatrongo/logs"
)

func SendEmailClaim(name, email, phoneNumber, incidentDate, facilityType, facilityName, location, severity, affiliation, description string) error {
	// SMTP server config
	var smptHost = "smtp.gmail.com"
	var smptPort = "587"
	var smptUser = "Bevs-n-Devs"     // app email
	var smptPassword = "Bevs-n-Devs" // app email password
	var recipient = "Bevs-n-Devs"    // destination email
	// create email message
	subject := "DEAR MATRON: New Medical Negligence Report"
	body := "Claimant Name: " + name + "\n" +
		"Claimant Email: " + email + "\n" +
		"Claimant Phone Number: " + phoneNumber + "\n" +
		"Incident Date: " + incidentDate + "\n" +
		"Facility Type: " + facilityType + "\n" +
		"Facility Name: " + facilityName + "\n" +
		"Location: " + location + "\n" +
		"Severity: " + severity + "\n" +
		"Affiliation: " + affiliation + "\n" +
		"Description: " + description + "\n"
	// send email
	auth := smtp.PlainAuth("", smptUser, smptPassword, smptHost)
	err := smtp.SendMail(smptHost+":"+smptPort, auth, smptUser, []string{recipient}, []byte("Subject: "+subject+"\n\n"+body))
	if err != nil {
		logs.Log(logs.ERROR, "Unable to send email: "+err.Error())
		return err
	}
	logs.Log(logs.INFO, "Email sent successfully!")
	return nil
}
