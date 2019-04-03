package models

import "gopkg.in/mgo.v2/bson"

type Movie struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	CoverImage  string        `bson:"cover_image" json:"cover_image"`
	Description string        `bson:"description" json:"description"`
}

type User struct {
	ID          	bson.ObjectId `bson:"_id" json:"id"`
	Name        	string        `bson:"name" json:"name"`
	ProfileImage  	string        `bson:"profile_image" json:"profile_image"`
	PhoneNumber 	string        `bson:"phone_number" json:"phone_number"`
	Password		string        `bson:"password" json:"password,omitempty"`
}
