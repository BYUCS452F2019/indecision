package handlers

import (
	"net/http"

	"github.com/jpw547/indecision/structs"
	"github.com/labstack/echo"
)

var food []structs.Restaurant

var movies []structs.Movie

func init() {
	food = []structs.Restaurant{
		{
			Choice: structs.Choice{
				ID:         "1",
				PictureURL: "",
				Type:       "restaurant",
			},
			RestaurantID: "1",
			Name:         "Wendy's",
		},
		{
			Choice: structs.Choice{
				ID:         "2",
				PictureURL: "",
				Type:       "restaurant",
			},
			RestaurantID: "2",
			Name:         "McDonald's",
		},
		{
			Choice: structs.Choice{
				ID:         "3",
				PictureURL: "",
				Type:       "restaurant",
			},
			RestaurantID: "3",
			Name:         "Chik-Fil-A",
		},
		{
			Choice: structs.Choice{
				ID:         "4",
				PictureURL: "",
				Type:       "restaurant",
			},
			RestaurantID: "4",
			Name:         "Olive Garden",
		},
		{
			Choice: structs.Choice{
				ID:         "5",
				PictureURL: "",
				Type:       "restaurant",
			},
			RestaurantID: "5",
			Name:         "Chipotle",
		},
		{
			Choice: structs.Choice{
				ID:         "6",
				PictureURL: "",
				Type:       "restaurant",
			},
			RestaurantID: "6",
			Name:         "Burger Supreme",
		},
		{
			Choice: structs.Choice{
				ID:         "7",
				PictureURL: "",
				Type:       "restaurant",
			},
			RestaurantID: "7",
			Name:         "Cafe Rio",
		},
		{
			Choice: structs.Choice{
				ID:         "8",
				PictureURL: "",
				Type:       "restaurant",
			},
			RestaurantID: "8",
			Name:         "Red Robin",
		},
	}

	movies = []structs.Movie{
		{
			Choice: structs.Choice{
				ID:         "9",
				PictureURL: "",
				Type:       "movie",
			},
			MovieID: "1",
			Title:   "The Lion King",
		},
		{
			Choice: structs.Choice{
				ID:         "10",
				PictureURL: "",
				Type:       "movie",
			},
			MovieID: "2",
			Title:   "The Prestige",
		},
		{
			Choice: structs.Choice{
				ID:         "11",
				PictureURL: "",
				Type:       "movie",
			},
			MovieID: "3",
			Title:   "Avengers: Infinity War",
		},
		{
			Choice: structs.Choice{
				ID:         "12",
				PictureURL: "",
				Type:       "movie",
			},
			MovieID: "4",
			Title:   "Wall-E",
		},
		{
			Choice: structs.Choice{
				ID:         "13",
				PictureURL: "",
				Type:       "movie",
			},
			MovieID: "5",
			Title:   "Willy Wonka and the Chocolate Factory",
		},
		{
			Choice: structs.Choice{
				ID:         "14",
				PictureURL: "",
				Type:       "movie",
			},
			MovieID: "6",
			Title:   "The Lord of the Rings: The Return of the King",
		},
		{
			Choice: structs.Choice{
				ID:         "15",
				PictureURL: "",
				Type:       "movie",
			},
			MovieID: "7",
			Title:   "Strange Brew",
		},
		{
			Choice: structs.Choice{
				ID:         "16",
				PictureURL: "",
				Type:       "movie",
			},
			MovieID: "8",
			Title:   "The Incredibles",
		},
	}
}

// GetChoicesByType returns a list of options to choose from that are of the specified type
func GetChoicesByType(ctx echo.Context) error {
	// get the type from the context
	choiceType := ctx.Param("type")

	// TODO: implement database access for choices
	if choiceType == "food" {
		return ctx.JSON(http.StatusOK, food)
	}
	if choiceType == "movies" {
		return ctx.JSON(http.StatusOK, movies)
	}

	return ctx.String(http.StatusBadRequest, "unknown choice type")
}
