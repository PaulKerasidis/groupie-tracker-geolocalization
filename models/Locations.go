package models

type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	DatesLink string   `json:"dates"`
	Dates     Dates
}