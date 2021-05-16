package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func getParams(r *http.Request) string {
	//this gets the params
	//unique id from the params
	vars := mux.Vars(r)
	key := vars["id"]
	return key
}