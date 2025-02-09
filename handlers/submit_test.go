package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/Bevs-n-Devs/dearmatrongo/logs"
)

func TestSubmitReport(t *testing.T) {
	go logs.ProcessLogs()
	go StartHTTPServer()
	mockData := []struct {
		testName            string
		method              string
		name                string
		email               string
		phone               string
		date                string
		facilityType        string
		facilityName        string
		incidentLocation    string
		severity            string
		affiliation         string
		incidentDescription string
		makeClaim           string
		expectStatusCode    int
	}{
		{
			testName:            "Testing valid post request",
			method:              http.MethodPost,
			name:                "John Doe",
			email:               "jdoe@email.com",
			phone:               "1234567890",
			date:                "2024-10-26",
			facilityType:        "Hospital",
			facilityName:        "St. Geroges",
			incidentLocation:    "At Home",
			severity:            "High",
			affiliation:         "Family Member",
			incidentDescription: "Something random",
			makeClaim:           "No",
			expectStatusCode:    http.StatusSeeOther,
		},
		{
			testName:         "Testing invalid method",
			method:           http.MethodGet,
			expectStatusCode: http.StatusSeeOther,
		},
	}

	for _, test := range mockData {
		t.Run(test.testName, func(t *testing.T) {
			form := url.Values{}
			if test.method == http.MethodPost {
				form.Set("name", test.name)
				form.Set("email", test.email)
				form.Set("phone", test.phone)
				form.Set("incident_date", test.date)
				form.Set("facility_type", test.facilityType)
				form.Set("facility_name", test.facilityName)
				form.Set("incident_location", test.incidentLocation)
				form.Set("severity", test.severity)
				form.Set("affiliation", test.affiliation)
				form.Set("incident_description", test.incidentDescription)
				form.Set("make_claim", test.makeClaim)
			}

			request := httptest.NewRequest(test.method, "/submit", bytes.NewBufferString(form.Encode()))
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			response := httptest.NewRecorder()
			SubmitReport(response, request)

			result := response.Result()
			defer result.Body.Close()

			if result.StatusCode != test.expectStatusCode {
				t.Errorf("Expected status code %d, got %d", test.expectStatusCode, result.StatusCode)
			}
		})
	}
}
