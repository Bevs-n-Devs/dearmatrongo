package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Bevs-n-Devs/dearmatrongo/logs"
)

func StartHTTPServer() {
	// initialize port
	httpServerPort := os.Getenv("PORT")
	if httpServerPort == "" {
		httpServerPort = "8080"
	}

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

	// Start the server
	logs.Logs(1, "Starting HTTP server...")
	logs.Logs(1, fmt.Sprintf("Server running on :%s", httpServerPort))
	err := http.ListenAndServe(fmt.Sprintf(":%s", httpServerPort), nil)
	if err != nil {
		logs.Logs(3, fmt.Sprintf("HTTP server failed to start: %s", err.Error()))
	}
}
