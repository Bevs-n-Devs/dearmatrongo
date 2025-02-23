package database

import "database/sql"

var (
	db *sql.DB // global db variable to hold db connection
)

type GetDearMatronReport struct {
	FacilityType        string `json:"facility_type"`
	FacilityName        []byte `json:"facility_name"`
	IncidentDate        []byte `json:"incident_date"`
	IncidentLocation    []byte `json:"incident_location"`
	IncidentDescription []byte `json:"incident_description"`
}

type ShowDearMatronReport struct {
	FacilityType        string `json:"facility_type"`
	FacilityName        string `json:"facility_name"`
	IncidentDate        string `json:"incident_date"`
	IncidentLocation    string `json:"incident_location"`
	IncidentDescription string `json:"incident_description"`
}
