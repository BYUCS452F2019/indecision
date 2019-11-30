package mongo

import "github.com/jpw547/indecision/structs"

// CreateChoice records a choice in the database
func (m *DB) CreateChoice(choice structs.Choice) (structs.Choice, error) {
	var toReturn structs.Choice

	// do the thing

	return toReturn, nil
}

// GetChoice returns a choice from the database
func (m *DB) GetChoice(id string) (structs.Choice, error) {
	var toReturn structs.Choice

	// do the thing

	return toReturn, nil
}

// UpdateChoice updates a choice in the database
func (m *DB) UpdateChoice(id string, choice structs.Choice) (structs.Choice, error) {
	var toReturn structs.Choice

	// do the thing

	return toReturn, nil
}

// DeleteChoice deletes a choice from the database
func (m *DB) DeleteChoice(id string) error {
	return nil
}
