package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
}

func main() {
	// Set application config
	var app application

	// Read from cmd line

	// Connect to DB
	app.Domain = "example.com"

	log.Println("Starting application on port", port)

	http.HandleFunc("/", Hello)

	// Start a web server
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
