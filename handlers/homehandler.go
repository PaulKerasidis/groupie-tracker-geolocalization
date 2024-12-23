package handlers

import (
	"groupie-tracker/models"
	"html/template"
	"log"
	"net/http"
	"sync"
)

// PageData holds the data to be passed to the template
type PageData struct {
	Artists         []models.Artist
	NotFoundMessage string
}

var (
	homeTemplate *template.Template
	once         sync.Once // Ensures template is loaded only once
)

// initTemplate initializes the template safely
func initTemplate() error {
	var err error
	homeTemplate, err = template.ParseFiles("templates/home.html")
	return err
}

// HomeHandler handles requests to the home page
func HomeHandler(w http.ResponseWriter, r *http.Request, artists []models.Artist) {
	// Initialize template once
	once.Do(func() {
		if err := initTemplate(); err != nil {
			log.Fatalf("Error parsing home template: %v", err)
		}
	})

	// Only handle root path
	if r.URL.Path != "/" {
		log.Printf("404 - Page not found: %s", r.URL.Path)
		ErrorHandler(w, r, http.StatusNotFound, "Page not found")
		return
	}

	// Handle only GET requests
	if r.Method != http.MethodGet {
		log.Printf("405 - Method not allowed: %s", r.Method)
		ErrorHandler(w, r, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	if artists == nil {
		log.Printf("500 - Error ")
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal server error")
		return
	}

	// Render template
	data := PageData{Artists: artists}

	if err := homeTemplate.Execute(w, data); err != nil {
		log.Printf("500 - Error executing template: %v", err)
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal server error")
		return
	}
}
