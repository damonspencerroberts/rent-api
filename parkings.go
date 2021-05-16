package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Parking struct {
	Id string `json:"id"`
	Address string `json:"address"`
	NumberOfCars string `json:"number_of_cars"`
	Images []string `json:"images"`
	Price string `json:"price"`
	Desc string `json:"description"`
	Ammenities []string `json:"ammenities"`
	Latitude string `json:"latitude"`
	Longitude string `json:"longitude"`
	User string `json:"user"`
}

var Parkings []Parking

func returnAllParkings(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: returnAllParkings")
	json.NewEncoder(w).Encode(Parkings)
}

func returnSingleParking(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnSingleParking")
	key := getParams(r, "id")
	for _, parking := range Parkings {
		if parking.Id == key {
			json.NewEncoder(w).Encode(parking)
		}
	}
}

func createNewParking(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createNewParking")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var parking Parking
	json.Unmarshal(reqBody, &parking)
	Parkings = append(Parkings, parking)
	json.NewEncoder(w).Encode(parking)
}

func updateParking(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: updateParking")
	id := getParams(r, "id")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var parking Parking
	json.Unmarshal(reqBody, &parking)
	for i, u := range Parkings {
		if u.Id == id {
			Parkings[i] = parking
		}
	}
}

func deleteParking(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteParking")
	id := getParams(r, "id")
	for i, u := range Parkings {
		if u.Id == id {
			Parkings = append(Parkings[:i], Parkings[i+1:]...)
		}
	}
}