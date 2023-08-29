package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/thedevsaddam/renderer"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var rnd *renderer.Render
var db mgo.Database // mongo db

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
	todo struct {
		ID        bson.ObjectId `"json:_id"` // use of back ticks here
		Title     string        `"json: title"`
		Completed bool          `"json:completed"`
		CreatedAt time.Time     `"json:created_at"`
	}
)

// This function is there to interact with the database
func init() {
	rnd = renderer.New()
	sess, err := mgo.Dial(hostname)
	checkErr(err)
	sess.SetMode(mgo.Monotonic,true)
	db = *sess.DB(dbName)

}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)	"gopkg.in/mgo.v2"
 
	r.Mount("/todo", todoHandler())

}

func todoHandler() http.Handler {
	rg := chi.NewRouter()
}
