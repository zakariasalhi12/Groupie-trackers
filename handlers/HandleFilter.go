package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	utils "groupie/utils"
)

func HandleFilter(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		utils.ErrorWriter(w, "Method Not Allowed", 405)
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

	if err := r.ParseForm(); err != nil {
		utils.ErrorWriter(w, err.Error(), 500)
		return
	}

	MinCreationDate, _ := strconv.Atoi(r.Form.Get("creationdate"))
	MaxCreationDate, _ := strconv.Atoi(r.Form.Get("creationdatemax"))
	MinFirstAlbum, _ := strconv.Atoi(r.Form.Get("firstalbum"))
	MaxFirstAlbum, _ := strconv.Atoi(r.Form.Get("firstalbummax"))

	FilterStruct := utils.FilterSearch{
		MinCreationDate: MinCreationDate,
		MaxCreationDate: MaxCreationDate,
		MinFirstAlbum:   MinFirstAlbum,
		MaxFirstAlbulm:  MaxFirstAlbum,
		Location:        r.FormValue("location"),
		Members:         r.Form["members[]"],
	}

	Art = utils.Filter(Art, FilterStruct)

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
