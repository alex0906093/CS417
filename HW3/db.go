package main

import(
	"fmt"
	"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func addToDB(s *Student, db *mgo.Session) error{
	c := db.DB("restful").C("students")
	count, err := c.Find(bson.M{"netid": s.NetID}).Limit(1).Count()
	if err != nil{
		return err
	}
	if count > 0{
		return fmt.Errorf("student with NetID %s already exists", s.NetID)
	}
	return c.Insert(s)
}


func getFromDB(query bson.M, db *mgo.Session) ([]byte, error){
	c:= db.DB("restful").C("students")
	var results []Student
	err1 := c.Find(query).All(&results)
	if err1 != nil{
		return nil, err1
	}
	b, err := json.Marshal(results)
		if err != nil{
			fmt.Println("problem ", err)
			return nil,err
		}

	fmt.Println("Results are: ", results)
	return b,nil
}

func getAllFromDB(db *mgo.Session)([]byte,error){
	c:= db.DB("restful").C("students")
	var results []Student
	err := c.Find(nil).All(&results)
	if err != nil{
		panic(err)
	}
	b, err := json.Marshal(results)
		if err != nil{
			fmt.Println("problem ", err)
			return nil,err
		}
	fmt.Println("Results are: ", results)
	return b,nil
}