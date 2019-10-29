package handlers

import (
	"database/sql"
	"github.com/jpw547/indecision/structs"
	"fmt"

)

type Store interface {
	CreateUser(user *structs.User) error
	GetUsers() ([]*structs.User, error)
	GetUser(user *structs.User) ([]*structs.User, error)
	GetFood() ([]*structs.Restaurant, error)
	GetMovies() ([]*structs.Movie, error)
}

type dbStore struct {
	db *sql.DB
}
var store Store

func InitStore(db  *sql.DB) {
	store = &dbStore{db: db}
}

func (store *dbStore) CreateUser(user *structs.User) error {
	sqlStatement:=fmt.Sprint("INSERT INTO user(username) VALUES ('",user.Username,"')")
	fmt.Println(sqlStatement)
	_, err := store.db.Query(sqlStatement)
	return err
}
func (store *dbStore) GetUser(user *structs.User) ([]*structs.User, error) {
	sqlStatement:=fmt.Sprint("SELECT * FROM user WHERE username='",user.Username,"'")
	rows, err := store.db.Query(sqlStatement)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := []*structs.User{}
	for rows.Next() {
		user := &structs.User{}
		if err := rows.Scan(&user.Username); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (store *dbStore) GetUsers() ([]*structs.User, error) {
	rows, err := store.db.Query("SELECT username from user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := []*structs.User{}
	for rows.Next() {
		user := &structs.User{}
		if err := rows.Scan(&user.Username); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
func (store *dbStore) GetFood() ([]*structs.Restaurant, error) {

	rows, err := store.db.Query("SELECT choice_id,picture_url,type,restaurant_id,name from choice join restaurant on choice.id=restaurant.choice_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	restaurants := []*structs.Restaurant{}
	for rows.Next() {
		restaurant := &structs.Restaurant{}
		if err := rows.Scan(&restaurant.Choice.ID,&restaurant.Choice.PictureURL,&restaurant.Choice.Type,&restaurant.ID,&restaurant.Name); err != nil {
			return nil, err
		}
		restaurants = append(restaurants, restaurant)
	}
	return restaurants, nil
}
func (store *dbStore) GetMovies() ([]*structs.Movie, error) {
	rows, err := store.db.Query("SELECT choice_id,picture_url,type,movie_id,title from choice join movie on choice.id=movie.choice_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	movies := []*structs.Movie{}
	for rows.Next() {
		movie := &structs.Movie{}
		if err := rows.Scan(&movie.Choice.ID,&movie.Choice.PictureURL,&movie.Choice.Type,&movie.ID,&movie.Title); err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}
