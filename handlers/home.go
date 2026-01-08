package handlers

import (
	"fmt"
    "html/template"
    "log"
    "net/http"
    "path/filepath"
    "sort"
    "strings"

    "groupie-tracker/models"
    "groupie-tracker/services"
)

type PageData struct {
	Query    		string
	Searched 		bool
	Artists  		[]models.Artist
	NoResult 		bool
	AlphaOrder 		bool
	PeriodeSelect	string
	MembresValue	string
}

// Fonction pour afficher les artistes 
func filterArtists(artists []models.Artist, query string) []models.Artist {
	query = strings.ToLower(query)
	var results []models.Artist

	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), query) {
			results = append(results, artist)
		}
	}
	return results
}

// Fonction pour générer le template 
func Home(w http.ResponseWriter, r *http.Request) {
    allArtists, err := services.GetArtists()
    if err != nil {
        log.Println("❌ Erreur lors du chargement des artistes")
        http.Error(w, "Erreur serveur", http.StatusInternalServerError)
        return
    }

    data := PageData{Artists: allArtists}

    // --- Recherche POST ---
    if r.Method == http.MethodPost {
        query := strings.TrimSpace(r.FormValue("group"))
        if query != "" {
            data.Query = query
            data.Searched = true
            data.Artists = filterArtists(allArtists, query)

            if len(data.Artists) == 0 {
                data.NoResult = true
                log.Println("⚠️ Aucun résultat trouvé pour:", query)
            }
        }
    }

    // --- Filtres GET (toujours appliqués) ---
    alpha := r.URL.Query().Get("alpha")
    periode := r.URL.Query().Get("periode")
    membres := r.URL.Query().Get("members")

    data.AlphaOrder = (alpha == "1")
    data.PeriodeSelect = periode
    data.MembresValue = membres

    if alpha != "" || periode != "" || membres != "" {
        data.Artists = applyFilters(data.Artists, alpha, periode, membres)

        if len(data.Artists) == 0 {
            data.NoResult = true
            log.Println("⚠️ Aucun artiste ne correspond aux filtres")
        }
    }

    tmpl := template.Must(template.ParseFiles(filepath.Join("templates", "home.html")))
    tmpl.Execute(w, data)
}

// Fonction pour vérifier la période 
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

// Fonction pour filtrer selon les critères
func applyFilters(artists []models.Artist, alpha, periode, membresStr string) []models.Artist {
    var results []models.Artist
    
    // Convertir membres en int si fourni
    var membres int
    if membresStr != "" {
        fmt.Sscanf(membresStr, "%d", &membres)
    }
    
    // Filtrer les artistes
    for _, artist := range artists {
        // Filtre période
        if periode != "" && periode != "all" {
            if !correspondPeriode(artist.CreationDate, periode) {
                continue
            }
        }
        
        // Filtre nombre de membres
        if membres > 0 {
            if len(artist.Members) != membres {
                continue
            }
        }
        
        // Si l'artiste passe tous les filtres
        results = append(results, artist)
    }
    
    // Trier alphabétiquement si demandé
    if alpha == "1" {
        sort.Slice(results, func(i, j int) bool {
            return results[i].Name < results[j].Name
        })
    }
    
    return results
}