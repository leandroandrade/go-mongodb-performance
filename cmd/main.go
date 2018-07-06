package main

import (
	"gopkg.in/mgo.v2"
	"log"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	//mongoDB()
	createPassw()

}

func createPassw() {
	password := []byte("123456")
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(hash))

	err = bcrypt.CompareHashAndPassword(hash, password)
	if err != nil {
		log.Fatal(err)
	}
}

func mongoDB() {
	session, err := mgo.Dial("localhost:27017")
	session.SetMode(mgo.Monotonic, true)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	user := User{
		ID:   uuid.New().String(),
		Name: "Leandro",
		Age:  29,
	}
	collection := session.DB("golang").C("people")
	err = collection.EnsureIndex(userIndex())
	err = collection.Insert(user)
	if err != nil {
		log.Fatal(err)
	}
	userReturn := User{}
	err = collection.Find(bson.M{"age": 29}).One(&userReturn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(userReturn)
}

func userIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

type User struct {
	ID       string
	Name     string
	Age      int
	Password string
}
