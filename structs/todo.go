package structs

import "gopkg.in/mgo.v2/bson"

// Todo - Structuring todo
type Todo struct {
	ID    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title string        `json:"title" bson:"title"`
}
