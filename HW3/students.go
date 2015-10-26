package main

import "gopkg.in/mgo.v2/bson"

type Student struct{
	Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
	NetID string `bson:"netid" json:"netid"`
	Name string `bson:"name" json: "name"`
	Major string `bson:"major" json: "major"`
	Year int `bson:"year" json: "year"`
	Grade int `bson:"grade" json: "grade"`
	Rating string `bson:"rating" json: "rating"`
}

type Students []Student