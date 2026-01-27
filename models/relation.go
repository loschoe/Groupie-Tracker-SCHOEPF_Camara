// Modèle relation dates - locations : 
package models

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type LocationDates struct {
	Location string
	Dates    []string
}
