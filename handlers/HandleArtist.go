package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	utils "groupie/utils"
)

func HandleArtists(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ErrorWriter(w, "Method Not Allowed", 405)
		return
	}

	id := r.FormValue("id")
	if id == "" {
		utils.ErrorWriter(w, "Missing id parameter", 400)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.ErrorWriter(w, "Bad Request artist with that id not found", 400)
		return
	}

	if idInt < 1 || idInt > 52 {
		utils.ErrorWriter(w, "Bad Request artist with that id not found", 400)
		return
	}

	var Art utils.ArtistData
	var Rolation utils.Rolation
	var Location utils.Location
	var Dates utils.Dates
	var GeoCode []utils.GeoLocations
	var Cordinations []utils.Cordinates

	if err := utils.Fetch(w, utils.ArtistsAPI+"/"+id, &Art); err != nil {
		utils.ErrorWriter(w, err.Error(), 500)
		return
	}
	if err := utils.Fetch(w, Art.Locations, &Location); err != nil {
		utils.ErrorWriter(w, err.Error(), 500)
		return
	}

	if err := utils.Fetch(w, Location.Dates, &Dates); err != nil {
		utils.ErrorWriter(w, err.Error(), 500)
		return
	}

	if err := utils.Fetch(w, Art.Relations, &Rolation); err != nil {
		utils.ErrorWriter(w, err.Error(), 500)
		return
	}

	Art.LocationsData = Location.Locations
	Art.DatesData = Dates.Dates
	Art.RelationsData = Rolation.DatesLocations

	if err := utils.GetGeoFromLocations(w, Art.LocationsData, &GeoCode); err != nil {
		return
	}
	utils.GetCordinations(&GeoCode, &Cordinations)
	json, err := json.Marshal(Cordinations)
	if err != nil {
		utils.ErrorWriter(w, err.Error(), 500)
		return
	}
	Art.CordStr = string(json)

	t, err := template.ParseFiles("./templates/ContentData.html")
	if err != nil {
		utils.ErrorWriter(w, err.Error(), 500)
		return
	}

	if err := t.Execute(w, Art); err != nil {
		utils.ErrorWriter(w, err.Error(), 500)
		return
	}
}
