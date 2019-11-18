package handlers

import (
	"fmt"
	"net/http"

	"github.com/jpw547/indecision/structs"
	"github.com/jpw547/indecision/yelp"
	"github.com/labstack/echo"
)

var food []structs.Restaurant

var movies []structs.Movie

func init() {
	food = []structs.Restaurant{
		{
			Choice: structs.Choice{
				ID:         "1",
				PictureURL: "https://130e178e8f8ba617604b-8aedd782b7d22cfe0d1146da69a52436.ssl.cf1.rackcdn.com/wendys-reaches-50-million-breach-settlement-banks-showcase_image-8-a-12032.jpg",
				Type:       "restaurant",
			},
			RestaurantID: "1",
			Name:         "Wendy's",
		},
		{
			Choice: structs.Choice{
				ID:         "2",
				PictureURL: "https://cdn.junglecreations.com/wp/junglecms/2019/04/77a4227c-mcdonalds-free-breakfast-feature.jpg",
				Type:       "restaurant",
			},
			RestaurantID: "2",
			Name:         "McDonald's",
		},
		{
			Choice: structs.Choice{
				ID:         "3",
				PictureURL: "https://assets.simpleviewinc.com/simpleview/image/fetch/c_limit,q_75,w_1200/https://assets.simpleviewinc.com/simpleview/image/upload/crm/sandysprings/Chick-fil-a-Northridge-DTO---logo-resized-95aaa9a05056b3a_95aaaa9b-5056-b3a8-493eb82fb0c5d336.jpg",
				Type:       "restaurant",
			},
			RestaurantID: "3",
			Name:         "Chik-Fil-A",
		},
		{
			Choice: structs.Choice{
				ID:         "4",
				PictureURL: "https://image.shutterstock.com/image-photo/columbusohusa-july-242017-exterior-olive-600w-1218524122.jpg",
				Type:       "restaurant",
			},
			RestaurantID: "4",
			Name:         "Olive Garden",
		},
		{
			Choice: structs.Choice{
				ID:         "5",
				PictureURL: "https://image.shutterstock.com/image-photo/miami-usa-august-22-2018-600w-1181057281.jpg",
				Type:       "restaurant",
			},
			RestaurantID: "5",
			Name:         "Chipotle",
		},
		{
			Choice: structs.Choice{
				ID:         "6",
				PictureURL: "https://static.wixstatic.com/media/5ac7a5_bae3cfebb7574f29b0bab1d58e4e5632~mv2_d_3000_1632_s_2.jpg/v1/fill/w_3000,h_1632,al_c/5ac7a5_bae3cfebb7574f29b0bab1d58e4e5632~mv2_d_3000_1632_s_2.jpg",
				Type:       "restaurant",
			},
			RestaurantID: "6",
			Name:         "Burger Supreme",
		},
		{
			Choice: structs.Choice{
				ID:         "7",
				PictureURL: "https://d3dx3xtlhhlssb.cloudfront.net/cms/meta/open-graph/caferio-ogimage.jpg",
				Type:       "restaurant",
			},
			RestaurantID: "7",
			Name:         "Cafe Rio",
		},
		{
			Choice: structs.Choice{
				ID:         "8",
				PictureURL: "https://ewscripps.brightspotcdn.com/dims4/default/8b3833e/2147483647/strip/true/crop/640x360+0+21/resize/1280x720!/quality/90/?url=https%3A%2F%2Fsharing.wxyz.com%2Fsharescnn%2Fphoto%2F2017%2F11%2F30%2F14805645242_7431b4e1b3_o_red-robin_73066708_ver1.0_640_480.jpg",
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
				PictureURL: "https://static01.nyt.com/images/2019/07/21/arts/21lionking-notebook1/21lionking-notebook1-articleLarge.jpg?quality=75&auto=webp&disable=upscale",
				Type:       "movie",
			},
			MovieID: "1",
			Title:   "The Lion King",
		},
		{
			Choice: structs.Choice{
				ID:         "10",
				PictureURL: "http://annyas.com/screenshots/images/2006/p-cn-04.jpg",
				Type:       "movie",
			},
			MovieID: "2",
			Title:   "The Prestige",
		},
		{
			Choice: structs.Choice{
				ID:         "11",
				PictureURL: "https://i.redd.it/dzqf2ph1e9v01.jpg",
				Type:       "movie",
			},
			MovieID: "3",
			Title:   "Avengers: Infinity War",
		},
		{
			Choice: structs.Choice{
				ID:         "12",
				PictureURL: "https://vignette.wikia.nocookie.net/logopedia/images/4/47/Wall-e-disneyscreencaps.com-248.jpg/revision/latest?cb=20161129004221",
				Type:       "movie",
			},
			MovieID: "4",
			Title:   "Wall-E",
		},
		{
			Choice: structs.Choice{
				ID:         "13",
				PictureURL: "https://i.ytimg.com/vi/uoRbFgGwDZY/maxresdefault.jpg",
				Type:       "movie",
			},
			MovieID: "5",
			Title:   "Willy Wonka and the Chocolate Factory",
		},
		{
			Choice: structs.Choice{
				ID:         "14",
				PictureURL: "https://tokybook.com/wp-content/uploads/2017/08/The-Return-of-the-King-2003-movie.jpg",
				Type:       "movie",
			},
			MovieID: "6",
			Title:   "The Lord of the Rings: The Return of the King",
		},
		{
			Choice: structs.Choice{
				ID:         "15",
				PictureURL: "http://annyas.com/screenshots/updates/wp-content/uploads/2016/01/strange-brew-blu-ray-movie-title-small.jpg",
				Type:       "movie",
			},
			MovieID: "7",
			Title:   "Strange Brew",
		},
		{
			Choice: structs.Choice{
				ID:         "16",
				PictureURL: "http://annyas.com/screenshots/images/2004/incredibles-blu-ray-movie-title.jpg",
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
		// val, err := store.GetFood()
		val, err := yelp.GetRestaurantList()
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSON(http.StatusOK, val)
		// return ctx.JSON(http.StatusOK, food)
	}
	if choiceType == "movies" {
		val, err := store.GetMovies()
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSON(http.StatusOK, val)
		// return ctx.JSON(http.StatusOK, movies)
	}

	return ctx.String(http.StatusBadRequest, "unknown choice type")
}

// DeleteChoiceByID deletes a choice with the given ID
func DeleteChoiceByID(ctx echo.Context) error {
	// get the type from the context
	choiceType := ctx.Param("type")
	fmt.Println(choiceType)
	err := store.DeleteChoiceByID(choiceType)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, "OK")
}

// CreateChoice records a choice in the database
func CreateChoice(ctx echo.Context) error {
	// get the type from the context
	// get the type from the context
	choiceType := ctx.Param("type")
	// TODO: implement database access for choices
	if choiceType == "food" {
		var r structs.Restaurant
		err := ctx.Bind(&r)
		err = store.CreateRestaurant(&r)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSON(http.StatusOK, "OK")
	}

	if choiceType == "movies" {
		var m structs.Movie
		err := ctx.Bind(&m)
		err = store.CreateMovie(&m)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSON(http.StatusOK, "OK")
	}

	return ctx.String(http.StatusBadRequest, "unknown choice type")
}

// GetNewList ...
func GetNewList(ctx echo.Context) error {
	y, err := yelp.GetRestaurantList()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, y)
}
