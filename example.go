package example

import (
	"encoding/json"
	"os"
)

func Add(a int, b int) int {
	return a + b
}

func Subtract(a int, b int) int {
	return a - b
}

const Name = "example"

type Person struct {
	Name string
	Age  int
}

func CountryStateCity(country string, state string, city string) []Final_Model {
	file, _ := os.Open("countries-states-cities.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	var country_state_city []Final_Model
	decoder.Decode(&country_state_city)

	// for _, country_model := range country_state_city {

	// 	if country == country_model.Country_Name {

	// 		for currentIndex, state_city := range country_model.StateCity {

	// 			if state_city.State == "" && len(state_city.Cities) == 0 {
	// 				if state != "" || city != "" {
	// 					return "This Country has no state & city"
	// 				}
	// 				if state == "" || city == "" {
	// 					return "1"
	// 				}
	// 			}
	// 			if len(state_city.State) != 0 && state == "" {
	// 				return "Invalid State"
	// 			}
	// 			if state_city.State == state {
	// 				if len(state_city.Cities) == 0 && city != "" {
	// 					return "This Country has no city"
	// 				}
	// 				if len(state_city.Cities) != 0 && city == "" {
	// 					return "Invalid city"
	// 				}
	// 				for _, city_name := range state_city.Cities {
	// 					if city_name == city {
	// 						return "2"
	// 					}
	// 				}
	// 				return "Invalid city"
	// 			} else {
	// 				if currentIndex == len(country_model.StateCity)-1 {
	// 					return "Invalid state"
	// 				}
	// 			}

	// 		}
	// 	}
	// }
	return country_state_city
}

type State_City struct {
	State  string   `json:"state" bson:"state"`
	Cities []string `json:"cities" bson:"cities"`
}

type Final_Model struct {
	Country_Name string       `json:"country_name" bson:"country_name"`
	StateCity    []State_City `json:"state_city" bson:"state_city"`
}
