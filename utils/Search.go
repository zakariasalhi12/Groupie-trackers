package utils

import (
	"strconv"
	"strings"
)

func Search(searchStr string, Artists []ArtistData) []ArtistData {
	NewArtists := []ArtistData(nil)
	CreationDate, _ := strconv.Atoi(searchStr)

	for _, artist := range Artists {
		// Search for Band by members
		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), strings.ToLower(searchStr)) {
				if !ArtistChecker(artist, NewArtists) { // Checker if the band is already in out NewArtists Slice
					NewArtists = append(NewArtists, artist) // if i found the band append it in our Newaritst Slice
				}
				break
			}
		}
		// Search for Band by Locations
		for _, location := range artist.LocationsData {
			if strings.Contains(strings.ToLower(location), strings.ToLower(searchStr)) {
				if !ArtistChecker(artist, NewArtists) { // Checker if the band is already in out NewArtists Slice
					NewArtists = append(NewArtists, artist) // if i found the band append it in our Newaritst Slice
				}
				break
			}
		}

		// Search for Band by Bandname
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(searchStr)) {
			if !ArtistChecker(artist, NewArtists) { // Checker if the band is already in out NewArtists Slice
				NewArtists = append(NewArtists, artist) // if i found the band append it in our Newaritst Slice
			}
			// Search for Band by Firstalbum
		} else if artist.Firstalbum == searchStr { // We dont lowercase cause Firstalbum is number
			if !ArtistChecker(artist, NewArtists) { // Checker if the band is already in out NewArtists Slice
				NewArtists = append(NewArtists, artist) // if i found the band append it in our Newaritst Slice
			}
			// Search for Band by CreationDate
		} else if artist.Creationdate == CreationDate { // also we dont lowercase here couse CreationDate is int
			if !ArtistChecker(artist, NewArtists) { // Checker if the band is already in out NewArtists Slice
				NewArtists = append(NewArtists, artist) // if i found the band append it in our Newaritst Slice
			}
		}
	}
	return NewArtists
}
