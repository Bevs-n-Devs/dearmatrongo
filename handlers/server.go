package handlers

import (
	"fmt"
	"net"
	"net/http"

	"github.com/Bevs-n-Devs/dearmatrongo/logs"
)

const (
	httpServerPort = ":9000"
	tcpServerPort  = ":8081"
)

func StartHTTPServer() {
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
	logs.Logs(1, fmt.Sprintf("Server running on http://localhost%s", httpServerPort))
	err := http.ListenAndServe(httpServerPort, nil)
	if err != nil {
		logs.Logs(3, fmt.Sprintf("HTTP server failed to start: %s", err.Error()))
	}
}

func StartTCPServer() {
	logs.Logs(1, fmt.Sprintf("Starting TCP server on port %s", tcpServerPort))
	listen, err := net.Listen("tcp", tcpServerPort)
	if err != nil {
		logs.Logs(3, fmt.Sprintf("TCP server failed to start: %s", err.Error()))
		return
	}
	defer listen.Close()

	for {
		var buff = make([]byte, 1024) // listen to incoming connections
		conn, err := listen.Accept()
		if err != nil {
			logs.Logs(3, fmt.Sprintf("TCP server failed to accept connection: %s", err.Error()))
			continue
		}
		n, err := conn.Read(buff)
		if err != nil {
			logs.Logs(3, fmt.Sprintf("TCP server failed to read data: %s", err.Error()))
			continue
		}

		logs.Logs(1, fmt.Sprintf("TCP server received data: %s", string(buff[:n])))

	}
}
