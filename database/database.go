package database

import (
	"log"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"
)

var mongo, err = mgo.Dial("localhost")

// GetConnection - Return mongo connection
func GetConnection() *mgo.Session {
	if err != nil {
		log.Fatal(err)
		mongo.Close()
	} else {
		mongo.SetMode(mgo.Monotonic, true)
	}
	return mongo
}

// GetBID - Return new Bson object id
func GetBID() bson.ObjectId {
	bid := []byte{80, 36, 13, 102, 47, 85, 184, 17, 20, 0, 0, 1}
	return bson.ObjectId(bid)
}
