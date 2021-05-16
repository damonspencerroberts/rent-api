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

	// Users
	myRouter.HandleFunc("/api/users", returnAllUsers)
	myRouter.HandleFunc("/api/user", createNewUser).Methods("POST")
	myRouter.HandleFunc("/api/user/{id}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/api/user/{id}", updateUser).Methods("PUT")
	myRouter.HandleFunc("/api/user/{id}", returnSingleUser)

	// Cars
	myRouter.HandleFunc("/api/cars", returnAllCars)
	myRouter.HandleFunc("/api/car", createNewCar).Methods("POST")
	myRouter.HandleFunc("/api/car/{id}", deleteCar).Methods("DELETE")
	myRouter.HandleFunc("/api/car/{id}", updateCar).Methods("PUT")
	myRouter.HandleFunc("/api/car/{id}", returnSingleCar)

	// Parkings
	myRouter.HandleFunc("/api/parkings", returnAllParkings)
	myRouter.HandleFunc("/api/parking", createNewParking).Methods("POST")
	myRouter.HandleFunc("/api/parking/{id}", deleteParking).Methods("DELETE")
	myRouter.HandleFunc("/api/parking/{id}", updateParking).Methods("PUT")
	myRouter.HandleFunc("/api/parking/{id}", returnSingleParking)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Rest api v2 - Mux Routers")
	Users = handleUsers()
	Cars = handleCars()
	Parkings = handleParkings()
	handleRequests()
}