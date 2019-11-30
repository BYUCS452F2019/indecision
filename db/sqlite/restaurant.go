package sqlite

import (
	"github.com/jpw547/indecision/structs"
)

// CreateRestaurant adds a restaurant object to the database
func (s *DB) CreateRestaurant(restaurant structs.Restaurant) (structs.Restaurant, error) {
	var toReturn structs.Restaurant

	// do the thing

	return toReturn, nil
}

// GetRestaurant gets a restaurant object from the database
func (s *DB) GetRestaurant(id string) (structs.Restaurant, error) {
	var toReturn structs.Restaurant

	// do the thing

	return toReturn, nil
}

// UpdateRestaurant updates a restaurant object in the database
func (s *DB) UpdateRestaurant(id string, restaurant structs.Restaurant) (structs.Restaurant, error) {
	var toReturn structs.Restaurant

	// do the thing

	return toReturn, nil
}

// DeleteRestaurant deletes a restaurant object from the database
func (s *DB) DeleteRestaurant(id string) error {
	return nil
}

// GetRestaurantsByCategory gets all restaurant objects of a certain category
func (s *DB) GetRestaurantsByCategory(category string) ([]structs.Restaurant, error) {
	var toReturn []structs.Restaurant

	// do the thing

	return toReturn, nil
}
