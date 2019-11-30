package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// DB in the object that implements the DB interface
type DB struct {
	conn *sql.DB
}

// NewDB returns a pointer to the DB object
func NewDB() *DB {
	return &DB{}
}

// OpenConnection opens a connection with the database
func (s *DB) OpenConnection() error {
	fmt.Print("Tyring to open sqlite db...")
	conn, err := sql.Open("sqlite3", "db/indecision.db")
	if err != nil {
		fmt.Printf("failed to open sqlite db: %s", err.Error())
		return err
	}

	s.conn = conn

	return nil
}

// CloseConnection closes the connection to the database
func (s *DB) CloseConnection() error {
	err := s.conn.Close()
	if err != nil {
		fmt.Printf("failed to close sqlite db: %s", err.Error())
	}

	return err
}
