package sqlite

import (
	"fmt"
	"strconv"

	"github.com/jpw547/indecision/structs"
)

// CreateUser creates a user in the database
func (s *DB) CreateUser(username string) (structs.User, error) {
	var toReturn structs.User

	if s.conn == nil {
		fmt.Print("The connection is nil")
		return toReturn, nil
	}

	stmt, err := s.conn.Prepare("INSERT INTO Users(username) values(?)")
	if err != nil {
		fmt.Printf("failed to prepare statement: %s", err.Error())
		return toReturn, err
	}

	result, err := stmt.Exec(username)
	if err != nil {
		fmt.Printf("failed to execute statement: %s", err.Error())
		return toReturn, err
	}

	id, err := result.LastInsertId()

	toReturn.Username = username
	toReturn.ID = strconv.FormatInt(id, 10)

	return toReturn, nil
}

// GetUser returns a user from the database
func (s *DB) GetUser(username string) (structs.User, error) {
	var toReturn structs.User

	// do the thing

	return toReturn, nil
}

// UpdateUser updates a user in the database
func (s *DB) UpdateUser(username string, user structs.User) (structs.User, error) {
	var toReturn structs.User

	// do the thing

	return toReturn, nil
}

// DeleteUser deletes a user from the database
func (s *DB) DeleteUser(username string) error {
	return nil
}
