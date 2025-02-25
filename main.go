package main

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/data"
	"groupie-tracker/handlers"
	"groupie-tracker/models"
	"groupie-tracker/utils"
	"net/http"
	"os"
	"strconv"
)

func main() {
	var artists []models.Artist
	fetchedArtists, relations, locations, dates, err := data.InitializeData()
	if err != nil {
		fmt.Println("Failed to fetch artists:", err)
	}
	artists = fetchedArtists

	allRelations := utils.MapAllRelations(relations)
	allLocations := utils.GetAllLocations(locations)
	allDates := utils.GetAllDates(dates)

	http.HandleFunc("/api/artists", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(artists); err != nil {
			handlers.ErrorHandler(w, r, http.StatusInternalServerError, "Failed to encode artists data.")
		}
	})

	http.HandleFunc("/api/relations", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(allRelations); err != nil {
			handlers.ErrorHandler(w, r, http.StatusInternalServerError, "Failed to encode artists data.")
		}
	})

	http.HandleFunc("/api/locations", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(allLocations); err != nil {
			handlers.ErrorHandler(w, r, http.StatusInternalServerError, "Failed to encode artists data.")
		}
	})

	http.HandleFunc("/api/dates", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(allDates); err != nil {
			handlers.ErrorHandler(w, r, http.StatusInternalServerError, "Failed to encode artists data.")
		}
	})

	http.HandleFunc("/suggestions", handlers.SuggestionsHandler)

	http.HandleFunc("/search", handlers.SearchHandler)

	http.HandleFunc("GET /map", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/map.html")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HomeHandler(w, r, artists)
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	fmt.Printf("Server is running on port http://localhost:%s ...\n", port)
	err = http.ListenAndServe(":"+port, nil)
	for err != nil {
		portInt, _ := strconv.Atoi(port)
		portInt++
		port = strconv.Itoa(portInt)
		fmt.Printf("\033[A\rServer is running on port http://localhost:%s ...\n", port)
		err = http.ListenAndServe(":"+port, nil)
	}
}
