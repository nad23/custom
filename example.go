package example

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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

func CountryStateCity(country string, state string, city string) string {

	files, err := http.Get("https://github.com//nad23/custompackagesknvnknk/main/countries-states-cities.json")
	defer files.Body.Close()
	if err != nil {
		return err.Error() + "  1"
	}
	countriesData, err := ioutil.ReadAll(files.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// defer file.Close()
	// decoder := json.NewDecoder(file)
	var country_state_city []Final_Model
	// err = decoder.Decode(&country_state_city)
	err = json.Unmarshal(countriesData, &country_state_city)
	if err != nil {

		return err.Error() + " 2"
	}
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
	return "country_state_city"
}

type State_City struct {
	State  string   `json:"state" bson:"state"`
	Cities []string `json:"cities" bson:"cities"`
}

type Final_Model struct {
	Country_Name string       `json:"country_name" bson:"country_name"`
	StateCity    []State_City `json:"state_city" bson:"state_city"`
}
