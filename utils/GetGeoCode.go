package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetGeoCode(w http.ResponseWriter, location string) ([]byte, error) {
	res, err := http.Get("https://maps.googleapis.com/maps/api/geocode/json?address=" + location + "&key=" + MapAPIKey)
	if err != nil {
		return []byte(nil), err
	}
	defer res.Body.Close()
	FullResponse, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte(nil), err
	}

	return FullResponse, nil
}

func GetGeoFromLocations(w http.ResponseWriter, LocationsData []string, GeoCodeStruct *[]GeoLocations) error {
	for _, location := range LocationsData {
		NewCord := GeoLocations{}
		Data, err := GetGeoCode(w, location)
		if err != nil {
			ErrorWriter(w, err.Error(), 500)
			return err
		}
		if err := json.Unmarshal(Data, &NewCord); err != nil {
			ErrorWriter(w, err.Error(), 500)
			return err
		}
		*GeoCodeStruct = append(*GeoCodeStruct, NewCord)
	}
	return nil
}

func GetCordinations(GeoCode *[]GeoLocations, Cordinations *[]Cordinates) {
	for _, Res := range *GeoCode {
		NewCord := Cordinates{}
		for _, Geo := range Res.Results {
			NewCord.Latitude = Geo.Geometry.Location.Lat
			NewCord.Longitude = Geo.Geometry.Location.Lng
		}
		*Cordinations = append(*Cordinations, NewCord)
	}
}
