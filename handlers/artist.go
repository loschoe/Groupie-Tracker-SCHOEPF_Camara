// Structures & variables pour créer l'artiste : 
package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"groupie-tracker/services"
	"groupie-tracker/utils"
)

type ArtistPageData struct {
	Artist   interface{}
	Relation interface{}
	ByLoc    interface{}
}

func Artist(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/artist/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	artists, _ := services.GetArtists()

	var artist interface{}
	for _, a := range artists {
		if a.ID == id {
			artist = a
			break
		}
	}
	if artist == nil {
		http.NotFound(w, r)
		return
	}

	relation, err := services.GetRelation(id)
	if err != nil {
		http.Error(w, "Erreur concerts", http.StatusInternalServerError)
		return
	}

	data := ArtistPageData{
		Artist:   artist,
		Relation: relation,
		ByLoc:    utils.GroupByLocation(relation),
	}

	tmpl := template.Must(
		template.New("artist.html").
			Funcs(template.FuncMap{"join": strings.Join}).
			ParseFiles(filepath.Join("templates", "artist.html")),
	)

	tmpl.ExecuteTemplate(w, "artist.html", data)
}