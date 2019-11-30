package db

import (
	"github.com/jpw547/indecision/db/mongo"
	"github.com/jpw547/indecision/db/sqlite"
	"github.com/jpw547/indecision/structs"
)

// DB .
type DB interface {
	OpenConnection() error
	CloseConnection() error

	CreateUser(username string) (structs.User, error)
	GetUser(username string) (structs.User, error)
	UpdateUser(username string, user structs.User) (structs.User, error)
	DeleteUser(username string) error

	CreateChoice(choice structs.Choice) (structs.Choice, error)
	GetChoice(id string) (structs.Choice, error)
	UpdateChoice(id string, choice structs.Choice) (structs.Choice, error)
	DeleteChoice(id string) error

	CreateRestaurant(restaurant structs.Restaurant) (structs.Restaurant, error)
	GetRestaurant(id string) (structs.Restaurant, error)
	UpdateRestaurant(id string, restaurant structs.Restaurant) (structs.Restaurant, error)
	DeleteRestaurant(id string) error
	GetRestaurantsByCategory(category string) ([]structs.Restaurant, error)

	CreateHistoryRecord(record structs.HistoryRecord) (structs.HistoryRecord, error)
	GetHistoryRecord(id string) (structs.HistoryRecord, error)
	UpdateHistoryRecord(id string, record structs.HistoryRecord) (structs.HistoryRecord, error)
	DeleteHistoryRecord(id string) error
	GetHistoryRecordsByUser(userID string) ([]structs.HistoryRecord, error)
	GetLastRecordByUser(userID string) (structs.HistoryRecord, error)
	GetHistoryRecordsByChoice(choiceID string) ([]structs.HistoryRecord, error)
	GetLastRecordByChoice(choiceID string) (structs.HistoryRecord, error)
}

var database DB
var dbType = "sql"

// GetDB returns the instance of the database object
func GetDB() DB {
	if database == nil {
		if dbType == "sql" {
			database = sqlite.NewDB()
		}
		if dbType == "mongo" {
			database = mongo.NewDB()
		}
	}

	return database
}
