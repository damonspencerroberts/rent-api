package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	Id string `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

var Users []User

func returnAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllUsers")
	json.NewEncoder(w).Encode(Users)
}

func returnSingleUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnSingleUser")
	key := getParams(r)
	for _, user := range Users {
		if user.Id == key {
			json.NewEncoder(w).Encode(user)
		}
	}
}

func createNewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createNewUser")

	// get the body of our POST request
  // unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(reqBody, &user)
	// update our global Articles array to include
  // our new Article
	Users = append(Users, user)

	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: updateUser")
	id := getParams(r)
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(reqBody, &user)

	for i, u := range Users {
		if u.Id == id {
			Users[i] = user
		}
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteUser")
	id := getParams(r)
	for i, u := range Users {
		//for all the users find the one looking for and remove
		if u.Id == id {
			Users = append(Users[:i], Users[i+1:]...)
		}
	}
}
