// Fonction pour ajouter un favori : 
package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"groupie-tracker/models"
	"groupie-tracker/services"
)

func Favorites(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("favorites")
	var favoriteIDs []int

	if err == nil {
		idsStr := strings.Split(cookie.Value, ",")
		for _, idStr := range idsStr {
			if idStr != "" {
				if id, err := strconv.Atoi(idStr); err == nil {
					favoriteIDs = append(favoriteIDs, id)
				}
			}
		}
	}

	allArtists, err := services.GetArtists()
	if err != nil {
		log.Println("Erreur lors de la récupération des artistes:", err)
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
		return
	}

	var favoriteArtists []models.Artist
	for _, artist := range allArtists {
		for _, favID := range favoriteIDs {
			if artist.ID == favID {
				favoriteArtists = append(favoriteArtists, artist)
				break
			}
		}
	}

	tmpl, err := template.ParseFiles(filepath.Join("templates", "favoris.html"))
	if err != nil {
		log.Println("Erreur template favoris.html:", err)
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
		return
	}

	data := struct {
		Artists []models.Artist
		Count   int
	}{
		Artists: favoriteArtists,
		Count:   len(favoriteArtists),
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println("Erreur exécution template favoris.html:", err)
	}
}