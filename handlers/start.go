// Page de démarrage :
package handlers

import (
    "html/template"
    "net/http"
    "path/filepath"
)

func Start(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles(filepath.Join("templates", "start.html")))
    tmpl.Execute(w, nil)
}