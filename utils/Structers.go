package utils

const (
	ArtistsAPI   = "https://groupietrackers.herokuapp.com/api/artists"
	LocationsAPI = "https://groupietrackers.herokuapp.com/api/locations"
	DatesAPI     = "https://groupietrackers.herokuapp.com/api/dates"
	RelationAPI  = "https://groupietrackers.herokuapp.com/api/relation"
	MapAPIKey    = "Your-API-Key"
)

type ArtistData struct {
	Id            int      `json:"id"`
	Image         string   `json:"image"`
	Name          string   `json:"name"`
	Members       []string `json:"members"`
	Creationdate  int      `json:"creationdate"`
	Firstalbum    string   `json:"firstalbum"`
	Locations     string   `json:"locations"`
	ConcertDates  string   `json:"concertDates"`
	Relations     string   `json:"relations"`
	RelationsData map[string][]string
	LocationsData []string
	DatesData     []string
	Cordinates    []Cordinates
	CordStr       string
}

type Location struct {
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Locations struct {
	Index []Location `json:"Index"`
}

type Dates struct {
	Dates []string `json:"dates"`
}

type Rolation struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

type FilterSearch struct {
	MinCreationDate int
	MaxCreationDate int
	MinFirstAlbum   int
	MaxFirstAlbulm  int
	Location        string
	Members         []string
}

type GeoLocations struct {
	Results []struct {
		Geometry struct {
			Location struct {
				Lat float64
				Lng float64
			}
		}
	}
}

type Cordinates struct {
	Latitude  float64
	Longitude float64
}

type Error struct {
	Error      string
	StatusCode int
}
