package main

import (
    //"encoding/json"
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)
//helper function to retreive http stuff
/*
func (Ls *LinkShortnerAPI) UrlShow(w http.ResponseWriter, r *http.Request) {
    //retrieve the variable from the request
    vars := mux.Vars(r)
}
*/
//handle the post function
func postHandle(w http.ResponseWriter, r *http.Request) {

	// read from r
	fmt.Println("Just show we're getting here")
    vars := mux.Vars(r)
    fmt.Println("vars is", vars)
	//write to w
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

}


//handle the get function
func getHandle(w http.ResponseWriter, r *http.Request) {

	// read from r 



	//write to w

}

//handle the delete function
func deleteHandle(w http.ResponseWriter, r *http.Request) {

	// read from r 



	//write to w
}

//handle the update function
func updateHandle(w http.ResponseWriter, r *http.Request) {

	// read from r 



	//write to w
}

//handle the list function
func listHandle(w http.ResponseWriter, r *http.Request) {

	// read from r 



	//write to w
}
