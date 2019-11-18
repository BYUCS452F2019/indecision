package structs

// A User represents the information needed about someone using the application
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

// A Choice is the base level object for different options to choose from
type Choice struct {
	ID         string `json:"choiceID"`
	PictureURL string `json:"pictureUrl"`
	Type       string `json:"type"`
}

// A Restaurant is a type of Choice
type Restaurant struct {
	Choice
	RestaurantID string `json:"restaurantID"`
	Name         string `json:"name"`
	Category     string `json:"category,omitempty"`
	Location     string `json:"location,omitempty"`
	Price        string `json:"price,omitempty"`
}

// A Movie is a type of Choice
type Movie struct {
	Choice
	MovieID string `json:"movieID"`
	Title   string `json:"title"`
	Genre   string `json:"genre,omitempty"`
}

// A YelpBusiness is the object returned by their API
type YelpBusiness struct {
	ID          string `json:"id"`
	Alias       string `json:"alias"`
	Name        string `json:"name"`
	ImageURL    string `json:"image_url"`
	IsClosed    bool   `json:"is_closed"`
	URL         string `json:"url"`
	ReviewCount int    `json:"review_count"`
	Categories  []struct {
		Alias string `json:"alias"`
		Title string `json:"title"`
	} `json:"categories"`
	Rating      float64 `json:"rating"`
	Coordinates struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"coordinates"`
	Transactions []string `json:"transactions"`
	Price        string   `json:"price"`
	Location     struct {
		Address1       string   `json:"address1"`
		Address2       string   `json:"address2"`
		Address3       string   `json:"address3"`
		City           string   `json:"city"`
		ZipCode        string   `json:"zip_code"`
		Country        string   `json:"country"`
		State          string   `json:"state"`
		DisplayAddress []string `json:"display_address"`
	} `json:"location"`
	Phone        string  `json:"phone"`
	DisplayPhone string  `json:"display_phone"`
	Distance     float64 `json:"distance"`
}

// A YelpResponse is the response object returned by their API
type YelpResponse struct {
	Businesses []YelpBusiness `json:"businesses"`
}
