package handlers

import (
	"fmt"
	"github.com/jpw547/indecision/structs"
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"context"
	"log"

)

type noSQLStore struct {
	clientOptions *options.ClientOptions
}

// InitStore initializes our data store
func InitNoSQLStore(insertTestData bool) error {
	
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

		if(insertTestData){

		// Connect to MongoDB
		client, err := mongo.Connect(context.TODO(), clientOptions)

		if err != nil {
			return err
		}

		// Check the connection
		err = client.Ping(context.TODO(), nil)

		if err != nil {
			return err
		}

		fmt.Println("Connected to MongoDB!")

		if err != nil {
			return err
		}
		choices := client.Database("indecision").Collection("choices")
			// insertMovies := []interface{}{movies}
			for _, movie := range movies{
				_,err=choices.InsertOne(context.TODO(), movie)
				if err != nil {
					fmt.Println(err)
				}
			}
			for _, f := range food{
				_,err=choices.InsertOne(context.TODO(), f)
				if err != nil {
					fmt.Println(err)
				}
			}
	}
	store = &noSQLStore{clientOptions: clientOptions}
	return nil
}

// CreateUser performs the database transaction to add a user to the database
func (store *noSQLStore) CreateUser(user *structs.User) error {
	client, err := mongo.Connect(context.TODO(), store.clientOptions)
    users := client.Database("indecision").Collection("users")
	_,err=users.InsertOne(context.TODO(), user)
    if err != nil {
        return err
    }
	return nil
}

// CreateMovie performs thee database transaction to add a movie to the database
func (store *noSQLStore) CreateMovie(movie *structs.Movie) error {
	client, err := mongo.Connect(context.TODO(), store.clientOptions)
    choices := client.Database("indecision").Collection("choices")
	_,err=choices.InsertOne(context.TODO(), movie)
    if err != nil {
        return err
    }
	return nil
}

// CreateRestaurant performs the database transaction to add a restaurant to the database
func (store *noSQLStore) CreateRestaurant(restaurant *structs.Restaurant) error {
	client, err := mongo.Connect(context.TODO(), store.clientOptions)
    choices := client.Database("indecision").Collection("choices")
	_,err=choices.InsertOne(context.TODO(), restaurant)
    if err != nil {
        return err
    }
	return nil
}

// DeleteUserByUsername performs the database transaction that deletes a user from the database
func (store *noSQLStore) DeleteUserByUsername(username string) error {
	client, err := mongo.Connect(context.TODO(), store.clientOptions)
    users := client.Database("indecision").Collection("users")
	_, err = users.DeleteOne(context.TODO(),bson.M{"username": username})
	if err != nil {
		return err
	}
	return nil
}

// DeleteChoiceByID performs the database transaction that deletes a choice from the database
func (store *noSQLStore) DeleteChoiceByID(id string) error {
	client, err := mongo.Connect(context.TODO(), store.clientOptions)
    choices := client.Database("indecision").Collection("choices")
	_, err = choices.DeleteOne(context.TODO(),bson.M{"id": id})
	if err != nil {
		return err
	}
	return nil

}

// GetUser performs the database transaction that gets a user from the database
func (store *noSQLStore) GetUser(user *structs.User) ([]*structs.User, error) {
	users := []*structs.User{}
    filter := bson.M{"username":user.Username}
	client, err := mongo.Connect(context.TODO(), store.clientOptions)
    choices := client.Database("indecision").Collection("choices")
	cur,err := choices.Find(context.TODO(),filter)
	for cur.Next(context.TODO()) {
		var elem structs.User
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
  
		users=append(users, &elem)
	}
	if err != nil {
		return nil,err
	} 
	return users, nil
}

// GetUsers performs the database transaction that gets a list of users from the database
func (store *noSQLStore) GetUsers() ([]*structs.User, error) {
	users := []*structs.User{}
	filter := bson.D{}
	client, err := mongo.Connect(context.TODO(), store.clientOptions)
    choices := client.Database("indecision").Collection("choices")

	cur,err := choices.Find(context.TODO(),filter)
	for cur.Next(context.TODO()) {
		var elem structs.User
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
  
		users=append(users, &elem)
	}
	if err != nil {
		return nil,err
	}  
	return users, nil
}

// GetFood performs the database transaction that gets a list of restaurants from the database
func (store *noSQLStore) GetFood() ([]*structs.Restaurant, error) {

	client, err := mongo.Connect(context.TODO(), store.clientOptions)
    choices := client.Database("indecision").Collection("choices")

	restaurants := []*structs.Restaurant{}

    filter := bson.M{"type":"restaurant"}

	cur,err := choices.Find(context.TODO(),filter)
	if err != nil {
		fmt.Println(err)
		log.Println(err)

		return nil,err
	}
	log.Println(cur)

	for cur.Next(context.TODO()) {
		var elem structs.Restaurant
		err := cur.Decode(&elem)
		if err != nil {
			fmt.Println(err)
			log.Println(cur)

			return nil, err
		}
  
		restaurants=append(restaurants, &elem)
	}
	return restaurants, nil
}

// GetMovies performs the database transaction that gets a list of movies from the database
func (store *noSQLStore) GetMovies() ([]*structs.Movie, error) {
	movies := []*structs.Movie{}

    filter := bson.M{"type":"movie"}
	client, err := mongo.Connect(context.TODO(), store.clientOptions)
    choices := client.Database("indecision").Collection("choices")
	cur,err := choices.Find(context.TODO(),filter)
	for cur.Next(context.TODO()) {
		var elem structs.Movie
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		movies=append(movies, &elem)
	}
	if err != nil {
		return nil,err
	}
	return movies, nil
}
