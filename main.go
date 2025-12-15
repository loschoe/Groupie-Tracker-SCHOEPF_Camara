package main
import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"encoding/json"
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


// =================== Handlers ===================
func handleHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(filepath.Join("templates", "home.html")))
	tmpl.Execute(w, nil)
}

func artistsHandler(w http.ResponseWriter, r *http.Request) {
	artists, err := getArtists()
	if err != nil {
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
		return
	}
	tmpl, err := template.ParseFiles("templates/artists.html")
	if err != nil {
		http.Error(w, "Erreur template", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(w, artists)
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes principales
	http.HandleFunc("/", artistsHandler)
	
	log.Println("✅ Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
