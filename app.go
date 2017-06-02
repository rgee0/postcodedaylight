package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func evalTime(inNum string, unit string) string {

	i, _ := strconv.Atoi(inNum)
	var outVal string

	if i == 1 {
		outVal = inNum + " " + unit
	} else {
		outVal = inNum + " " + unit + "s"
	}
	return outVal
}

func splitTime(intime string) string {
	parts := strings.Split(intime, ":")
	return evalTime(parts[0], "hour") + ", " + evalTime(parts[1], "minute") + " & " + evalTime(parts[2], "second")
}

func stripSpaces(input string) string {
	//split on a new line as stdin appends \n
	parts := strings.Split(input, "\n")
	//return te part befoew the first newline with all spaces removed
	return strings.Replace(parts[0], " ", "", -1)
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

	fmt.Printf("The duration of daylight at %s was %s.\n", strings.ToUpper(postcode), splitTime(suntimes.TimeOf.DayLength))

}
