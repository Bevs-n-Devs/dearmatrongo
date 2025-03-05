package sendemail

import (
	"fmt"
	"net/smtp"
	"os"
	"time"

	"github.com/Bevs-n-Devs/dearmatrongo/env"
	"github.com/Bevs-n-Devs/dearmatrongo/logs"
)

const (
	logInfo    = 1
	logWarning = 2
	logError   = 3
)

/*
SendEmailUK sends an email to DEARMATRON_UK_EMAIL with the report details:

(name, email, phone number, incident date, facility type, facility name, location, severity, affiliation, description)

It also sends a CC to DEAR_MATRON_RECIEVE_EMAIL as a backup. If the email credentials are not set in the environment, it will attempt to load them from a .env file.
The function logs the email credentials loading process and any errors that occur during the email sending process.
*/
func SendEmailUK(name, email, phoneNumber, incidentDate, facilityType, facilityName, location, severity, affiliation, description string) error {
	logs.Logs(logInfo, "Uploading environment variables to database...")
	if os.Getenv("DEAR_MATRON_SEND_EMAIL") == "" || os.Getenv("DEAR_MATRON_SEND_EMAIL_PASSWORD") == "" || os.Getenv("DEAR_MATRON_RECIEVE_EMAIL") == "" || os.Getenv("DEARMATRON_UK_EMAIL") == "" {
		logs.Logs(logWarning, "Could not get email credentials from Heroku. Loading from .env file...")
		err := env.LoadEnv("env/.env")
		if err != nil {
			logs.Logs(logError, "Unable to load environment variables: "+err.Error())
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
		logs.Logs(logError, "Email credentials are empty!")
		return nil
	}

	if recipient == ccEmail {
		logs.Logs(logWarning, "Primary and secondary email addresses are the same, skipping CC")
		ccEmail = ""
	}

	// create email message
	subject := fmt.Sprintf("DEAR MATRON: New Medical Negligence Report for %s", name)
	body := fmt.Sprintf(`
	Claimant Name: %s
	Claimant Email: %s
	Claimant Phone Number: %s
	Incident Date: %s
	Facility Type: %s
	Facility Name: %s
	Location of Incident: %s
	Severity of Incident: %s
	Affiliation: %s
	Description of Incident: %s
	Timestamp: %s

	Dear Matron
	dearmatron@gmail.com

	https://dearmatron.com
	`, name, email, phoneNumber, incidentDate, facilityType, facilityName, location, severity, affiliation, description, time.Now().Format("2006-01-02 15:04:05"))
	// send email
	auth := smtp.PlainAuth("", smptUser, smptPassword, smptHost)
	err := smtp.SendMail(smptHost+":"+smptPort, auth, smptUser, []string{recipient, ccEmail}, []byte("Subject: "+subject+"\n\n"+body))
	if err != nil {
		logs.Logs(logError, "Unable to send email: "+err.Error())
		return err
	}
	logs.Logs(logInfo, "Email sent successfully!")
	return nil
}

/*
SendEmailUSA sends an email to DEARMATRON_USA_EMAIL with the report details:

(name, email, phone number, incident date, facility type, facility name, location, severity, affiliation, description)

It also sends a CC to DEAR_MATRON_RECIEVE_EMAIL as a backup. If the email credentials are not set in the environment, it will attempt to load them from a .env file.
The function logs the email credentials loading process and any errors that occur during the email sending process.
*/
func SendEmailUSA(name, email, phoneNumber, incidentDate, facilityType, facilityName, location, severity, affiliation, description string) error {
	logs.Logs(logInfo, "Uploading environment variables to database...")
	if os.Getenv("DEAR_MATRON_SEND_EMAIL") == "" || os.Getenv("DEAR_MATRON_SEND_EMAIL_PASSWORD") == "" || os.Getenv("DEAR_MATRON_RECIEVE_EMAIL") == "" || os.Getenv("DEARMATRON_USA_EMAIL") == "" {
		logs.Logs(logWarning, "Could not get email credentials from Heroku. Loading from .env file...")
		err := env.LoadEnv("env/.env")
		if err != nil {
			logs.Logs(logError, "Unable to load environment variables: "+err.Error())
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
		logs.Logs(logError, "Email credentials are empty!")
		return nil
	}

	if recipient == ccEmail {
		logs.Logs(logWarning, "Primary and secondary email addresses are the same, skipping CC")
		ccEmail = ""
	}

	// create email message
	subject := fmt.Sprintf("DEAR MATRON: New Medical Malpractice Report for %s", name)
	body := fmt.Sprintf(`
	Claimant Name: %s
	Claimant Email: %s
	Claimant Phone Number: %s
	Incident Date: %s
	Facility Type: %s
	Facility Name: %s
	Location of Incident: %s
	Severity of Incident: %s
	Affiliation: %s
	Description of Incident: %s
	Timestamp: %s

	Dear Matron
	dearmatron@gmail.com

	https://dearmatron.com
	`, name, email, phoneNumber, incidentDate, facilityType, facilityName, location, severity, affiliation, description, time.Now().Format("2006-01-02 15:04:05"))
	// send email
	auth := smtp.PlainAuth("", smptUser, smptPassword, smptHost)
	err := smtp.SendMail(smptHost+":"+smptPort, auth, smptUser, []string{recipient, ccEmail}, []byte("Subject: "+subject+"\n\n"+body))
	if err != nil {
		logs.Logs(logError, "Unable to send email: "+err.Error())
		return err
	}
	logs.Logs(logInfo, "Email sent successfully!")
	return nil
}

func SearchDataPolicyDataFailed(name, email string) error {
	if os.Getenv("DEAR_MATRON_SEND_EMAIL") == "" || os.Getenv("DEAR_MATRON_SEND_EMAIL_PASSWORD") == "" || os.Getenv("DEAR_MATRON_RECIEVE_EMAIL") == "" {
		logs.Logs(logWarning, "Could not get email credentials from Heroku. Loading from .env file...")
		err := env.LoadEnv("env/.env")
		if err != nil {
			logs.Logs(logError, "Unable to load environment variables: "+err.Error())
		}
	}

	// SMTP server config
	smptHost := "smtp.gmail.com"
	smptPort := "587"
	smptUser := os.Getenv("DEAR_MATRON_SEND_EMAIL")              // app email
	smptPassword := os.Getenv("DEAR_MATRON_SEND_EMAIL_PASSWORD") // app email password
	ccEmail := os.Getenv("DEAR_MATRON_RECIEVE_EMAIL")            // 2nd destination email as backup

	if smptUser == "" || smptPassword == "" || ccEmail == "" {
		logs.Logs(logError, "Email credentials are empty!")
		return nil
	}

	subject := fmt.Sprintf("DEAR MATRON: Data could not be deleted for %s", name)
	body := `
We received your request to delete your data from our records. However, we were unable to locate a matching report based on the information you provided.

Because we use strong encryption to protect your data, all details must be entered exactly as they were originally submitted—including names, dates, and facility details. Even small differences in spelling, formatting, or spacing can prevent us from finding your record.

If you believe there may have been a mistake in your submission, we kindly ask you to try again using the exact details from your original report.

Alternatively, if you are unable to locate the correct details, please contact us directly at dearmatron@gmail.com. To help us conduct a more thorough search, please include all the details you originally provided, including:

- Full Name
- Email Address
- Phone Number (if applicable)
- Date of Incident
- Facility Type 
- Facility Name
- Incident Location
- Severity of Incident (low, moderate, or severe)
- Your Affiliation (patient, family member, employee, other)

Once we have verified the details, we will make every effort to process your deletion request promptly.

If you have any further questions, please do not hesitate to reach out.

Best regards,

Dear Matron
dearmatron@gmail.com

https://dearmatron.com
	`
	auth := smtp.PlainAuth("", smptUser, smptPassword, smptHost)
	err := smtp.SendMail(smptHost+":"+smptPort, auth, smptUser, []string{email, ccEmail}, []byte("Subject: "+subject+"\n\n"+body))
	if err != nil {
		logs.Logs(logError, "Unable to send email: "+err.Error())
		return err
	}
	logs.Logs(logInfo, "Email sent successfully!")
	return nil
}

func SearchDataPolicyDataFound(name, email, phoneNumber, incidentDate, facilityType, facilityName, location, severity, affiliation, description, claim, public, country, sendToEmail string) error {
	if os.Getenv("DEAR_MATRON_SEND_EMAIL") == "" || os.Getenv("DEAR_MATRON_SEND_EMAIL_PASSWORD") == "" || os.Getenv("DEAR_MATRON_RECIEVE_EMAIL") == "" {
		logs.Logs(logWarning, "Could not get email credentials from Heroku. Loading from .env file...")
		err := env.LoadEnv("env/.env")
		if err != nil {
			logs.Logs(logError, "Unable to load environment variables: "+err.Error())
		}
	}

	// SMTP server config
	smptHost := "smtp.gmail.com"
	smptPort := "587"
	smptUser := os.Getenv("DEAR_MATRON_SEND_EMAIL")              // app email
	smptPassword := os.Getenv("DEAR_MATRON_SEND_EMAIL_PASSWORD") // app email password
	ccEmail := os.Getenv("DEAR_MATRON_RECIEVE_EMAIL")            // 2nd destination email as backup

	if smptUser == "" || smptPassword == "" || ccEmail == "" {
		logs.Logs(logError, "Email credentials are empty!")
		return nil
	}

	subject := fmt.Sprintf("DEAR MATRON: Data deleted for %s", name)
	body := fmt.Sprintf(`
Your data has been successfully deleted from our records.

Name: %s
Email: %s
Phone Number: %s
Incident Date: %s
Facility Type: %s
Facility Name: %s
Location: %s
Severity: %s
Affiliation: %s
Description: %s
Claim: %s
Public: %s
Country: %s

If you have any further questions, please do not hesitate to reach out.

Best regards,

Dear Matron
dearmatron@gmail.com

https://dearmatron.com
`, name, email, phoneNumber, incidentDate, facilityType, facilityName, location, severity, affiliation, description, claim, public, country)

	// Create a new MIME message
	msg := "From: " + smptUser + "\n" +
		"To: " + sendToEmail + "\n" +
		"Cc: " + ccEmail + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	// Authenticate with the SMTP server
	auth := smtp.PlainAuth("", smptUser, smptPassword, smptHost)

	// Send the email
	err := smtp.SendMail(smptHost+":"+smptPort, auth, smptUser, []string{sendToEmail, ccEmail}, []byte(msg))
	if err != nil {
		logs.Logs(logError, "Unable to send email: "+err.Error())
		return err
	}
	logs.Logs(logInfo, "Email sent successfully!")
	return nil
}
