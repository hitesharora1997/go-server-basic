package main

import "net/http"

func main() {
	fileserver := http.FileServer(http.Dir("./static")) // telling the golang to check the in the static folder
	http.Handle("/", fileserver)                        // setting up a route for "/" and pointing it at our file server

}
