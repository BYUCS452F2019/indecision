package sqlite

import (
	"github.com/jpw547/indecision/structs"
)

// CreateChoice records a choice in the database
func (s *DB) CreateChoice(choice structs.Choice) (structs.Choice, error) {
	var toReturn structs.Choice

	// do the thing

	return toReturn, nil
}

// GetChoice returns a choice from the database
func (s *DB) GetChoice(id string) (structs.Choice, error) {
	var toReturn structs.Choice

	// do the thing

	return toReturn, nil
}

// UpdateChoice updates a choice in the database
func (s *DB) UpdateChoice(id string, choice structs.Choice) (structs.Choice, error) {
	var toReturn structs.Choice

	// do the thing

	return toReturn, nil
}

// DeleteChoice deletes a choice from the database
func (s *DB) DeleteChoice(id string) error {
	return nil
}
