package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
	fmt.Println("Endpoint hit: Homepage")
}

func handleRequests() {
	//new instance of mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/api/users", returnAllUsers)
	myRouter.HandleFunc("/api/user", createNewUser).Methods("POST")
	myRouter.HandleFunc("/api/user/{id}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/api/user/{id}", updateUser).Methods("PUT")
	myRouter.HandleFunc("/api/user/{id}", returnSingleUser)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Rest api v2 - Mux Routers")
	Users = []User{
		User{Id: "1", FirstName: "John", LastName: "Doe", Email: "john@gmail.com", PhoneNumber: "07644655293"},
		User{Id: "2", FirstName: "Melanie", LastName: "Smith", Email: "ms10101@aol.com", PhoneNumber: "1474558011"},
		User{Id: "3", FirstName: "Leonardo", LastName: "DiCaprio", Email: "leo@dicap.com", PhoneNumber: "2123459862"},
	}
	handleRequests()
}