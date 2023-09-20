package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movie []Movie // movie is the array of the struct movie

func main() {
	r := mux.NewRouter()

	movie = append(movie, Movie{Id: "1", Isbn: "438227", Title: "Movie one", Director: &Director{Firstname: "Hitesh", Lastname: "Arora"}})
	movie = append(movie, Movie{Id: "1", Isbn: "45455", Title: "Movie Two", Director: &Director{Firstname: "John", Lastname: "Stell"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting the server at 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
