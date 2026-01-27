// Utiliser l'API externe et récupérer les données :
package services

import (
	"encoding/json"
	"net/http"
	"strconv"

	"groupie-tracker/models"
)

func GetArtists() ([]models.Artist, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var artists []models.Artist
	err = json.NewDecoder(resp.Body).Decode(&artists)
	return artists, err
}

func GetRelation(id int) (*models.Relation, error) {
	url := "https://groupietrackers.herokuapp.com/api/relation/" + strconv.Itoa(id)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var relation models.Relation
	err = json.NewDecoder(resp.Body).Decode(&relation)
	return &relation, err
}
