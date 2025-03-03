package pokeapi

// Structs for the Location Areas - used by map and mapb commands.
type LocationAreas struct {
	Count    int               `json:"count"`
	Next     string            `json:"next"`
	Previous string            `json:"previous"`
	Results  []LocationResults `json:"results"`
}

type LocationResults struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
