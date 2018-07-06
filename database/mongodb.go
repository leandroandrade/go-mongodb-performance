package database

import (
	"gopkg.in/mgo.v2"
	"log"
)

type MongoDatabase struct {
	session *mgo.Session
}

func NewMongoInstance() *MongoDatabase {
	session, err := mgo.Dial("localhost:27017")
	session.SetMode(mgo.Monotonic, true)
	if err != nil {
		log.Fatal(err)
	}

	return &MongoDatabase{session: session}
}

func (s *MongoDatabase) Copy() *mgo.Session {
	return s.session.Copy()
}

func (s *MongoDatabase) Clone() *mgo.Session {
	return s.session.Clone()
}

func (s *MongoDatabase) Get() *mgo.Session {
	return s.session
}

func (s *MongoDatabase) Close() {
	if s.session != nil {
		s.session.Close()
	}
}
