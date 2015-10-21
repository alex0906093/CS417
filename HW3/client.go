package main

import (
	"flag"
	"strings"
	"net/http"
	//"net/url"
	"bytes"
	"io/ioutil"
	"fmt"
)

func main(){
	//get all of the command line arguments and parse them
	methodPtr := flag.String("method", "none","flag for the method to use")
	urlPtr := flag.String("url", "none", "flag for the URL")
	dataPtr := flag.String("data", "none", "flag for the data")
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
		if strings.EqualFold(*dataPtr, "none") {
			fmt.Println("You need to enter data with the -data flag, rerun with data")
			return
		}
		if strings.EqualFold(*urlPtr, "none"){
			fmt.Println("You need to enter a URL with the -url flag")
		}
		// create http request
		var query = []byte(*dataPtr)
		/*var urlSend, err = url.Parse(*urlPtr)
		if err != nil {
			fmt.Println("Problem with URL")
			return
		}*/
		req, err := http.NewRequest("POST", *urlPtr, bytes.NewBuffer(query))
		fmt.Println("req method is", req.Method)
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