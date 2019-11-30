package mongo

import "github.com/jpw547/indecision/structs"

// CreateUser creates a user in the database
func (m *DB) CreateUser(username string) (structs.User, error) {
	var toReturn structs.User

	// do the thing

	return toReturn, nil
}

// GetUser returns a user from the database
func (m *DB) GetUser(username string) (structs.User, error) {
	var toReturn structs.User

	// do the thing

	return toReturn, nil
}

// UpdateUser updates a user in the database
func (m *DB) UpdateUser(username string, user structs.User) (structs.User, error) {
	var toReturn structs.User

	// do the thing

	return toReturn, nil
}

// DeleteUser deletes a user from the database
func (m *DB) DeleteUser(username string) error {
	return nil
}
