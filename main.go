package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Car Sharing API!")
	fmt.Println("Endpoint hit: Homepage")
}

func handleRequests() {
	//new instance of mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	
	userRoutes(myRouter)
	bookingRoutes(myRouter)
	parkingRoutes(myRouter)
	carRoutes(myRouter)
	reviewRoutes(myRouter)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Rest api v2 - Mux Routers")
	Users = createUsers()
	Cars = createCars()
	Parkings = createParkings()
	Bookings = createBookings()
	Reviews = createReviews()
	handleRequests()
}