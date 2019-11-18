package yelp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/jpw547/indecision/structs"
)

var testLat = 40.233845
var testLong = -111.658531

// GetRestaurantList gets a list of restaurant businesses from the Yelp API
func GetRestaurantList() ([]structs.Restaurant, error) {
	c := http.Client{}

	url := fmt.Sprintf("https://api.yelp.com/v3/businesses/search?latitude=%v&longitude=%v", testLat, testLong)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("failed to create the request: %s", err.Error())
		return []structs.Restaurant{}, err
	}

	key := os.Getenv("YELP_API_KEY")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", key))
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("failed to execute request: %s", err.Error())
		return []structs.Restaurant{}, err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("failed to read response body: %s", err.Error())
		return []structs.Restaurant{}, err
	}

	var yResponse structs.YelpResponse
	err = json.Unmarshal(data, &yResponse)
	if err != nil {
		fmt.Printf("failed to unmarshal response body: %s", err.Error())
		return []structs.Restaurant{}, err
	}

	var restaurants []structs.Restaurant
	for _, b := range yResponse.Businesses {
		restaurants = append(restaurants, structs.Restaurant{
			RestaurantID: b.ID,
			Name:         b.Name,
			Location:     fmt.Sprintf("%s, %s", b.Location.DisplayAddress[0], b.Location.DisplayAddress[1]),
			Price:        b.Price,
			Choice: structs.Choice{
				PictureURL: b.ImageURL,
				Type:       "restaurant",
			},
		})
	}

	return restaurants, nil
}
