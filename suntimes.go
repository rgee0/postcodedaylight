package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type suntimesType struct {
	Status string `json:"status"`
	TimeOf struct {
		Sunrise   string `json:"sunrise"`
		Sunset    string `json:"sunset"`
		DayLength int    `json:"day_length"`
	} `json:"results"`
}

func makeDate(inDateStr string) time.Time {
	//https://golang.org/src/time/format.go tells us how to specify date formats
	const apiForm = "2006-01-02T15:04:05+00:00"
	//convert the date string to a time var
	outDate, _ := time.Parse(apiForm, inDateStr)
	return outDate
}

func daylength(uptime string, downtime string) time.Duration {
	upTime := makeDate(uptime)
	downTime := makeDate(downtime)
	return downTime.Sub(upTime)
}

func requestSuntimes(longitude float64, latitude float64) suntimesType {
	client := http.Client{}
	uri := "https://api.sunrise-sunset.org/json?lat=" + strconv.FormatFloat(latitude, 'f', -1, 32) + "&lng=" + strconv.FormatFloat(longitude, 'f', -1, 32) + "&formatted=0"
	res, err := client.Get(uri)
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
