package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)



func postHandle(w http.ResponseWriter, r *http.Request) {

	// read from r
	 



	//write to w
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(todos); err != nil {
        panic(err)
    }
}



func getHandle(w http.ResponseWriter, r *http.Request) {

	// read from r 



	//write to w
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(todos); err != nil {
        panic(err)
    }

}


func deleteHandle(w http.ResponseWriter, r *http.Request) {

	// read from r 



	//write to w
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(todos); err != nil {
        panic(err)
    }

}


func updateHandle(w http.ResponseWriter, r *http.Request) {

	// read from r 



	//write to w
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(todos); err != nil {
        panic(err)
    }
}


func listHandle(w http.ResponseWriter, r *http.Request) {

	// read from r 



	//write to w
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(todos); err != nil {
        panic(err)
    }

}
