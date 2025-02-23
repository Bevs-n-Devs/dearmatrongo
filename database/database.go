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

type GetDearMatronReport struct {
	FacilityType string `json:"facility_type"`
	FacilityName string `json:"facility_name"`
	IncidentDate string `json:"incident_date"`
	Location     string `json:"location"`
	Description  string `json:"description"`
}

// connect to database via external DB URL
func ConnectDB() error {
	var err error
	if os.Getenv("DATABASE_URL") == "" {
		logs.Logs(2, "Could not get database URL from Heroku. Loading from .env file...")
		err := env.LoadEnv("env/.env")
		if err != nil {
			logs.Logs(3, fmt.Sprintf("Unable to load environment variables: %s", err.Error()))
			return err
		}
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		logs.Logs(5, "Database URL is empty!")
		return fmt.Errorf("database URL is empty")
	}

	logs.Logs(4, "Connecting to database...")
	db, err = sql.Open("postgres", databaseURL) // open db connection from global db variable
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
	VALUES ('john doe', 'jdoe"email.com', '1234567890', '2024-10-26', 'clinic', 'st geroges', 'at home', 'high', 'family member', 'something random', 'yes', NOW(), 'Yes', 'UK');
*/
func InsertDearMatron(name, email, phoneNumber, incidentDate, facilityType, facilityName, location, severity, affiliation, description, makeClaim, makePublic, country string) error {
	logs.Logs(4, "Creating new report for Dear Matron...")
	if db == nil {
		logs.Logs(5, "Database connection is not initialized")
		return errors.New("database connection is not initialized")
	}
	// SQL query
	query := `
	INSERT INTO dear_matron (name, email, phone_number, incident_date, facility_type, facility_name, location, severity, affiliation, description, make_claim, submitted, make_public, country)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, NOW(), $12, $13);
	`
	// execute query
	_, err := db.Exec(query, name, email, phoneNumber, incidentDate, facilityType, facilityName, location, severity, affiliation, description, makeClaim, makePublic, country)
	if err != nil {
		logs.Logs(5, fmt.Sprintf("Unable to create new report for Dear Matron: %s", err.Error()))
		return err
	}
	logs.Logs(4, "New report created for Dear Matron successfully!")
	return nil
}

/*
Gets all UK reports from dear_matron table

Only for users who are in the UK that want their reports to be public.

Returns a list of GetDearMatronReport structs.
*/
func GetAllReportsUK() ([]GetDearMatronReport, error) {
	query := `
	SELECT facility_type, facility_name, incident_date, location, description
	FROM dear_matron
	WHERE make_public = 'Yes'
	AND country = 'UK';`
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
	var reports []GetDearMatronReport

	// loop through the rows and add them to the slice
	for rows.Next() {
		var report GetDearMatronReport
		err := rows.Scan(
			&report.FacilityType,
			&report.FacilityName,
			&report.IncidentDate,
			&report.Location,
			&report.Description,
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

/*
Gets all USA reports from dear_matron table

Only for users who are in the USA that want their reports to be public.

Returns a slice of GetDearMatronReport structs.
*/
func GetAllReportsUSA() ([]GetDearMatronReport, error) {
	query := `
	SELECT facility_type, facility_name, incident_date, location, description
	FROM dear_matron
	WHERE make_public = 'Yes'
	AND country = 'USA';`
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
	var reports []GetDearMatronReport

	// loop through the rows and add them to the slice
	for rows.Next() {
		var report GetDearMatronReport
		err := rows.Scan(
			&report.FacilityType,
			&report.FacilityName,
			&report.IncidentDate,
			&report.Location,
			&report.Description,
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
