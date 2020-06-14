package locations

// Country represents a country structure
type Country struct {
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	TimeZone       string         `json:"time_zone"`
	GeoInformation GeoInformation `json:"geo_information"`
	States         []State        `json:"states"`
}

// GeoInformation represents geo information
type GeoInformation struct {
	Location GeoLocation `json:"location"`
}

// GeoLocation represents a geo location
type GeoLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// State represent a state of a country
type State struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
