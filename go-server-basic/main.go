package main

import "net/http"

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
}
