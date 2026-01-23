package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strings"

	"groupie-tracker/models"
	"groupie-tracker/services"
	"groupie-tracker/utils"
)

type CompareData struct {
	LeftQuery   string
	LeftArtist  *models.Artist
	LeftByLoc   interface{}

	RightQuery  string
	RightArtist *models.Artist
	RightByLoc  interface{}

	NoResult bool
}

func Compare(w http.ResponseWriter, r *http.Request) {
	leftQuery := r.URL.Query().Get("left")
	rightQuery := r.URL.Query().Get("right")

	data := CompareData{
		LeftQuery:  leftQuery,
		RightQuery: rightQuery,
	}

	if leftQuery != "" {
		artist, relation := searchArtist(leftQuery)
		if artist != nil {
			data.LeftArtist = artist
			if relation != nil {
				data.LeftByLoc = utils.GroupByLocation(relation)
			}
		}
	}

	if rightQuery != "" {
		artist, relation := searchArtist(rightQuery)
		if artist != nil {
			data.RightArtist = artist
			if relation != nil {
				data.RightByLoc = utils.GroupByLocation(relation)
			}
		}
	}

	if (leftQuery != "" && data.LeftArtist == nil) ||
		(rightQuery != "" && data.RightArtist == nil) {
		data.NoResult = true
	}

	funcMap := template.FuncMap{
		"join": strings.Join,
	}

	tmpl := template.Must(template.New("Compare.html").Funcs(funcMap).ParseFiles(filepath.Join("templates", "Compare.html")))
	tmpl.Execute(w, data)
}

func searchArtist(query string) (*models.Artist, *models.Relation) {
	artists, err := services.GetArtists()
	if err != nil {
		return nil, nil
	}

	query = strings.ToLower(strings.TrimSpace(query))

	for _, artist := range artists {
		if strings.ToLower(artist.Name) == query {
			relation, err := services.GetRelation(artist.ID)
			if err != nil {
				return &artist, nil
			}
			return &artist, relation
		}
	}

	return nil, nil
}