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

type GetDearMatronFullReport struct {
	Name                []byte `json:"name"`
	Email               []byte `json:"email"`
	PhoneNumber         []byte `json:"phone_number"`
	IncidentDate        []byte `json:"incident_date"`
	FacilityType        string `json:"facility_type"`
	FacilityName        []byte `json:"facility_name"`
	IncidentLocation    []byte `json:"incident_location"`
	Severity            []byte `json:"severity"`
	Affiliation         []byte `json:"affiliation"`
	IncidentDescription []byte `json:"incident_description"`
	MakeClaim           []byte `json:"make_claim"`
	MakePublic          string `json:"make_public"`
	Country             string `json:"country"`
}

type ShowDearMatronFullReport struct {
	Name                string `json:"name"`
	Email               string `json:"email"`
	PhoneNumber         string `json:"phone_number"`
	IncidentDate        string `json:"incident_date"`
	FacilityType        string `json:"facility_type"`
	FacilityName        string `json:"facility_name"`
	IncidentLocation    string `json:"incident_location"`
	Severity            string `json:"severity"`
	Affiliation         string `json:"affiliation"`
	IncidentDescription string `json:"incident_description"`
	MakeClaim           string `json:"make_claim"`
	MakePublic          string `json:"make_public"`
	Country             string `json:"country"`
}
