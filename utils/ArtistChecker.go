package utils

func ArtistChecker(ar ArtistData, arr []ArtistData) bool {
	for _, artist := range arr {
		if ar.Name == artist.Name {
			return true
		}
	}
	return false
}
