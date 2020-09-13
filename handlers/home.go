package handlers

import (
	"html/template"
	"net/http"
	"tv-guide/models"
	"tv-guide/services"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/home.html"))

	data := struct {
		Games []models.Game
	}{
		Games: services.GetGames(),
	}

	tmpl.Execute(w, data)
}
