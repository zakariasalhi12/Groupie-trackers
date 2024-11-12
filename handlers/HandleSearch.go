package handlers

import (
	"html/template"
	"net/http"

	utils "groupie/utils"
)

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ErrorWriter(w, "Methode not allowed", 405)
		return
	}

	var Art []utils.ArtistData

	if err := utils.Fetch(w, utils.ArtistsAPI, &Art); err != nil {
		utils.ErrorWriter(w, err.Error(), 500)
		return
	}

	if err := utils.LocationsAdder(w, &Art); err != nil {
		utils.ErrorWriter(w, err.Error(), 500)
		return
	}

	if Search := r.FormValue("search"); Search != "" {
		Art = utils.Search(Search, Art)
	}

	// Parse the template
	t, err := template.ParseFiles("./templates/Result.html")
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
