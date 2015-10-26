package main

import (
    "encoding/json"
    "fmt"
    //"bytes"
    "io"
    "strconv"
    "io/ioutil"
    "net/http"
    //"github.com/gorilla/mux"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
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
	//fmt.Println("Just show we're getting here")
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
    fmt.Println("got to GET handle from request")
    name := r.URL.Query().Get("name")
    netid := r.URL.Query().Get("netid")
    major := r.URL.Query().Get("major")
    year := r.URL.Query().Get("year")
	grade := r.URL.Query().Get("grade")
    rating := r.URL.Query().Get("rating")
    if len(name) == 0 && len(major) == 0 && len(year) == 0 && len(grade) == 0 && len(rating) == 0 {
        fmt.Println("No valid arguments sent in request")
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422)
        return
    }
    query := bson.M {}
    query["$and"] = []bson.M{}
    if len(name) != 0{
        query["$and"] = append(query["$and"].([]bson.M), bson.M{"name": name})
    }
    if len(netid) != 0{
        query["$and"] = append(query["$and"].([]bson.M), bson.M{"netid": netid})
    }
    if len(major) != 0{
        query["$and"] = append(query["$and"].([]bson.M), bson.M{"major": major})
    }
    if len(rating) != 0{
        query["$and"] = append(query["$and"].([]bson.M), bson.M{"rating": rating})
    }
    if len(year) != 0{
        yInt, err := strconv.Atoi(year)
        if err != nil{
            fmt.Println("Problem with year, must be an integer")
        }
        query["$and"] = append(query["$and"].([]bson.M), bson.M{"year":yInt})
    }
    if len(grade) != 0{
        gInt, err := strconv.Atoi(grade)
        if err != nil{
            fmt.Println("Problem with grade, must be an integer")
        }
        query["$and"] = append(query["$and"].([]bson.M), bson.M{"grade": gInt})
    }

    dbsess, err := mgo.Dial("localhost:27017")
    if err != nil {
        fmt.Println("DB connection error")
        panic(err)
    }
    b, err := getFromDB(query, dbsess)
    if(err != nil){
        fmt.Println(err)
    }
    dbsess.Close()
    //write to w
    w.Header().Set("Content-Type", "application/json")
    w.Write(b)

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
    fmt.Println("got to LIST handle from request")
    dbsess, err := mgo.Dial("localhost:27017")
    if err != nil {
        fmt.Println("DB connection error")
        panic(err)
    }
    b, err := getAllFromDB(dbsess)
    if(err != nil){
        fmt.Println("error: ", err)
    }
    dbsess.Close()


	//write to w
    w.Header().Set("Content-Type", "application/json")
    w.Write(b)
}
