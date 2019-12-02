package mongo

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// DB is the MongoDB proxy object that implements our DB interface
type DB struct {
	Client *mongo.Client
}

// NewDB returns a pointer to the DB object
func NewDB() *DB {
	return &DB{}
}

// OpenConnection opens a connection with the database
func (m *DB) OpenConnection() error {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
		return err
	}

	m.Client = client

	err = m.Client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		fmt.Printf("couldn't connect to the database: %v", err.Error())
		return err
	}

	fmt.Println("Successfully connected to the database")

	return nil
}

// CloseConnection closes the connection to the database
func (m *DB) CloseConnection() error {
	return nil
}
