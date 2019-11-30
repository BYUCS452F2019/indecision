package mongo

import "github.com/jpw547/indecision/structs"

// CreateRestaurant adds a restaurant object to the database
func (m *DB) CreateRestaurant(restaurant structs.Restaurant) (structs.Restaurant, error) {
	var toReturn structs.Restaurant

	// do the thing

	return toReturn, nil
}

// GetRestaurant gets a restaurant object from the database
func (m *DB) GetRestaurant(id string) (structs.Restaurant, error) {
	var toReturn structs.Restaurant

	// do the thing

	return toReturn, nil
}

// UpdateRestaurant updates a restaurant object in the database
func (m *DB) UpdateRestaurant(id string, restaurant structs.Restaurant) (structs.Restaurant, error) {
	var toReturn structs.Restaurant

	// do the thing

	return toReturn, nil
}

// DeleteRestaurant deletes a restaurant object from the database
func (m *DB) DeleteRestaurant(id string) error {
	return nil
}

// GetRestaurantsByCategory gets all restaurant objects of a certain category
func (m *DB) GetRestaurantsByCategory(category string) ([]structs.Restaurant, error) {
	var toReturn []structs.Restaurant

	// do the thing

	return toReturn, nil
}
