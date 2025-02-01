package database

import (
	"database/sql"
	"errors"
	"net"
	"os"
	"time"

	"github.com/Bevs-n-Devs/dearmatrongo/env"
	"github.com/Bevs-n-Devs/dearmatrongo/logs"
)

var db *sql.DB

func databaseDriver() {
	logs.Log(logs.INFO, "Uploading environment variables to database...")
	err := env.LoadEnv("env/.env")
	if err != nil {
		logs.Log(logs.ERROR, "Unable to load environment variables: "+err.Error())
	}

	var host = os.Getenv("DB_HOST")
	var port = os.Getenv("DB_PORT")
	var user = os.Getenv("DB_USER")
	var password = os.Getenv("DB_PASSWORD")
	var database = os.Getenv("DB_DATABASE")
	var sslMode = os.Getenv("DB_SSLMODE")
	var connectionString = "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + database + " sslmode=" + sslMode

	// PostgreSQL DSN (data source name)
	dsn := connectionString

	// open a raw TCP connection
	conn, err := net.Dial("tcp", dsn)
	if err != nil {
		logs.Log(logs.ERROR, "Unable to connect to database: "+err.Error())
	}
	defer conn.Close()

	// open database connection
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		logs.Log(logs.ERROR, "Unable to open database connection: "+err.Error())
	}
	defer db.Close()

	// test database connection
	err = db.Ping()
	if err != nil {
		logs.Log(logs.ERROR, "Cannot connect to database: "+err.Error())
	}
	logs.Log(logs.INFO, "Database connection successful!")
	logs.Log(logs.INFO, "Database driver initialized with connection string: "+connectionString)
}

func GetDB() *sql.DB {
	return db
}

// insert into dear_matron_db table
/*
INSERT INTO dear_matron_db (report_id, name, email, phone_number, incident_date, facility_type, facility_name, location, severity, affiliation, description, make_claim, submitted)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
*/
func InsertDearMatron(name, email, phoneNumber, incidentDate, facilityType, facilityName, location, severity, affiliation, description, makeClaim string) error {
	var current = time.Now()
	var submitted = current.Format("15:04:05 02-01-2006")
	var query = `
	INSERT INTO dear_matron_db (report_id, name, email, phone_number, incident_date, facility_type, facility_name, location, severity, affiliation, description, make_claim, submitted)
	VALUES (DEFAULT, $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);
	`
	logs.Log(logs.INFO, "Creating new report for Dear Matron...")
	if db == nil {
		logs.Log(logs.ERROR, "Database connection is not initialized")
		return errors.New("database connection is not initialized")
	}
	_, err := db.Exec(query, name, email, phoneNumber, incidentDate, facilityType, facilityName, location, severity, affiliation, description, makeClaim, submitted)
	if err != nil {
		logs.Log(logs.ERROR, "Unable to create new report for Dear Matron: "+err.Error())
	}
	logs.Log(logs.INFO, "New report created for Dear Matron successfully!")
	return nil
}

// get all data from dear_matron_db table
/*
SELECT * FROM dear_matron_db
*/
func GetAllData() (*sql.Rows, error) {
	var query = `
	SELECT * FROM dear_matron_db
	`
	logs.Log(logs.INFO, "Retrieving all data from dear_matron_db table...")
	if db == nil {
		logs.Log(logs.ERROR, "Database connection is not initialized")
		return nil, errors.New("database connection is not initialized")
	}
	rows, err := db.Query(query)
	if err != nil {
		logs.Log(logs.ERROR, "Unable to retrieve all data from dear_matron_db table: "+err.Error())
	}
	return rows, nil
}
