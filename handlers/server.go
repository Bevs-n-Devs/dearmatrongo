package handlers

import (
	"log"
	"net/http"
)

func StartServer() {
	// static file server for assets like css, js, images
	var staticFiles = http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", staticFiles))

	// define routes
	http.HandleFunc("/", HomePage)
	// http.HandleFunc("/about", AboutPage)

	// start the server
	log.Println("Starting application...")
	log.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to load:", err.Error())
	}
}
