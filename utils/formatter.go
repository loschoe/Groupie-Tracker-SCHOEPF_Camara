// Formattage des données pour le modele de relation:
package utils

import (
	"sort"
	"strings"

	"groupie-tracker/models"
)

func FormatLocation(raw string) string {
	parts := strings.Split(raw, "-")
	if len(parts) != 2 {
		return strings.Title(strings.ReplaceAll(raw, "_", " "))
	}

	city := strings.Title(strings.ReplaceAll(parts[0], "_", " "))
	country := strings.ToUpper(parts[1])

	return city + " - " + country
}

func FormatDate(date string) string {
	return strings.ReplaceAll(date, "-", "/")
}

func GroupByLocation(rel *models.Relation) []models.LocationDates {
	resultMap := make(map[string][]string)

	for loc, dates := range rel.DatesLocations {
		formattedLoc := FormatLocation(loc)
		for _, d := range dates {
			resultMap[formattedLoc] = append(resultMap[formattedLoc], FormatDate(d))
		}
	}

	var result []models.LocationDates
	for loc, dates := range resultMap {
		sort.Strings(dates)
		result = append(result, models.LocationDates{
			Location: loc,
			Dates:    dates,
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Location < result[j].Location
	})

	return result
}
