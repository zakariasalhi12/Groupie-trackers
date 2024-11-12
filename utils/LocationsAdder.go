package utils

import (
	"net/http"
)

func LocationsAdder(w http.ResponseWriter, art *[]ArtistData) error {
	ArtistLocations := Locations{}
	if err := Fetch(w, LocationsAPI, &ArtistLocations); err != nil {
		return err
	}

	for i := range *art {
		(*art)[i].LocationsData = ArtistLocations.Index[i].Locations
	}

	return nil
}
