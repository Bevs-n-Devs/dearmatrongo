package handlers

import (
	"log"
	"net/http"
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

	// Start the server
	log.Println("Starting application...")
	log.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to load:", err.Error())
	}
}
