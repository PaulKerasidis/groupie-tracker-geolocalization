package data

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/models"
	"net/http"
	"time"
)

const baseURL = "https://groupietrackers.herokuapp.com/api"

// generic fetch function to reduce code duplication
func fetch(endpoint string, target interface{}) error {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(baseURL + endpoint)
	if err != nil {
		return fmt.Errorf("failed to get %s: %w", endpoint, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(target); err != nil {
		return fmt.Errorf("failed to decode %s: %w", endpoint, err)
	}

	return nil
}

func GetDates() (models.AllDates, error) {
	var dates models.AllDates
	if err := fetch("/dates", &dates); err != nil {
		return models.AllDates{}, err
	}
	return dates, nil
}

func GetLocations() (models.AllLocations, error) {
	var locations models.AllLocations
	if err := fetch("/locations", &locations); err != nil {
		return models.AllLocations{}, err
	}
	return locations, nil
}

func GetRelations() (models.AllRelations, error) {
	var relations models.AllRelations
	if err := fetch("/relation", &relations); err != nil {
		return models.AllRelations{}, err
	}
	return relations, nil
}

func GetArtists() ([]models.Artist, error) {
	var artists []models.Artist
	if err := fetch("/artists", &artists); err != nil {
		return nil, err
	}
	return artists, nil
}

func InitializeData() ([]models.Artist, models.AllRelations, models.AllLocations, models.AllDates, error) {
	var (
		fetchedArtists []models.Artist
		relations      models.AllRelations
		dates          models.AllDates
		locations      models.AllLocations
	)

	// Use channels for concurrent fetching
	errChan := make(chan error, 4)

	go func() {
		var err error
		fetchedArtists, err = GetArtists()
		errChan <- err
	}()

	go func() {
		var err error
		relations, err = GetRelations()
		errChan <- err
	}()

	go func() {
		var err error
		dates, err = GetDates()
		errChan <- err
	}()

	go func() {
		var err error
		locations, err = GetLocations()
		errChan <- err
	}()

	// Check for errors
	for i := 0; i < 4; i++ {
		if err := <-errChan; err != nil {
			return nil, relations, locations, dates, err
		}
	}

	return fetchedArtists, relations, locations, dates, nil
}

