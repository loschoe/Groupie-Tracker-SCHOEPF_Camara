package main

import (
	"log"
	"net/http"
	"groupie-tracker/handlers"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlers.Start)
	http.HandleFunc("/accueil", handlers.Home)
	http.HandleFunc("/compare", handlers.Compare)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/artist/", handlers.Artist)

	log.Println("Démarrage du serveur...")
	log.Println("✅ Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}