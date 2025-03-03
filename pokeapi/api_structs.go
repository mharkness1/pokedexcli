package pokeapi

type LocationAreas struct {
	Count    int               `json:"count"`
	Next     string            `json:"next"`
	Previous any               `json:"previous"`
	Results  []LocationResults `json:"results"`
}

type LocationResults struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
