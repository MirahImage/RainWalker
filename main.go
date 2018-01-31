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
func NumberSheds(shedReq *soda.GetRequest) (uint, error) {
	return shedReq.Count()
}

// IsCovered queries the NYC sheds API for a specific address.
// It returns whether there is a shed at the given address.
func IsCovered(shedReq *soda.GetRequest, borough string, streetName string, houseNumber string) (bool, error) {
	shedReq.Format = "json"
	shedReq.Filters["borough"] = borough
	shedReq.Filters["street_name"] = streetName
	shedReq.Filters["house__"] = houseNumber
	shedReq.Query.Limit = 10

	count, err := shedReq.Count()
	return count > 0, err
}

func main() {
	shedReq := soda.NewGetRequest(shedsEndpoint, appToken)
	numSheds, err := NumberSheds(shedReq)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("There are", numSheds, "sheds in NYC")
}
