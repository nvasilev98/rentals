package rentals

type Model struct {
	ID              int
	Name            string
	Description     string
	Type            string
	VehicleMake     string
	VehicleModel    string
	VehicleYear     int
	VehicleLength   float32
	Sleeps          int
	PrimaryImageURL string
	PricePerDay     int
	HomeCity        string
	HomeState       string
	HomeZIP         string
	HomeCountry     string
	LAT             float32
	LNG             float32
	UserID          int
	FirstName       string
	LastName        string
}
