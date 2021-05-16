package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Car struct {
	Id string `json:"id"`
	Model string `json:"model"`
	Color string `json:"color"`
	Plate string `json:"plate"`
	User string `json:"user"`
}

var Cars []Car

func returnAllCars(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllCars")
	json.NewEncoder(w).Encode(Cars)
}

func returnSingleCar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnSingleCar")
	key := getParams(r, "id")
	for _, car := range Cars {
		if car.Id == key {
			json.NewEncoder(w).Encode(car)
		}
	}
}

func createNewCar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createNewCar")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var car Car
	json.Unmarshal(reqBody, &car)
	Cars = append(Cars, car)
	json.NewEncoder(w).Encode(car)
}

func updateCar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: updateCar")
	id := getParams(r, "id")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var car Car
	json.Unmarshal(reqBody, &car)
	for i, u := range Cars {
		if u.Id == id {
			Cars[i] = car
		}
	}
}

func deleteCar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteCars")
	id := getParams(r, "id")
	for i, u := range Cars {
		if u.Id == id {
			Cars = append(Cars[:i], Cars[i+1:]...)
		}
	}
}