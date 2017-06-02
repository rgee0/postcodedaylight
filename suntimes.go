package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type suntimesType struct {
	Status string `json:"status"`
	TimeOf struct {
		Sunrise   string `json:"sunrise"`
		Sunset    string `json:"sunset"`
		DayLength string `json:"day_length"`
	} `json:"results"`
}

func requestSuntimes(longitude float64, latitude float64) suntimesType {
	client := http.Client{}
	res, err := client.Get("https://api.sunrise-sunset.org/json?lat=" + strconv.FormatFloat(latitude, 'f', -1, 32) + "&lng=" + strconv.FormatFloat(longitude, 'f', -1, 32))
	if err != nil {
		log.Fatalln("Unable to reach sunrise-sunset.org endpoint.")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("Unable to parse response from server.")
	}

	suntimes := suntimesType{}
	json.Unmarshal(body, &suntimes)
	return suntimes
}
