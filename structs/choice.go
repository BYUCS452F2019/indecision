package structs

// A User represents the information needed about someone using the application
type User struct {
	ID       string `json:"id"  bson:"id,omitempty"`
	Username string `json:"username"  bson:"username,omitempty"`
}

// A Choice is the base level object for different options to choose from
type Choice struct {
	ID         string `json:"choiceID"  bson:"id,omitempty"`
	PictureURL string `json:"pictureUrl"  bson:"pictureURL,omitempty"`
	Type       string `json:"type"  bson:"type,omitempty"`
}

// A Restaurant is a type of Choice
type Restaurant struct {
	Choice		 `bson:",inline"`
	RestaurantID string `json:"restaurantID"  bson:"restaurantID,omitempty"`
	Name         string `json:"name"  bson:"name,omitempty"`
	Category     string `json:"category,omitempty"  bson:"category,omitempty"`
	Location     string `json:"location,omitempty"  bson:"location,omitempty"`
	Price        string `json:"price,omitempty"  bson:"price,omitempty"`
}

// A Movie is a type of Choice
type Movie struct {
	Choice  `bson:",inline"`
	MovieID string `json:"movieID"  bson:"movieID,omitempty"`
	Title   string `json:"title"  bson:"title,omitempty"`
	Genre   string `json:"genre,omitempty"  bson:"genre,omitempty"`
}
