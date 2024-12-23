package utils

import (
	"fmt"
	"groupie-tracker/models"
	"sort"
	"strings"
	"time"
)

func MapAllRelations(relations models.AllRelations) []map[string][]string {
	var allRelations []map[string][]string
	for _, rel := range relations.Index {
		newRel := mapDatesLocations(rel.DatesLocations)
		allRelations = append(allRelations, newRel)
	}
	return allRelations
}

func mapDatesLocations(datesLocations map[string][]string) map[string][]string {
	newDatesLocations := make(map[string][]string)
	for k, v := range datesLocations {
		newKey := strings.ReplaceAll(k, "_", " ")
		newKey = strings.ReplaceAll(newKey, "-", ", ")
		newKey = strings.ReplaceAll(newKey, newKey, capitalizeWords(newKey))
		newKey = strings.ReplaceAll(newKey, "Uk", "UK")
		newKey = strings.ReplaceAll(newKey, "Usa", "USA")
		newDatesLocations[newKey] = v
	}
	return newDatesLocations
}

func GetAllLocations(locations models.AllLocations) [][]string {
	var allLocations [][]string
	for _, loc := range locations.Index {
		newLoc := locationsFormat(loc.Locations)
		allLocations = append(allLocations, newLoc)
	}

	return allLocations
}

func locationsFormat(locations []string) []string {
	for i, location := range locations {
		locations[i] = strings.ReplaceAll(location, "_", " ")
		locations[i] = strings.ReplaceAll(locations[i], "-", ", ")
		locations[i] = strings.ReplaceAll(locations[i], locations[i], capitalizeWords(locations[i]))
		locations[i] = strings.ReplaceAll(locations[i], "Uk", "UK")
		locations[i] = strings.ReplaceAll(locations[i], "Usa", "USA")
		locations[i] = capitalizeWords(locations[i])
	}
	return locations
}

func GetAllDates(dates models.AllDates) [][]string {
	var allDates [][]string
	for _, date := range dates.Index {
		newDate := datesFormat(date.Dates)
		allDates = append(allDates, newDate)
	}

	return allDates
}

func datesFormat(dates []string) []string {
	layout := "02-01-2006"
	var timeSlice []time.Time
	for _, date := range dates {
		if strings.HasPrefix(date, "*") {
			date = date[1:]
		}
		t, err := time.Parse(layout, date)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			continue
		}
		timeSlice = append(timeSlice, t)

	}
	// Sort the slice of time.Time
	sort.Slice(timeSlice, func(i, j int) bool {
		return timeSlice[i].Before(timeSlice[j])
	})
	// Convert the sorted slice of time.Time back to a slice of strings
	for i, t := range timeSlice {
		dates[i] = t.Format(layout)
	}

	return dates
}

func capitalizeWords(sentence string) string {
	words := strings.Fields(sentence)
	for i, word := range words {
		words[i] = strings.ToUpper(string(word[0])) + word[1:]
	}
	return strings.Join(words, " ")
}
