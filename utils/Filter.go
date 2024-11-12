package utils

import (
	"strconv"
	"strings"
)

func Filter(arr []ArtistData, filter FilterSearch) []ArtistData {
	NewArtists := []ArtistData{}
	for _, artist := range arr {
		FirstAlbYear := strings.Split(artist.Firstalbum, "-")
		Year, err := strconv.Atoi(FirstAlbYear[2])
		if err != nil {
			// Handle invalid year
			continue
		}

		// Filter by creation date range
		if filter.MinCreationDate != 0 || filter.MaxCreationDate != 0 {
			if artist.Creationdate < filter.MinCreationDate || artist.Creationdate > filter.MaxCreationDate {
				continue
			}
		}

		// Filter by first album year range
		if filter.MinFirstAlbum != 0 || filter.MaxFirstAlbulm != 0 {
			if Year < filter.MinFirstAlbum || Year > filter.MaxFirstAlbulm {
				continue
			}
		}

		// Filter by location
		if filter.Location != "" {
			foundLocation := false
			for _, location := range artist.LocationsData {
				if ParseLocation(location) == ParseLocation(filter.Location) || strings.EqualFold(location, filter.Location) {
					foundLocation = true
					break
				}
			}
			if !foundLocation {
				continue
			}
		}

		// Filter by number of members
		memberCount := len(artist.Members)
		memberMatch := false
		for _, Number := range filter.Members {
			nm, _ := strconv.Atoi(Number)
			if memberCount == nm {
				memberMatch = true
				break
			}
		}
		if len(filter.Members) > 0 && !memberMatch {
			continue
		}

		// Ensure no duplicate artists
		if !ArtistChecker(artist, NewArtists) {
			NewArtists = append(NewArtists, artist)
		}
	}

	return NewArtists
}

func ParseLocation(s string) string {
	res := ""
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			res += string(char)
		} else if char >= 'A' && char <= 'Z' {
			res += string(char + 32)
		}
	}

	return res
}
