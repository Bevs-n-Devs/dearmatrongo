package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Bevs-n-Devs/dearmatrongo/encrypt"
	"github.com/Bevs-n-Devs/dearmatrongo/logs"
)

func StartHTTPServer() {
	logs.Logs(logInfo, "Starting Dear Matron app...")
	// Initialize templates
	InitTemplates()
	err := encrypt.InitEncryption()
	if err != nil {
		logs.Logs(logError, fmt.Sprintf("Error initializing encryption: %s", err.Error()))
	}

	// Static file server for assets like CSS, JS, images
	var staticFiles = http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", staticFiles))

	// Define routes
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/uk/report", ReportUK)
	http.HandleFunc("/uk/submit", SubmitUK)
	http.HandleFunc("/uk/view", ViewUK)
	http.HandleFunc("/usa", HomeUSA)
	http.HandleFunc("/usa/report", ReportUSA)
	http.HandleFunc("/usa/submit", SubmitUSA)
	http.HandleFunc("/usa/view", ViewUSA)
	http.HandleFunc("/uk/data-policy", GDPR)
	// http.HandleFunc("/uk/data-policy-delete", DeleteGDPRData)
	http.HandleFunc("/usa/data-policy", CCPA)
	// http.HandleFunc("/usa/data-policy-delete", DeleteCCPAData)

	// initialize port
	httpServerPort := os.Getenv("PORT")
	// start server on local machine
	if httpServerPort == "" {
		logs.Logs(logWarning, "Could not get PORT from Heroku. Starting server on default port http://localhost:9000...")
		httpServerPort = "9000"
		err := http.ListenAndServe(fmt.Sprintf(":%s", httpServerPort), nil)
		if err != nil {
			logs.Logs(logError, fmt.Sprintf("HTTP server failed to start: %s", err.Error()))
		}
	}

	// Start the server on Heroku port
	logs.Logs(logInfo, fmt.Sprintf("Server running on :%s", httpServerPort))
	err = http.ListenAndServe(fmt.Sprintf(":%s", httpServerPort), nil)
	if err != nil {
		logs.Logs(logError, fmt.Sprintf("HTTP server failed to start: %s", err.Error()))
	}
}
