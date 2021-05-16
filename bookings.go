package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Booking struct {
	Id string `json:"id"`
	StartDate time.Time `json:"start_date"`
	EndDate time.Time `json:"end_date"`
	StartTime time.Time `json:"start_time"`
	EndTime time.Time `json:"end_time"`
	Car string `json:"car"`
	Parking string `json:"parking"`
}

var Bookings []Booking

func returnAllBookings(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: returnAllBookings")
	json.NewEncoder(w).Encode(Bookings)
}

func returnSingleBooking(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnSingleBooking")
	key := getParams(r, "id")
	for _, booking := range Bookings {
		if booking.Id == key {
			json.NewEncoder(w).Encode(booking)
		}
	}
}

func createNewBooking(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createNewBooking")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var booking Booking
	json.Unmarshal(reqBody, &booking)
	Bookings = append(Bookings, booking)
	json.NewEncoder(w).Encode(booking)
}

func updateBooking(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: updateBooking")
	id := getParams(r, "id")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var booking Booking
	json.Unmarshal(reqBody, &booking)
	for i, u := range Bookings {
		if u.Id == id {
			Bookings[i] = booking
		}
	}
}

func deleteBooking(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteBooking")
	id := getParams(r, "id")
	for i, u := range Bookings {
		if u.Id == id {
			Bookings = append(Bookings[:i], Bookings[i+1:]...)
		}
	}
}