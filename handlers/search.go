package handlers

import (
	"encoding/json"
	"groupie-tracker/data"
	"groupie-tracker/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var artists []models.Artist
var locations models.AllLocations

func init() {
	artists, _ = data.GetArtists()
	locations, _ = data.GetLocations()
}

func SuggestionsHandler(w http.ResponseWriter, r *http.Request) {

	// Only allow GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get the search query from URL parameters
	query := strings.ToLower(r.URL.Query().Get("q"))
	if query == "" {
		json.NewEncoder(w).Encode([]interface{}{})
		return
	}

	seen := make(map[string]bool)
	type Suggestions struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	}

	sugs := []Suggestions{}
	addSuggestion := func(sType, sValue string) {
		if !seen[sValue] {
			seen[sValue] = true
			sugs = append(sugs, Suggestions{Type: sType, Value: sValue})
		}
	}

	for _, artist := range artists {

		name := strings.ToLower(artist.Name)
		if len(query) <= 2 && strings.HasPrefix(name, query) {
			addSuggestion("artist/band", artist.Name)
		} else if len(query) > 2 && strings.Contains(name, query) {
			addSuggestion("artist/band", artist.Name)
		}
		members := artist.Members
		for _, member := range members {
			if len(query) <= 2 && strings.HasPrefix(strings.ToLower(member), query) {
				addSuggestion("member", member)
			} else if len(query) > 2 && strings.Contains(strings.ToLower(member), query) {
				addSuggestion("member", member)
			}
		}

		locations := locations.Index[artist.ID-1].Locations
		for _, location := range locations {
			if len(query) <= 2 && strings.HasPrefix(strings.ToLower(location), query) {
				addSuggestion("location", location)
			} else if len(query) > 2 && strings.Contains(strings.ToLower(location), query) {
				addSuggestion("location", location)
			}
		}

		firstAlbum := strings.ToLower(artist.FirstAlbum)
		if strings.Contains(firstAlbum, query) {
			addSuggestion("first album", artist.FirstAlbum)
		}

		creationDate := artist.CreationDate
		crDate := strconv.Itoa(creationDate)
		if strings.Contains(strings.ToLower(crDate), query) {
			addSuggestion("creation date", crDate)
		}

	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Return the suggestions as JSON
	if err := json.NewEncoder(w).Encode(sugs); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

}

func SearchHandler(w http.ResponseWriter, r *http.Request) {

	homeTemplate, _ = template.ParseFiles("templates/home.html")
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	query := strings.ToLower(r.URL.Query().Get("q"))

	type Results struct {
		Artists         []models.Artist
		NotFoundMessage string
	}
	results := Results{}

	if query == "" {
		results = Results{Artists: artists}
		//json.NewEncoder(w).Encode(results)
		homeTemplate.Execute(w, results)
		return
	}

	for _, artist := range artists {
		name := strings.ToLower(artist.Name)
		members := strings.Join(artist.Members, " ")
		locations := strings.Join(locations.Index[artist.ID-1].Locations, " ")
		firstAlbum := strings.ToLower(artist.FirstAlbum)
		creationDate := strconv.Itoa(artist.CreationDate)
		if strings.Contains(name, query) ||
			strings.Contains(strings.ToLower(members), query) ||
			strings.Contains(strings.ToLower(locations), query) ||
			strings.Contains(firstAlbum, query) ||
			strings.Contains(creationDate, query) {
			results.Artists = append(results.Artists, artist)
		}
	}

	if len(results.Artists) == 0 {
		notFound := Results{NotFoundMessage: "We couldn't find any matches for your search. Please refine your query and try again."}
		//json.NewEncoder(w).Encode(notFound)
		homeTemplate.Execute(w, notFound)
		return
	}

	if err := homeTemplate.Execute(w, results); err != nil {
		log.Printf("500 - Error executing template: %v", err)
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal server error")
		return
	}

}
