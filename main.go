package main

import (
	"fmt"
	"log"

	soda "github.com/SebastiaanKlippert/go-soda"
)

const shedsEndpoint = "https://data.cityofnewyork.us/resource/2jy7-cddj"
const appToken = "Hl8EuekONWAZmQj7osuAIKPTg"

// NumberSheds queries the NYC sheds API.
// It returns the number of sheds and any errors.
func NumberSheds() (uint, error) {
	nyData := soda.NewGetRequest(shedsEndpoint, appToken)
	return nyData.Count()
}

func main() {
	numSheds, err := NumberSheds()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("There are", numSheds, "sheds in NYC")
}
