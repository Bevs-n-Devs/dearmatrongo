package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "embed"

	_ "github.com/lib/pq"

	"github.com/Bevs-n-Devs/dearmatrongo/encrypt"
	"github.com/Bevs-n-Devs/dearmatrongo/env"
	"github.com/Bevs-n-Devs/dearmatrongo/logs"
)

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

/*
CreateDearMatronReport creates a new entry in the dearmatron table in the database
This function will hash and encrypt the data before inserting it into the database.
If the database connection is not established, it will return an error.

Parameters:

- name: the name of the person making the report

- email: the email of the person making the report

- phoneNumber: the phone number of the person making the report

- incidentDate: the date of the incident

- facilityType: the type of facility the incident occurred in (e.g. hospital, clinic, etc.)

- facilityName: the name of the facility the incident occurred in

- incidentLocation: the location of the incident

- severity: the severity of the incident (e.g. high, medium, low)

- affiliation: the affiliation of the person making the report (e.g. family member, friend, etc.)

- description: the description of the incident

- makeClaim: whether the person making the report wants to make a claim

- makePublic: whether the person making the report wants to make the report public

- country: the country where the incident occurred

Returns:

- error: an error if the database connection is not established or if there is a problem inserting the data into the database
*/
func CreateDearMatronReport(name, email, phoneNumber, incidentDate, facilityType, facilityName, incidentLocation, severity, affiliation, description, makeClaim, makePublic, country string) error {
	if db == nil {
		logs.Logs(5, "Database connection is not initialized")
		return errors.New("database connection is not initialized")
	}

	// hash & encrypt data
	hashName := encrypt.HashData(name)
	encryptName, err := encrypt.Encrypt([]byte(name))
	if err != nil {
		logs.Logs(5, fmt.Sprintf("Could not encrypt name: %s", err.Error()))
		return err
	}

	hashEmail := encrypt.HashData(email)
	encryptEmail, err := encrypt.Encrypt([]byte(email))
	if err != nil {
		logs.Logs(5, fmt.Sprintf("Could not encrypt email: %s", err.Error()))
		return err
	}

	encryptNumber, err := encrypt.Encrypt([]byte(phoneNumber))
	if err != nil {
		logs.Logs(5, fmt.Sprintf("Could not encrypt phone number: %s", err.Error()))
		return err
	}

	hashDate := encrypt.HashData(incidentDate)
	encryptDate, err := encrypt.Encrypt([]byte(incidentDate))
	if err != nil {
		logs.Logs(5, fmt.Sprintf("Could not encrypt incident date: %s", err.Error()))
		return err
	}

	hashFacilityName := encrypt.HashData(facilityName)
	encryptFacilityName, err := encrypt.Encrypt([]byte(facilityName))
	if err != nil {
		logs.Logs(5, fmt.Sprintf("Could not encrypt facility name: %s", err.Error()))
		return err
	}

	encryptIncidentLocation, err := encrypt.Encrypt([]byte(incidentLocation))
	if err != nil {
		logs.Logs(5, fmt.Sprintf("Could not encrypt incident location: %s", err.Error()))
		return err
	}

	encryptSeverity, err := encrypt.Encrypt([]byte(severity))
	if err != nil {
		logs.Logs(5, fmt.Sprintf("Could not encrypt severity: %s", err.Error()))
		return err
	}

	encryptAffiliation, err := encrypt.Encrypt([]byte(affiliation))
	if err != nil {
		logs.Logs(5, fmt.Sprintf("Could not encrypt affiliation: %s", err.Error()))
		return err
	}

	encryptDescription, err := encrypt.Encrypt([]byte(description))
	if err != nil {
		logs.Logs(5, fmt.Sprintf("Could not encrypt description: %s", err.Error()))
		return err
	}

	encryptMakeClaim, err := encrypt.Encrypt([]byte(makeClaim))
	if err != nil {
		logs.Logs(5, fmt.Sprintf("Could not encrypt make claim: %s", err.Error()))
		return err
	}

	// SQL query
	query := `
	INSERT INTO tbl_dearmatron (
		name,
		hash_name, 
		email,
		hash_email, 
		phone_number, 
		incident_date,
		hash_incident_date, 
		facility_type, 
		facility_name,
		hash_facility_name, 
		incident_location, 
		severity, 
		affiliation, 
		incident_description, 
		make_claim,  
		make_public, 
		country, 
		submitted
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, NOW());
	`

	// execute query
	_, err = db.Exec(
		query,
		encryptName,
		hashName,
		encryptEmail,
		hashEmail,
		encryptNumber,
		encryptDate,
		hashDate,
		facilityType,
		encryptFacilityName,
		hashFacilityName,
		encryptIncidentLocation,
		encryptSeverity,
		encryptAffiliation,
		encryptDescription,
		encryptMakeClaim,
		makePublic,
		country,
	)
	if err != nil {
		logs.Logs(5, fmt.Sprintf("Could not insert data into database: %s", err.Error()))
		return err
	}

	return nil
}

/*
Gets all UK reports from tbl_dearmatron table

Only for users who are in the UK that want their reports to be public.

Returns a list of GetDearMatronReport structs.
*/
func GetAllReportsUK() ([]GetDearMatronReport, error) {
	query := `
	SELECT facility_type, facility_name, incident_date, incident_location, incident_description
	FROM tbl_dearmatron
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
			&report.IncidentLocation,
			&report.IncidentDescription,
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
Gets all USA reports from tbl_dearmatron table

Only for users who are in the USA that want their reports to be public.

Returns a slice of GetDearMatronReport structs.
*/
func GetAllReportsUSA() ([]GetDearMatronReport, error) {
	query := `
	SELECT facility_type, facility_name, incident_date, incident_location, incident_description
	FROM tbl_dearmatron
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
			&report.IncidentLocation,
			&report.IncidentDescription,
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
Checks if a report exists in the database based on the parameters provided.

Parameters:

- name: name of the user.

- email: email of the user.

- incidentDate: date of the incident.

- facilityType: Type of the facility.

- facilityName: name of the facility.

Returns:

- A boolean indicating whether a matching report exists in the database.
- An error if the database query fails.
*/
func CheckGDPRData(name, email, incidentDate, facilityType, facilityName string) (bool, error) {
	// validate input
	if name == "" || email == "" || incidentDate == "" || facilityName == "" {
		return false, errors.New("all parameters are required")
	}

	// hash params
	hashName := encrypt.HashData(name)
	hashEmail := encrypt.HashData(email)
	hashIncidentDate := encrypt.HashData(incidentDate)
	hashFacilityName := encrypt.HashData(facilityName)

	// create a query with encrypted params
	query := `
		SELECT hash_name, hash_email, hash_incident_date, facility_type, hash_facility_name
		FROM tbl_dearmatron
		WHERE hash_name = $1
		AND hash_email = $2
		AND hash_incident_date = $3
		AND facility_type = $4
		AND hash_facility_name = $5;`

	// execute query to find matching data within database
	var dbHashName string
	var dbHashEmail string
	var dbHashIncidentDate string
	var dbFacilityType string
	var dbHashFacilityName string

	err := db.QueryRow(
		query,
		hashName,
		hashEmail,
		hashIncidentDate,
		facilityType,
		hashFacilityName,
	).Scan(
		&dbHashName,
		&dbHashEmail,
		&dbHashIncidentDate,
		&dbFacilityType,
		&dbHashFacilityName,
	)

	if err == sql.ErrNoRows {
		logs.Logs(2, "No matching data found in database")
		return false, nil
	}

	if err != nil {
		logs.Logs(5, fmt.Sprintf("Unable to retrieve data from database: %s", err.Error()))
		return false, err
	}

	// check if both hashes match
	checkName := encrypt.VerifyHash(hashName, dbHashName)
	checkEmail := encrypt.VerifyHash(hashEmail, dbHashEmail)
	checkIncidentDate := encrypt.VerifyHash(hashIncidentDate, dbHashIncidentDate)
	checkFacilityName := encrypt.VerifyHash(hashFacilityName, dbHashFacilityName)

	// if any of the checks fail, return false
	if !checkName {
		logs.Logs(3, "Name does not match")
		return false, nil
	}

	if !checkEmail {
		logs.Logs(3, "Email does not match")
		return false, nil
	}

	if !checkIncidentDate {
		logs.Logs(3, "Incident date does not match")
		return false, nil
	}

	if !checkFacilityName {
		logs.Logs(3, "Facility name does not match")
		return false, nil
	}

	// if all checks pass, return true
	return true, nil
}

/*
DearMatronFullReport retrieves a report from the database given the parameters.

All parameters are required.
The function hashes the parameters and creates a query with the encrypted params.
The function then queries the database and returns the report.
*/
func DearMatronFullReport(name, email, incidentDate, facilityType, facilityName, country string) (GetDearMatronFullReport, error) {
	// validate input
	if name == "" || email == "" || incidentDate == "" || facilityName == "" {
		return GetDearMatronFullReport{}, errors.New("all parameters are required")
	}

	// hash params
	hashName := encrypt.HashData(name)
	hashEmail := encrypt.HashData(email)
	hashIncidentDate := encrypt.HashData(incidentDate)
	hashFacilityName := encrypt.HashData(facilityName)

	// create a query with encrypted params
	query := `
			SELECT name, email, phone_number, incident_date, facility_type, facility_name, incident_location, severity, affiliation, incident_description, make_claim, make_public, country
			FROM tbl_dearmatron
			WHERE hash_name = $1
			AND hash_email = $2
			AND hash_incident_date = $3
		AND facility_type = $4
		AND hash_facility_name = $5
		AND country = $6;`
	if db == nil {
		logs.Logs(5, "Database connection is not initialized")
		return GetDearMatronFullReport{}, errors.New("database connection is not initialized")
	}

	rows, err := db.Query(
		query,
		hashName,
		hashEmail,
		hashIncidentDate,
		facilityType,
		hashFacilityName,
		country,
	)
	if err != nil {
		logs.Logs(5, fmt.Sprintf("Unable to retrieve data from database: %s", err.Error()))
		return GetDearMatronFullReport{}, err
	}
	defer rows.Close()

	var report GetDearMatronFullReport

	for rows.Next() {
		err := rows.Scan(
			&report.Name,
			&report.Email,
			&report.PhoneNumber,
			&report.IncidentDate,
			&report.FacilityType,
			&report.FacilityName,
			&report.IncidentLocation,
			&report.Severity,
			&report.Affiliation,
			&report.IncidentDescription,
			&report.MakeClaim,
			&report.MakePublic,
			&report.Country,
		)
		if err != nil {
			logs.Logs(5, fmt.Sprintf("Unable to scan row: %s", err.Error()))
			return GetDearMatronFullReport{}, err
		}
	}

	err = rows.Err()
	if err != nil {
		logs.Logs(5, fmt.Sprintf("Unable to retrieve data from database: %s", err.Error()))
		return GetDearMatronFullReport{}, err
	}
	return report, nil
}

/*
DeleteGDPRData deletes a record from the 'tbl_dearmatron' database table based on the given encrypted data.

Parameters:

- name: name of the user.

- email: email of the user.

- incidentDate: date of the incident.

- facilityType: Type of the facility.

- facilityName: name of the facility.

Returns:

- An error if the deletion fails, otherwise nil.
*/
func DeleteGDPRData(name, email, incidentDate, facilityType, facilityName, country string) error {
	if db == nil {
		logs.Logs(5, "Database connection is not initialized")
		return errors.New("database connection is not initialized")
	}

	// hash params
	hashName := encrypt.HashData(name)
	hashEmail := encrypt.HashData(email)
	hashIncidentDate := encrypt.HashData(incidentDate)
	hashFacilityName := encrypt.HashData(facilityName)

	query := `
	DELETE
	FROM tbl_dearmatron
		WHERE hash_name = $1
		AND hash_email = $2
		AND hash_incident_date = $3
		AND facility_type = $4
		AND hash_facility_name = $5
		AND country = $6;`

	rows, err := db.Query(
		query,
		hashName,
		hashEmail,
		hashIncidentDate,
		facilityType,
		hashFacilityName,
		country,
	)

	if err != nil {
		logs.Logs(5, fmt.Sprintf("Unable to delete data from database: %s", err.Error()))
		return err
	}
	defer rows.Close()

	return nil
}
