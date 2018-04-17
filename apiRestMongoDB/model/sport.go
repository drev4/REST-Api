package model

import "gopkg.in/mgo.v2/bson"

//Sport Objeto deporte
type Sport struct {
	ID   bson.ObjectId `json:"id,omitempty"`
	Name string        `json:"firstname,omitempty"`

}

