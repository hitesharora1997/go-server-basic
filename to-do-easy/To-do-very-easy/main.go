package main

import (
	"time"

	"github.com/thedevsaddam/renderer"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var rnd *renderer.Render
var db *mgo.Database // mongo db

// My custom struct
const (
	hostname       string = "localhost:27017"
	dbName         string = "demo_todo"
	collectionName string = "todo"
	port           string = "9000"
)

type ( // meed to look on my own
	todoModel struct {
		ID        bson.ObjectId `"bson:_id,omitempty"` // use of back ticks here
		Title     string        `"bson: title"`
		Completed bool          `"bson:completed"`
		CreatedAt time.Time     `"bson:createdAt"`
	}
)

func main() {
	println("hello")
}
