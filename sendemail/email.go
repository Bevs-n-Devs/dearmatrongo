package sendemail

import (
	"net/smtp"
	"os"
	"time"

	"github.com/Bevs-n-Devs/dearmatrongo/env"
	"github.com/Bevs-n-Devs/dearmatrongo/logs"
)

func SendEmailClaim(name, email, phoneNumber, incidentDate, facilityType, facilityName, location, severity, affiliation, description string) error {
	logs.Logs(1, "Uploading environment variables to database...")
	err := env.LoadEnv("env/.env")
	if err != nil {
		logs.Logs(3, "Unable to load environment variables: "+err.Error())
	}
	// SMTP server config
	smptHost := "smtp.gmail.com"
	smptPort := "587"
	smptUser := os.Getenv("DEAR_MATRON_SEND_EMAIL")              // app email
	smptPassword := os.Getenv("DEAR_MATRON_SEND_EMAIL_PASSWORD") // app email password
	recipient := os.Getenv("DEAR_MATRON_RECIEVE_EMAIL")          // destination email
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
		"Description: " + description + "\n" +
		"Timestamp: " + time.Now().String()
	// send email
	auth := smtp.PlainAuth("", smptUser, smptPassword, smptHost)
	err = smtp.SendMail(smptHost+":"+smptPort, auth, smptUser, []string{recipient}, []byte("Subject: "+subject+"\n\n"+body))
	if err != nil {
		logs.Logs(3, "Unable to send email: "+err.Error())
		return err
	}
	logs.Logs(1, "Email sent successfully!")
	return nil
}
