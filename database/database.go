package database

import (
	"database/sql"
	"encoding/hex"
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

// insert into dear_matron table
/*
eg.
INSERT INTO public.dear_matron(
	name, email, phone_number, incident_date, facility_type, facility_name, location, severity, affiliation, description, make_claim, make_public, country, submitted)
	VALUES ('john doe', 'jdoe"email.com', '1234567890', '2024-10-26', 'clinic', 'st geroges', 'at home', 'high', 'family member', 'something random', 'Yes', 'Yes', 'UK', NOW());
*/
func InsertDearMatron(name, email, phoneNumber, incidentDate []byte, facilityType string, facilityName, location, severity, affiliation, description, makeClaim []byte, makePublic, country string) error {
	logs.Logs(4, "Creating new report for Dear Matron...")
	if db == nil {
		logs.Logs(5, "Database connection is not initialized")
		return errors.New("database connection is not initialized")
	}
	// SQL query
	query := `
	INSERT INTO dearmatron (name, email, phone_number, incident_date, facility_type, facility_name, incident_location, severity, affiliation, incident_description, make_claim,  make_public, country, submitted)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, NOW());
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
	SELECT facility_type, facility_name, incident_date, incident_location, incident_description
	FROM dearmatron
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
Gets all USA reports from dear_matron table

Only for users who are in the USA that want their reports to be public.

Returns a slice of GetDearMatronReport structs.
*/
func GetAllReportsUSA() ([]GetDearMatronReport, error) {
	query := `
	SELECT facility_type, facility_name, incident_date, incident_location, incident_description
	FROM dearmatron
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
CheckGDPRData checks if a record exists in the 'dearmatron' database table
that matches the provided name, email, incident date, and facility name.

All parameters are encrypted before querying the database.

Parameters:

- name: The unencrypted name of the user.

- email: The unencrypted email of the user.

- incidentDate: The unencrypted date of the incident.

- facilityName: The unencrypted name of the facility.

Returns:

- A boolean indicating whether a matching record is found (true) or not (false).
*/
func CheckGDPRData(name, email, incidentDate, facilityName string) (bool, error) {
	// validate input
	if name == "" || email == "" || incidentDate == "" || facilityName == "" {
		return false, errors.New("all parameters are required")
	}

	// encrypt params
	encryptName, err := encrypt.Encrypt([]byte(name))
	if err != nil {
		logs.Logs(5, fmt.Sprintf("Unable to encrypt name: %s", err.Error()))
		return false, err
	}
	encryptEmail, err := encrypt.Encrypt([]byte(email))
	if err != nil {
		logs.Logs(5, fmt.Sprintf("Unable to encrypt email: %s", err.Error()))
		return false, err
	}
	encryptIncidentDate, err := encrypt.Encrypt([]byte(incidentDate))
	if err != nil {
		logs.Logs(5, fmt.Sprintf("Unable to encrypt incident date: %s", err.Error()))
		return false, err
	}
	encryptFacilityName, err := encrypt.Encrypt([]byte(facilityName))
	if err != nil {
		logs.Logs(5, fmt.Sprintf("Unable to encrypt facility name: %s", err.Error()))
		return false, err
	}

	// create a query with encrypted params
	query := `
		SELECT 1 FROM dearmatron
		WHERE name = decode($1, 'hex')
		AND email = decode($2, 'hex')
		AND incident_date = decode($3, 'hex')
		AND facility_name = decode($4, 'hex');
	`

	// execute query to find matching data within database
	var exists int
	err = db.QueryRow(
		query,
		hex.EncodeToString(encryptName),
		hex.EncodeToString(encryptEmail),
		hex.EncodeToString(encryptIncidentDate),
		hex.EncodeToString(encryptFacilityName),
	).Scan(&exists)

	// if found, return true; false otherwise
	if err == sql.ErrNoRows {
		logs.Logs(5, "GDPR search could not find data in database")
		return false, nil
	}

	if err != nil {
		logs.Logs(5, fmt.Sprintf("Unable to retrieve data from database: %s", err.Error()))
		return false, err
	}

	logs.Logs(5, "GDPR search found data in database")
	return true, nil
}

/*
Returns a single report from the 'dearmatron' database table, given the encrypted name, email, incident date and facility name.

Parameters:

- name: Encrypted name of the user.

- email: Encrypted email of the user.

- incidentDate: Encrypted date of the incident.

- facilityName: Encrypted name of the facility.

Returns:

- The struct of the report if a matching record is found.

- An empty struct and an error if no matching record is found or if there is an error in querying the database.
*/
func GDPRDearMatronFullReport(name, email, incidentDate, facilityType string) (GetDearMatronFullReport, error) {
	// convert strings to encrypted bytes
	encrytName, err := encrypt.Encrypt([]byte(name))
	if err != nil {
		logs.Logs(5, fmt.Sprintf("Unable to encrypt name: %s", err.Error()))
		return GetDearMatronFullReport{}, err
	}
	encrytEmail, err := encrypt.Encrypt([]byte(email))
	if err != nil {
		logs.Logs(5, fmt.Sprintf("Unable to encrypt email: %s", err.Error()))
		return GetDearMatronFullReport{}, err
	}
	encrytIncidentDate, err := encrypt.Encrypt([]byte(incidentDate))
	if err != nil {
		logs.Logs(5, fmt.Sprintf("Unable to encrypt incident date: %s", err.Error()))
		return GetDearMatronFullReport{}, err
	}
	encrytFacilityType, err := encrypt.Encrypt([]byte(facilityType))
	if err != nil {
		logs.Logs(5, fmt.Sprintf("Unable to encrypt facility name: %s", err.Error()))
		return GetDearMatronFullReport{}, err
	}

	query := `
	SELECT * FROM dearmatron
	WHERE name = decode($1, 'hex')
    AND email = decode($2, 'hex')
    AND incident_date = decode($3, 'hex')
    AND facility_type = decode($4, 'hex');`
	if db == nil {
		logs.Logs(5, "Database connection is not initialized")
		return GetDearMatronFullReport{}, errors.New("database connection is not initialized")
	}

	rows, err := db.Query(
		query,
		hex.EncodeToString(encrytName),
		hex.EncodeToString(encrytEmail),
		hex.EncodeToString(encrytIncidentDate),
		hex.EncodeToString(encrytFacilityType),
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

// DeleteGDPRData deletes a record from the 'dearmatron' database table based on the given encrypted data.
//
// Parameters:
// - name: Encrypted name of the user.
// - email: Encrypted email of the user.
// - incidentDate: Encrypted date of the incident.
// - facilityName: Encrypted name of the facility.
//
// Returns:
// - An error if the deletion fails, otherwise nil.

func DeleteGDPRData(name, email, incidentDate, facilityType string) error {
	// convert strings to encrypted bytes
	encryptName, err := encrypt.Encrypt([]byte(name))
	if err != nil {
		logs.Logs(5, fmt.Sprintf("Unable to encrypt name: %s", err.Error()))
		return err
	}
	encryptEmail, err := encrypt.Encrypt([]byte(email))
	if err != nil {
		logs.Logs(5, fmt.Sprintf("Unable to encrypt email: %s", err.Error()))
		return err
	}
	encryptIncidentDate, err := encrypt.Encrypt([]byte(incidentDate))
	if err != nil {
		logs.Logs(5, fmt.Sprintf("Unable to encrypt incident date: %s", err.Error()))
		return err
	}
	encryptFacilityType, err := encrypt.Encrypt([]byte(facilityType))
	if err != nil {
		logs.Logs(5, fmt.Sprintf("Unable to encrypt facility name: %s", err.Error()))
		return err
	}

	query := `
	DELETE FROM dearmatron
	WHERE name = decode($1, 'hex')
	AND email = decode($2, 'hex')
	AND incident_date = decode($3, 'hex')
	AND facility_type = decode($4, 'hex');
	`
	_, err = db.Exec(
		query,
		hex.EncodeToString(encryptName),
		hex.EncodeToString(encryptEmail),
		hex.EncodeToString(encryptIncidentDate),
		hex.EncodeToString(encryptFacilityType),
	)
	if err != nil {
		logs.Logs(5, fmt.Sprintf("Unable to delete data from database: %s", err.Error()))
		return err
	}
	return nil
}
