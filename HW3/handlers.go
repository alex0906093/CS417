package main

import (
    "encoding/json"
    "fmt"
    //"bytes"
    "io"
    "io/ioutil"
    "net/http"
    "github.com/gorilla/mux"
    "gopkg.in/mgo.v2"
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
    var student Student
	// read from r
	fmt.Println("Just show we're getting here")
    vars := mux.Vars(r)
    fmt.Println("vars is", vars)
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil{
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &student); err != nil{
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
       }
    }
    fmt.Println("Json Data ", student)
    dbsess, err := mgo.Dial("localhost:27017")
    if err != nil {
        fmt.Println("DB connection error")
        panic(err)
    }
    err1 := addToDB(&student, dbsess)
    fmt.Println(err1)
    dbsess.Close()

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
