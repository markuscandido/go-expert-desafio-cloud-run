package entity

// Location represents a geographic location from CEP lookup
type Location struct {
	City  string
	State string
}

// NewLocation creates a new Location instance
func NewLocation(city, state string) *Location {
	return &Location{
		City:  city,
		State: state,
	}
}
