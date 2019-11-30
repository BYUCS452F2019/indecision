package mongo

// DB is the MongoDB proxy object that implements our DB interface
type DB struct {
}

// NewDB returns a pointer to the DB object
func NewDB() *DB {
	return &DB{}
}

// OpenConnection opens a connection with the database
func (m *DB) OpenConnection() error {
	return nil
}

// CloseConnection closes the connection to the database
func (m *DB) CloseConnection() error {
	return nil
}
