package main

import (
	"flag"
	"strings"
	"strconv"
	"bytes"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type sslice []string

func (i *sslice) String() string{
	return fmt.Sprintf("%d", *i)
}

func (i *sslice) Set(value string) error{
	fmt.Printf("%s\n", value)
	tmp := value;
	*i = append(*i, tmp)
	return nil
}

var dataPtr sslice

func main(){
	//get all of the command line arguments and parse them

	methodPtr := flag.String("method", "none","flag for the method to use")
	urlPtr := flag.String("url", "none", "flag for the URL")
	flag.Var(&dataPtr, "data", "List of data")
	yearPtr := flag.Int("year", 0, "flag for year to remove students")
	flag.Parse()
	
	fmt.Println("Printing data vars")
	for i := 0; i < len(dataPtr); i++{
		fmt.Printf("%s\n", dataPtr[i])
	}
		//Cases for method flags
	//no method flag entered
	if strings.EqualFold(*methodPtr, "none") {
		fmt.Println("Invalid method argument, program exiting")
		return
	}

	//create, add student to database
	if strings.EqualFold(*methodPtr, "Create"){
		//check flags
		if strings.EqualFold(*urlPtr, "none"){
			fmt.Println("You need to enter a URL with the -url flag")
		}
		//get data
		var s1, s2, s3, s4, s5, s6 []string
		var netID, name, major, year, grade, rating string
		if len(dataPtr) != 6 {
			fmt.Println("Invalid data input, check format")
			return
		}
		for i := 0; i < len(dataPtr); i++ {
			dataPtr[i] = strings.Replace(dataPtr[i], "’","", -1) //clear useless chars
			if i == 0{
			s1 = strings.Split(dataPtr[i],":")
				if s1[0] == "NetID"{
					netID = s1[1]
				}else{
					fmt.Println("Invalid data input, check format")
					return
				}
			}
			if i == 1{
			s2 = strings.Split(dataPtr[i],":")
				if s2[0] == "Name"{
					name = s2[1]
				}else{
					fmt.Println("Invalid data input, check format")
					return
				}
			}
			if i == 2{
			s3 = strings.Split(dataPtr[i],":")
				if s3[0] == "Major"{
					major = s3[1]
				}else{
					fmt.Println("Invalid data input, check format")
					return
				}
			}
			if i == 3{
			s4 = strings.Split(dataPtr[i],":")
				if s4[0] == "Year"{
					year = s4[1]
				}else{
					fmt.Println("Invalid data input, check format")
					return
				}
			}
			if i == 4{
			s5 = strings.Split(dataPtr[i],":")
				if s5[0] == "Grade" {
					grade = s5[1]
				}else{
					fmt.Println("Invalid data input, check format")
					return
				}
			}
			if i == 5{
			s6 = strings.Split(dataPtr[i],":")
				if(s6[0] == "Rating"){
					rating = s6[1]
				}else{
					fmt.Println("Invalid data input, check format")
					return
				}
			}
		}

		year1, err := strconv.Atoi(year)
		if err != nil{
			fmt.Println("problem converting year")
			return
		}
		grade1, err := strconv.Atoi(grade)
		if err != nil{
			fmt.Println("problem converting grade")
			return
		}
		o := bson.NewObjectId()
		qStudent := 
			Student{
				o,
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
		fmt.Println("Response Status:", resp.Status)
		fmt.Println("Response Headers", resp.Header)
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