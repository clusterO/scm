package main

import (
	service "scm/db/cmd/service"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	var MgoSession *mgo.Session
	println(MgoSession)

	defer MgoSession.Close()

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	MgoSession = session
	println(MgoSession)

	c := MgoSession.DB("test").C("people")
	println(c)

	c.Insert(bson.M{"a": 1, "b": true})

	service.Run()
}
