package handlers

import (
	"net/http"

	"github.com/Bevs-n-Devs/dearmatrongo/logs"
)

func StartServer() {
	// Initialize templates
	InitTemplates()

	// Static file server for assets like CSS, JS, images
	var staticFiles = http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", staticFiles))

	// Define routes
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/report", ReportPage)
	http.HandleFunc("/submit", SubmitReport)

	// Start the server
	logs.Log(logs.INFO, "Starting server...")
	logs.Log(logs.INFO, "Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logs.Log(logs.ERROR, "Server failed to load: "+err.Error())
	}
}
