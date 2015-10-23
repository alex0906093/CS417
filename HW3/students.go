package main

import "gopkg.in/mgo.v2/bson"

type Student struct{
	Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
	NetID string `json:"id"`
	Name string `json: "name"`
	Major string `json: "major"`
	Year int `json: "year"`
	Grade int `json: "grade"`
	Rating string `json: "rating"`
}

type Students []Student