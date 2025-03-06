package utils

import (
	"fmt"
	"time"

	"github.com/Bevs-n-Devs/dearmatrongo/logs"
)

const (
	logError = 3
)

/*
Checks to see if the user wants to make a claim

Returns true if the user wants to make a claim, false otherwise
*/
func MakeCklaimCheck(claim string) bool {
	return claim == "Yes"
}

func ValidateDate(date string) bool {
	dateFormat := "2006-01-02" // Define the expected date format

	incidentDate, err := time.Parse(dateFormat, date)
	if err != nil {
		logs.Logs(logError, fmt.Sprintf("Failed to parse date '%s' with format '%s': %s", date, dateFormat, err.Error()))
		return false
	}

	currentDate := time.Now()

	// compare the incident date with current date
	isBeforeOrEqual := incidentDate.Before(currentDate) || incidentDate.Equal(currentDate)
	return isBeforeOrEqual
}
