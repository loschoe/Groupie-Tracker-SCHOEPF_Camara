// Page d'accueil du siite 
package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"sort"
	"strings"

	"groupie-tracker/models"
	"groupie-tracker/services"
)

type PageData struct {
	Query         string
	Searched      bool
	Artists       []models.Artist
	NoResult      bool
	AlphaOrder    bool
	PeriodeSelect string
	MembresValue  string
}

func filterArtists(artists []models.Artist, query string) []models.Artist {
	query = strings.ToLower(query)
	var results []models.Artist

	for _, artist := range artists {
		name := strings.ToLower(artist.Name)
		creation := fmt.Sprintf("%d", artist.CreationDate)

		if strings.Contains(name, query) || strings.Contains(creation, query) {
			results = append(results, artist)
		}
	}
	return results
}

func Home(w http.ResponseWriter, r *http.Request) {
	allArtists, err := services.GetArtists()
	if err != nil {
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
		return
	}

	query := r.URL.Query().Get("group")
	alpha := r.URL.Query().Get("alpha")
	periode := r.URL.Query().Get("periode")
	membres := r.URL.Query().Get("members")

	filtered := allArtists

	if query != "" {
		filtered = filterArtists(filtered, query)
	}

	if alpha != "" || periode != "" || membres != "" {
		filtered = applyFilters(filtered, alpha, periode, membres)
	}

	data := PageData{
		Query:         query,
		Artists:       filtered,
		NoResult:      len(filtered) == 0,
		AlphaOrder:    alpha == "1",
		PeriodeSelect: periode,
		MembresValue:  membres,
	}

	tmpl := template.Must(template.ParseFiles(filepath.Join("templates", "home.html")))
	tmpl.Execute(w, data)
}

func correspondPeriode(creation int, periode string) bool {
	switch periode {
	case "1950-1970":
		return creation >= 1950 && creation <= 1970
	case "1970-1990":
		return creation >= 1970 && creation <= 1990
	case "1990-2000":
		return creation >= 1990 && creation <= 2000
	case "2000+":
		return creation >= 2000
	default:
		return true
	}
}

func applyFilters(artists []models.Artist, alpha, periode, membresStr string) []models.Artist {
	var results []models.Artist

	var membres int
	if membresStr != "" {
		fmt.Sscanf(membresStr, "%d", &membres)
	}

	for _, artist := range artists {

		if periode != "" && periode != "all" {
			if !correspondPeriode(artist.CreationDate, periode) {
				continue
			}
		}

		if membres > 0 {
			if len(artist.Members) != membres {
				continue
			}
		}

		results = append(results, artist)
	}

	if alpha == "1" {
		sort.Slice(results, func(i, j int) bool {
			return results[i].Name < results[j].Name
		})
	}

	return results
}