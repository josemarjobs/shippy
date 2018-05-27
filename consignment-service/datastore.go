package main

import (
	"log"

	"gopkg.in/mgo.v2"
)

func CreateSession(host string) (*mgo.Session, error) {
	log.Print("Connecting to MongoDB... ")
	session, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}
	log.Println("Connected.")

	session.SetMode(mgo.Monotonic, true)

	return session, nil
}
