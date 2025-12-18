package main
import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"encoding/json"
	"strings"
)

// =========== Fonctions & Variables =============
type Artist struct {
	ID				int			`json:"id"`
	Name			string 		`json:"name"`
	Members 		[]string 	`json:"members"`
	CreationDate 	int 		`json:"creationDate"`
	FirstAlbum 		string  	`json:"firstAlbum"`
	Image        	string   	`json:"image"`
}

func (a Artist) NbMembers() int {
    return len(a.Members)
}

type PageData struct {
	Query 			string
	Searched 		bool
	Artists 		[]Artist
}

func getArtists() ([]Artist, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var artists []Artist
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		return nil, err
	}

	return artists, nil
}

func filterArtists(artists []Artist, query string) []Artist {
	query = strings.ToLower(query)
	var results []Artist

	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), query) {
			results = append(results, artist)
		}
	}

	return results
}

// =================== Handlers ===================
func handleHome(w http.ResponseWriter, r *http.Request) {
	allArtists, err := getArtists()
	if err != nil {
		log.Println("Erreur getArtists:", err)
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
		return
	}

	data := PageData{
		Artists: allArtists, // par défaut : tous les artistes
	}

	if r.Method == http.MethodPost {
		query := strings.TrimSpace(r.FormValue("group"))

		if query != "" {
			// Champ rempli → filtre
			data.Query = query
			data.Searched = true
			data.Artists = filterArtists(allArtists, query)
		} else {
			// Champ vide → retourner à tous les artistes
			data.Query = ""
			data.Searched = false
			data.Artists = allArtists
		}
	}

	tmpl := template.Must(template.ParseFiles(filepath.Join("templates", "home.html")))
	if err := tmpl.Execute(w, data); err != nil {
		log.Println("Erreur template home.html :", err)
	}
}

func handleArtist(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(filepath.Join("templates", "artist.html")))
	tmpl.Execute(w, nil)
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(filepath.Join("templates", "AboutUS.html")))
	tmpl.Execute(w, nil)
}

// =================== Fonction main ===================
func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	
	// Les pages !
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/about", handleAbout)
	http.HandleFunc("/artist", handleArtist)
	
	// Démarrage du serveur :
	log.Println("✅ Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
