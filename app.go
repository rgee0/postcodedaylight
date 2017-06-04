package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func stripSpaces(inStr string) string {
	//split on a new line as stdin appends \n
	parts := strings.Split(inStr, "\n")
	//return te part befoew the first newline with all spaces removed
	return strings.Replace(parts[0], " ", "", -1)
}

func daylightRemaining(downtime string) string {
	//remove the sub-second values
	t := time.Now().UTC().Round(time.Second)

	//convert string representation to a time
	downTime := makeDate(downtime)
	//initialise with default value to save elsing
	outVal := "Sunset has passed"

	//if there is remaining daylight
	if t.Before(downTime) {
		//over-ride the default with remaining duration
		outVal = downTime.Sub(t).String() + " until sunset"
	}
	return outVal
}

func main() {

	input, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		log.Fatal("Unable to read standard input:", err)
	}

	postcode := string(input)

	//Should do the stripping before the validation otherwise
	// a series of spaces would pass the following len test
	postcode = stripSpaces(postcode)

	if len(postcode) == 0 {
		log.Fatalln("A UK postcode is required.\n")
	}

	//go get the postcode information
	postcodeInfo := requestPostcode(postcode)

	//check whether the postcode request returned a valid status
	if postcodeInfo.Status != 200 {
		fmt.Printf("The postcode entered was not found.\n")
		os.Exit(0)
	}

	//if there is a valid response there is more work to do
	suntimes := requestSuntimes(postcodeInfo.LongLat.Longitude, postcodeInfo.LongLat.Latitude)

	if suntimes.Status != "OK" {
		fmt.Printf("Sun times for the entered postcode were not found.\n")
		os.Exit(0)
	}
	//Everything seems to be in order so output some results
	fmt.Printf("Duration of daylight today at %s : %s (%s).\n", strings.ToUpper(postcode), daylength(suntimes.TimeOf.Sunrise, suntimes.TimeOf.Sunset), daylightRemaining(suntimes.TimeOf.Sunset))

}
