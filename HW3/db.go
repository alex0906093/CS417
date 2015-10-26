package main

import(
	"fmt"
	//"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func addToDB(s *Student, db *mgo.Session) error{
	c := db.DB("test").C("students")
	count, err := c.Find(bson.M{"netid": s.NetID}).Limit(1).Count()
	if err != nil{
		return err
	}
	if count > 0{
		return fmt.Errorf("student with NetID %s already exists", s.NetID)
	}
	return c.Insert(s)
}