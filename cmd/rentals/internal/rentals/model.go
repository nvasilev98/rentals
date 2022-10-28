package rentals

type RentalResponse struct {
	ID              int              `json:"id"`
	Name            string           `json:"name"`
	Description     string           `json:"description"`
	Type            string           `json:"type"`
	VehicleMake     string           `json:"make"`
	VehicleModel    string           `json:"model"`
	VehicleYear     string           `json:"year"`
	VehicleLength   float32          `json:"length"`
	Sleeps          int              `json:"sleeps"`
	PrimaryImageURL string           `json:"primary_image_url"`
	Price           PriceResponse    `json:"price"`
	Location        LocationResponse `json:"location"`
	User            UserResponse     `json:"user"`
}

type PriceResponse struct {
	Day int `json:"day"`
}

type LocationResponse struct {
	HomeCity    string  `json:"city"`
	HomeState   string  `json:"state"`
	HomeZIP     string  `json:"zip"`
	HomeCountry string  `json:"country"`
	LAT         float32 `json:"lat"`
	LNG         float32 `json:"lng"`
}

type UserResponse struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
