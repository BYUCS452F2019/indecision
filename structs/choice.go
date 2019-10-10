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
