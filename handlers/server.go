package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Bevs-n-Devs/dearmatrongo/logs"
)

func StartHTTPServer() {
	logs.Logs(1, "Starting Dear Matron app...")
	// Initialize templates
	InitTemplates()

	// Static file server for assets like CSS, JS, images
	var staticFiles = http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", staticFiles))

	// Define routes
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/report", ReportPage)
	http.HandleFunc("/submit", SubmitReport)
	http.HandleFunc("/getReports", GetReports)

	// initialize port
	httpServerPort := os.Getenv("PORT")
	// start server on local machine
	if httpServerPort == "" {
		logs.Logs(2, "Could not get PORT from Heroku. Starting server on default port http://localhost:9000...")
		httpServerPort = "9000"
		err := http.ListenAndServe(fmt.Sprintf(":%s", httpServerPort), nil)
		if err != nil {
			logs.Logs(3, fmt.Sprintf("HTTP server failed to start: %s", err.Error()))
		}
	}

	// Start the server on Heroku port
	logs.Logs(1, fmt.Sprintf("Server running on :%s", httpServerPort))
	err := http.ListenAndServe(fmt.Sprintf(":%s", httpServerPort), nil)
	if err != nil {
		logs.Logs(3, fmt.Sprintf("HTTP server failed to start: %s", err.Error()))
	}
}
