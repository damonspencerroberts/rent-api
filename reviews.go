package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Review struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Stars string `json:"stars"`
	Parking string `json:"parking"`
}

var Reviews []Review

func returnAllReviews(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: returnAllReviews")
	key := getParams(r, "parking_id")
	var curReviews []Review
	for _, review := range Reviews {
		if review.Parking == key {
			curReviews = append(curReviews, review)
		}
	}
	json.NewEncoder(w).Encode(curReviews)
}

func createNewReview(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: createNewReview")
	reqBody, _ := ioutil.ReadAll(r.Body)
	key := getParams(r, "parking_id")
	var curReviews []Review
	for i, review := range Reviews {
		if review.Parking == key {
			curReviews = append(curReviews, review)
			Reviews = append(Reviews[:i], Reviews[i+1:]...)
		}
	}
	var review Review
	json.Unmarshal(reqBody, &review)
	curReviews = append(curReviews, review)
	Reviews = append(Reviews, curReviews...)
	json.NewEncoder(w).Encode(review)
}

func updateReview(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: updateReview")
	parkingId := getParams(r, "parking_id")
	key := getParams(r, "id")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var review Review
	json.Unmarshal(reqBody, &review)
	for i, rev := range Reviews {
		if rev.Parking == parkingId {
			if rev.Id == key {
				Reviews[i] = review
			}
		}
	}
}

func deleteReview(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: deleteReview")
	parkingId := getParams(r, "parking_id")
	key := getParams(r, "id")
	for i, review := range Reviews {
		if review.Parking == parkingId {
			if review.Id == key {
				Reviews = append(Reviews[:i], Reviews[i+1:]...)
			}
		}
	} 
}

func returnSingleReview(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: returnSingleReview")
	parkingId := getParams(r, "parking_id")
	key := getParams(r, "id")
	for _, review := range Reviews {
		if review.Parking == parkingId {
			if review.Id == key {
				json.NewEncoder(w).Encode(review)
			}
		}
	} 
}

// func deleteCar(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint Hit: deleteCars")
// 	id := getParams(r, "id")
// 	for i, u := range Cars {
// 		if u.Id == id {
// 			Cars = append(Cars[:i], Cars[i+1:]...)
// 		}
// 	}
// }