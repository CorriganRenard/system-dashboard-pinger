package models

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	// IpAddress represents the structure of our resource
	IpAddress struct {
		Id          bson.ObjectId `json:"id" bson:"_id"`
		Ip          string        `json:"ip" bson:"ip"`
		LastPing    string        `json:"lastPing" bson:"lastPing"`
		PingOk      bool          `json:"pingOk" bson:"pingOk"`
		Description string        `json:"description" bson:"description"`
		Location    string        `json:"location" bson:"location"`
	}
)
