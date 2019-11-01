package handlers

import (
	"database/sql"
	"fmt"

	"github.com/jpw547/indecision/structs"
)

// Store is an interface for our database usage
type Store interface {
	CreateUser(user *structs.User) error
	GetUsers() ([]*structs.User, error)
	GetUser(user *structs.User) ([]*structs.User, error)
	GetFood() ([]*structs.Restaurant, error)
	GetMovies() ([]*structs.Movie, error)
	DeleteUserByUsername(username string) error
	DeleteChoiceByID(id string) error
	CreateMovie(movie *structs.Movie) error
	CreateRestaurant(restaurant *structs.Restaurant) error
}

type dbStore struct {
	db *sql.DB
}

var store Store

// InitStore initializes our data store
func InitStore(db *sql.DB) {
	store = &dbStore{db: db}
}

// CreateUser performs the database transaction to add a user to the database
func (store *dbStore) CreateUser(user *structs.User) error {
	sqlStatement := fmt.Sprint("INSERT INTO user(username) VALUES ('", user.Username, "')")

	fmt.Println(sqlStatement)

	_, err := store.db.Query(sqlStatement)
	if err != nil {
		fmt.Printf("failed to execute create user query on the database: %v", err)
		return err
	}

	return nil
}

// CreateMovie performs thee database transaction to add a movie to the database
func (store *dbStore) CreateMovie(movie *structs.Movie) error {
	sqlStatement := fmt.Sprint("INSERT INTO choice(type,picture_url) VALUES ('", movie.Choice.Type, "','", movie.Choice.PictureURL, "')")

	fmt.Println(sqlStatement)

	res, err := store.db.Exec(sqlStatement)

	choiceID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	sqlStatement = fmt.Sprint("INSERT INTO movie(choice_id,title,genre) VALUES ('", choiceID, "','", movie.Title, "','", movie.Genre, "')")

	fmt.Println(sqlStatement)

	_, err = store.db.Query(sqlStatement)
	if err != nil {
		fmt.Printf("failed to execute create movie query on the database: %v", err)
		return err
	}

	return nil
}

// CreateRestaurant performs the database transaction to add a restaurant to the database
func (store *dbStore) CreateRestaurant(restaurant *structs.Restaurant) error {
	sqlStatement := fmt.Sprint("INSERT INTO choice(type,picture_url) VALUES ('", restaurant.Choice.Type, "','", restaurant.Choice.PictureURL, "')")

	fmt.Println(sqlStatement)

	res, err := store.db.Exec(sqlStatement)
	if err != nil {
		return err
	}

	choiceID, err := res.LastInsertId()

	sqlStatement = fmt.Sprint("INSERT INTO restaurant(choice_id,name,category,location,price) VALUES ('", choiceID, "','", restaurant.Name, "','", restaurant.Category, "','", restaurant.Location, "','", restaurant.Price, "')")

	fmt.Println(sqlStatement)

	_, err = store.db.Query(sqlStatement)
	if err != nil {
		fmt.Printf("failed to execute create restaurant query on the database: %v", err)
		return err
	}

	return nil
}

// DeleteUserByUsername performs the database transaction that deletes a user from the database
func (store *dbStore) DeleteUserByUsername(username string) error {
	sqlStatement := fmt.Sprint("DELETE FROM user where username= '", username, "'")

	fmt.Println(sqlStatement)

	_, err := store.db.Query(sqlStatement)
	if err != nil {
		fmt.Printf("failed to execute delete user query on the database: %v", err)
		return err
	}

	return nil
}

// DeleteChoiceByID performs the database transaction that deletes a choice from the database
func (store *dbStore) DeleteChoiceByID(id string) error {

	sqlStatement := fmt.Sprint("DELETE FROM movie where choice_id= ", id)
	fmt.Println(id)

	fmt.Println(sqlStatement)
	_, err := store.db.Query(sqlStatement)
	if err != nil {
		fmt.Printf("failed to execute delete movie query on the database: %v", err)
		return err
	}

	sqlStatement = fmt.Sprint("DELETE FROM restaurant where choice_id= ", id)

	fmt.Println(sqlStatement)

	_, err = store.db.Query(sqlStatement)
	if err != nil {
		fmt.Printf("failed to execute delete restaurant query on the database: %v", err)
		return err
	}

	sqlStatement = fmt.Sprint("DELETE FROM choice where id= ", id)

	fmt.Println(sqlStatement)

	_, err = store.db.Query(sqlStatement)
	if err != nil {
		fmt.Printf("failed to execute delete choice query on the database: %v", err)
		return err
	}

	return nil
}

// GetUser performs the database transaction that gets a user from the database
func (store *dbStore) GetUser(user *structs.User) ([]*structs.User, error) {
	sqlStatement := fmt.Sprint("SELECT * FROM user WHERE username='", user.Username, "'")

	rows, err := store.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []*structs.User{}

	for rows.Next() {
		user := &structs.User{}
		if err := rows.Scan(&user.ID, &user.Username); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// GetUsers performs the database transaction that gets a list of users from the database
func (store *dbStore) GetUsers() ([]*structs.User, error) {
	rows, err := store.db.Query("SELECT * from user")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []*structs.User{}

	for rows.Next() {
		user := &structs.User{}
		if err := rows.Scan(&user.ID, &user.Username); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// GetFood performs the database transaction that gets a list of restaurants from the database
func (store *dbStore) GetFood() ([]*structs.Restaurant, error) {

	rows, err := store.db.Query("SELECT choice_id,picture_url,type,restaurant_id,name from choice join restaurant on choice.id=restaurant.choice_id")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	restaurants := []*structs.Restaurant{}

	for rows.Next() {
		restaurant := &structs.Restaurant{}
		if err := rows.Scan(&restaurant.Choice.ID, &restaurant.Choice.PictureURL, &restaurant.Choice.Type, &restaurant.RestaurantID, &restaurant.Name); err != nil {
			return nil, err
		}

		restaurants = append(restaurants, restaurant)
	}

	return restaurants, nil
}

// GetMovies performs the database transaction that gets a list of movies from the database
func (store *dbStore) GetMovies() ([]*structs.Movie, error) {
	rows, err := store.db.Query("SELECT choice_id,picture_url,type,movie_id,title from choice join movie on choice.id=movie.choice_id")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	movies := []*structs.Movie{}

	for rows.Next() {
		movie := &structs.Movie{}
		if err := rows.Scan(&movie.Choice.ID, &movie.Choice.PictureURL, &movie.Choice.Type, &movie.ID, &movie.Title); err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	return movies, nil
}
