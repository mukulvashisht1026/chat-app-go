package datastore

import "github.com/golang/protobuf/ptypes/timestamp"

type Group struct {
	Id string
	Type string 
	Name string
}

type Message struct {
	Type string
	Value string
	Sender string
	Timestamp timestamp.Timestamp
	GroupId string
}

type User struct {
	Id string
	Name string
	MobileNumber string
}

type GroupsToUserMapping struct {
	GroupId string
	MemberId string
}


