package main

import (
	"flag"
	"strings"
	"net/http"
	"strconv"
	"os"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"fmt"
)


func main(){
	//get all of the command line arguments and parse them

	methodPtr := flag.String("method", "none","flag for the method to use")
	urlPtr := flag.String("url", "none", "flag for the URL")
	dataPtr := flag.String("data", "none", "be used")
	yearPtr := flag.Int("year", 0, "flag for year to remove students")
	flag.Parse()

		//Cases for method flags
	
	//no method flag entered
	if strings.EqualFold(*methodPtr, "none") {
		fmt.Println("Invalid method argument, program exiting")
		return
	}

	//create, add student to database
	if strings.EqualFold(*methodPtr, "Create"){
		//check flags
		a1 := os.Args[3]
		a2 := os.Args[4]
		a3 := os.Args[5]
		a4 := os.Args[6]
		a5 := os.Args[7]
		a6 := os.Args[8]
		//fmt.Println(os.Args)
		if strings.EqualFold(*dataPtr, "none") {
			fmt.Println("You need to enter data with the -data flag, rerun with data")
			return
		}
		if strings.EqualFold(*urlPtr, "none"){
			fmt.Println("You need to enter a URL with the -url flag")
		}
		netID := fixString(a1)
		name := fixString(a2)
		major := fixString(a3)
		year := fixString(a4)
		grade := fixString(a5)
		rating := fixString(a6)
		grade = grade[0:2]
		year = year[0:4]
		fmt.Println(netID)
		/*
		fmt.Println("name is ", name)
		fmt.Println("major is ", major)
		fmt.Println("year is ", year)
		fmt.Println("grade is ", grade)
		fmt.Println("rating is ", rating)
		*/
		year1, err := strconv.Atoi(year)
		if err != nil{
			fmt.Println("problem1", err)
			
		}
		grade1, err := strconv.Atoi(grade)
		if err != nil{
			fmt.Println("problem2", err)
			
		}

		qStudent := 
			Student{
				netID,
				name,
				major,
				year1,
				grade1,
				rating,
		}
		fmt.Println("qStudent: ", qStudent.Grade)
		b, err := json.Marshal(qStudent)
		if err != nil{
			fmt.Println("problem ", err)
		}
		//fmt.Println("b is ", b)

		// create http request
		req, err := http.NewRequest("POST", *urlPtr, bytes.NewBuffer(b))
		req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Set("Content-Type", "application/json;charset=UTF-8")
		
		//send http request
		client := &http.Client{}
		resp, err := client.Do(req)
	
		if err != nil {
			panic(err)
		}
		
		defer resp.Body.Close()
		//fmt.Println("Response Status:", resp.Status)
		//fmt.Println("Response Headers", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("Response Body", string(body))
	}
	
	//remove, delete student from DB
	if strings.EqualFold(*methodPtr, "remove") {
		if *yearPtr == 0{
			fmt.Println("no year flag input, run again with -year flag")
			return
		}
	}

	//update, set student grades in DB 
	if strings.EqualFold(*methodPtr, "update"){

	}
	//list, for LIST Operation and GET operation
	if strings.EqualFold(*methodPtr, "list") {

	}
	
}

//helper method get rid of single quote
func fixString(s string) string{
	ns := strings.Split(s, ":")[1]
	ns = strings.Trim(ns, " ")
	l := len(ns)
	if l > 0 {
	ns = ns[:l-1]
	}
	return ns 

}