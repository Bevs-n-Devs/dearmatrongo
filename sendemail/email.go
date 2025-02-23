package sendemail

import (
	"net/smtp"
	"os"
	"time"

	"github.com/Bevs-n-Devs/dearmatrongo/env"
	"github.com/Bevs-n-Devs/dearmatrongo/logs"
)

/*
SendEmailUK sends an email to DEARMATRON_UK_EMAIL with the report details:

(name, email, phone number, incident date, facility type, facility name, location, severity, affiliation, description)

It also sends a CC to DEAR_MATRON_RECIEVE_EMAIL as a backup. If the email credentials are not set in the environment, it will attempt to load them from a .env file.
The function logs the email credentials loading process and any errors that occur during the email sending process.
*/
func SendEmailUK(name, email, phoneNumber, incidentDate, facilityType, facilityName, location, severity, affiliation, description string) error {
	logs.Logs(1, "Uploading environment variables to database...")
	if os.Getenv("DEAR_MATRON_SEND_EMAIL") == "" || os.Getenv("DEAR_MATRON_SEND_EMAIL_PASSWORD") == "" || os.Getenv("DEAR_MATRON_RECIEVE_EMAIL") == "" || os.Getenv("DEARMATRON_UK_EMAIL") == "" {
		logs.Logs(2, "Could not get email credentials from Heroku. Loading from .env file...")
		err := env.LoadEnv("env/.env")
		if err != nil {
			logs.Logs(3, "Unable to load environment variables: "+err.Error())
		}
	}

	// SMTP server config
	smptHost := "smtp.gmail.com"
	smptPort := "587"
	smptUser := os.Getenv("DEAR_MATRON_SEND_EMAIL")              // app email
	smptPassword := os.Getenv("DEAR_MATRON_SEND_EMAIL_PASSWORD") // app email password
	recipient := os.Getenv("DEARMATRON_UK_EMAIL")                // 1st destination email
	ccEmail := os.Getenv("DEAR_MATRON_RECIEVE_EMAIL")            // 2nd destination email as backup

	if smptUser == "" || smptPassword == "" || recipient == "" || ccEmail == "" {
		logs.Logs(3, "Email credentials are empty!")
		return nil
	}

	if recipient == ccEmail {
		logs.Logs(2, "Primary and secondary email addresses are the same, skipping CC")
		ccEmail = ""
	}

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
	err := smtp.SendMail(smptHost+":"+smptPort, auth, smptUser, []string{recipient, ccEmail}, []byte("Subject: "+subject+"\n\n"+body))
	if err != nil {
		logs.Logs(3, "Unable to send email: "+err.Error())
		return err
	}
	logs.Logs(1, "Email sent successfully!")
	return nil
}

/*
SendEmailUSA sends an email to DEARMATRON_USA_EMAIL with the report details:

(name, email, phone number, incident date, facility type, facility name, location, severity, affiliation, description)

It also sends a CC to DEAR_MATRON_RECIEVE_EMAIL as a backup. If the email credentials are not set in the environment, it will attempt to load them from a .env file.
The function logs the email credentials loading process and any errors that occur during the email sending process.
*/
func SendEmailUSA(name, email, phoneNumber, incidentDate, facilityType, facilityName, location, severity, affiliation, description string) error {
	logs.Logs(1, "Uploading environment variables to database...")
	if os.Getenv("DEAR_MATRON_SEND_EMAIL") == "" || os.Getenv("DEAR_MATRON_SEND_EMAIL_PASSWORD") == "" || os.Getenv("DEAR_MATRON_RECIEVE_EMAIL") == "" || os.Getenv("DEARMATRON_USA_EMAIL") == "" {
		logs.Logs(2, "Could not get email credentials from Heroku. Loading from .env file...")
		err := env.LoadEnv("env/.env")
		if err != nil {
			logs.Logs(3, "Unable to load environment variables: "+err.Error())
		}
	}

	// SMTP server config
	smptHost := "smtp.gmail.com"
	smptPort := "587"
	smptUser := os.Getenv("DEAR_MATRON_SEND_EMAIL")              // app email
	smptPassword := os.Getenv("DEAR_MATRON_SEND_EMAIL_PASSWORD") // app email password
	recipient := os.Getenv("DEARMATRON_USA_EMAIL")               // 1st destination email
	ccEmail := os.Getenv("DEAR_MATRON_RECIEVE_EMAIL")            // 2nd destination email as backup

	if smptUser == "" || smptPassword == "" || recipient == "" || ccEmail == "" {
		logs.Logs(3, "Email credentials are empty!")
		return nil
	}

	if recipient == ccEmail {
		logs.Logs(2, "Primary and secondary email addresses are the same, skipping CC")
		ccEmail = ""
	}

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
	err := smtp.SendMail(smptHost+":"+smptPort, auth, smptUser, []string{recipient, ccEmail}, []byte("Subject: "+subject+"\n\n"+body))
	if err != nil {
		logs.Logs(3, "Unable to send email: "+err.Error())
		return err
	}
	logs.Logs(1, "Email sent successfully!")
	return nil
}
