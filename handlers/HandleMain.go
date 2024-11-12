package handlers

import (
	"html/template"
	"net/http"

	utils "groupie/utils"
)

func HandleMain(w http.ResponseWriter, r *http.Request) {

	// Handle non-root paths
	if r.URL.Path != "/" {
		utils.ErrorWriter(w, "Error Page Not Found", 404)
		return
	}

	if r.Method != http.MethodGet {
		utils.ErrorWriter(w, "Method Not Allowed", 405)
		return
	}

	var Art []utils.ArtistData

	// Fetch artist data
	if err := utils.Fetch(w, utils.ArtistsAPI, &Art); err != nil {
		utils.ErrorWriter(w, err.Error(), 500)
		return
	}

	if err := utils.LocationsAdder(w, &Art); err != nil {
		utils.ErrorWriter(w, err.Error(), 500)
		return
	}

	// Parse the template
	t, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		utils.ErrorWriter(w, err.Error(), 500)
		return
	}
	// Execute the template with the paginated data
	if err := t.Execute(w, Art); err != nil {
		utils.ErrorWriter(w, err.Error(), 500)
		return
	}
}
