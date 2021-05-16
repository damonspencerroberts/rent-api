package main

import "github.com/gorilla/mux"

func userRoutes(r *mux.Router) {
	r.HandleFunc("/api/users", returnAllUsers)
	r.HandleFunc("/api/user", createNewUser).Methods("POST")
	r.HandleFunc("/api/user/{id}", deleteUser).Methods("DELETE")
	r.HandleFunc("/api/user/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/api/user/{id}", returnSingleUser)
}

func parkingRoutes(r *mux.Router) {
	r.HandleFunc("/api/parkings", returnAllParkings)
	r.HandleFunc("/api/parking", createNewParking).Methods("POST")
	r.HandleFunc("/api/parking/{id}", deleteParking).Methods("DELETE")
	r.HandleFunc("/api/parking/{id}", updateParking).Methods("PUT")
	r.HandleFunc("/api/parking/{id}", returnSingleParking)
}

func bookingRoutes(r *mux.Router) {
	r.HandleFunc("/api/bookings", returnAllBookings)
	r.HandleFunc("/api/booking", createNewBooking).Methods("POST")
	r.HandleFunc("/api/booking/{id}", deleteBooking).Methods("DELETE")
	r.HandleFunc("/api/booking/{id}", updateBooking).Methods("PUT")
	r.HandleFunc("/api/booking/{id}", returnSingleBooking)
}

func reviewRoutes(r *mux.Router) {
	r.HandleFunc("/api/parking/{parking_id}/reviews", returnAllReviews)
	r.HandleFunc("/api/parking/{parking_id}/review", createNewReview).Methods("POST")
	r.HandleFunc("/api/parking/{parking_id}/review/{id}", deleteReview).Methods("DELETE")
	r.HandleFunc("/api/parking/{parking_id}/review/{id}", updateReview).Methods("PUT")
	r.HandleFunc("/api/parking/{parking_id}/review/{id}", returnSingleReview)
}

func carRoutes(r *mux.Router) {
	r.HandleFunc("/api/cars", returnAllCars)
	r.HandleFunc("/api/car", createNewCar).Methods("POST")
	r.HandleFunc("/api/car/{id}", deleteCar).Methods("DELETE")
	r.HandleFunc("/api/car/{id}", updateCar).Methods("PUT")
	r.HandleFunc("/api/car/{id}", returnSingleCar)
}