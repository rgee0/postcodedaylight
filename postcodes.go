package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type postcodeType struct {
	Status  int `json:"status"`
	LongLat struct {
		Longitude float64 `json:"longitude"`
		Latitude  float64 `json:"latitude"`
	} `json:"result"`
}

func requestPostcode(postcode string) postcodeType {
	client := http.Client{}
	res, err := client.Get("https://api.postcodes.io/postcodes/" + postcode)
	if err != nil {
		log.Fatalln("Unable to reach postcodes.io endpoint.")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("Unable to parse response from server.")
	}

	postcodeLocation := postcodeType{}
	json.Unmarshal(body, &postcodeLocation)
	return postcodeLocation
}
