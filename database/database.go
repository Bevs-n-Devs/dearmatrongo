package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "embed"

	_ "github.com/lib/pq"

	"github.com/Bevs-n-Devs/dearmatrongo/env"
	"github.com/Bevs-n-Devs/dearmatrongo/logs"
)

var db *sql.DB

type DearMatronReport struct {
	ID            int    `json:"dear_matron_id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	PhoneNumber   string `json:"phone_number"`
	IncidentDate  string `json:"incident_date"`
	FacilityType  string `json:"facility_type"`
	FacilityName  string `json:"facility_name"`
	Location      string `json:"location"`
	Severity      string `json:"severity"`
	Affiliation   string `json:"affiliation"`
	Description   string `json:"description"`
	MakeClaim     string `json:"make_claim"`
	DateSubmitted string `json:"submitted"`
}

// connect to database via external DB URL
func ConnectDB() error {
	err := env.LoadEnv("env/.env")
	if err != nil {
		logs.Logs(3, fmt.Sprintf("Unable to load environment variables: %s", err.Error()))
		return err
	}
	databaseURL := os.Getenv("DATABASE_URL")

	logs.Logs(4, "Connecting to database...")
	db, err = sql.Open("postgres", databaseURL)
	if err != nil {
		logs.Logs(5, fmt.Sprintf("Unable to open database connection: %s", err.Error()))
		return err
	}
	// verify connection
	logs.Logs(4, "Verifying database connection...")
	if db == nil {
		logs.Logs(5, "Database connection is nil!")
		return errors.New("database connection not establioshed")
	}
	err = db.Ping()
	if err != nil {
		logs.Logs(5, fmt.Sprintf("Cannot connect to database: %s", err.Error()))
		return err
	}
	logs.Logs(4, "Database connection successful!")
	return nil
}

func CloseDB() error {
	if db != nil {
		db.Close()
		logs.Logs(4, "Database connection closed")
		return nil
	}
	logs.Logs(5, "Database connection is not initialized. Could not close database.")
	return errors.New("database connection is not initialized")
}

// insert into dear_matron table
/*
eg.
INSERT INTO public.dear_matron(
	name, email, phone_number, incident_date, facility_type, facility_name, location, severity, affiliation, description, make_claim, submitted)
	VALUES ('john doe', 'jdoe"email.com', '1234567890', '2024-10-26', 'clinic', 'st geroges', 'at home', 'high', 'family member', 'something random', 'yes', NOW());
*/
func InsertDearMatron(name, email, phoneNumber, incidentDate, facilityType, facilityName, location, severity, affiliation, description, makeClaim string) error {
	logs.Logs(4, "Creating new report for Dear Matron...")
	if db == nil {
		logs.Logs(5, "Database connection is not initialized")
		return errors.New("database connection is not initialized")
	}
	// SQL query
	query := `
	INSERT INTO dear_matron (name, email, phone_number, incident_date, facility_type, facility_name, location, severity, affiliation, description, make_claim, submitted)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, NOW());
	`
	// execute query
	_, err := db.Exec(query, name, email, phoneNumber, incidentDate, facilityType, facilityName, location, severity, affiliation, description, makeClaim)
	if err != nil {
		logs.Logs(5, fmt.Sprintf("Unable to create new report for Dear Matron: %s", err.Error()))
		return err
	}
	logs.Logs(4, "New report created for Dear Matron successfully!")
	return nil
}

// get all data from dear_matron_db table
/*
SELECT * FROM dear_matron_db
*/
func GetAllData() (*sql.Rows, error) {
	query := `
	SELECT * FROM dear_matron
	`
	logs.Logs(1, "Retrieving all data from database...")
	if db == nil {
		logs.Logs(5, "Database connection is not initialized")
		return nil, errors.New("database connection is not initialized")
	}
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// get all reports from dear_matron table
/*
SELECT * FROM dear_matron
*/
// returns a list of DearMatronReport structs
func GetAllReports() ([]DearMatronReport, error) {
	query := "SELECT * FROM dear_matron"
	if db == nil {
		logs.Logs(5, "Database connection is not initialized")
		return nil, errors.New("database connection is not initialized")
	}

	rows, err := db.Query(query)
	if err != nil {
		logs.Logs(5, fmt.Sprintf("Unable to retrieve data from database: %s", err.Error()))
		return nil, err
	}
	defer rows.Close()

	// create a slice to hold the results
	var reports []DearMatronReport

	// loop through the rows and add them to the slice
	for rows.Next() {
		var report DearMatronReport
		err := rows.Scan(
			&report.ID,
			&report.Name,
			&report.Email,
			&report.PhoneNumber,
			&report.IncidentDate,
			&report.FacilityType,
			&report.FacilityName,
			&report.Location,
			&report.Severity,
			&report.Affiliation,
			&report.Description,
			&report.MakeClaim,
			&report.DateSubmitted,
		)
		if err != nil {
			logs.Logs(5, fmt.Sprintf("Unable to scan row: %s", err.Error()))
			return nil, err
		}
		reports = append(reports, report)
	}
	// check for errors
	err = rows.Err()
	if err != nil {
		logs.Logs(5, fmt.Sprintf("Unable to retrieve data from database: %s", err.Error()))
		return nil, err
	}
	return reports, nil

}
