package main

import (
	"flag"
	"net/http"
	"os"
)

func main() {
	//dir, _ := os.Getwd() //Get Working Directionary , first var is dir ,second is error
	//Just because we dont care about the error this case

	//http.ListenAndServe(":3000",http.FileServer(http.Dir(dir)))
	var dir string
	port := flag.String("port", "3000", "port to serve HTTP on")
	path := flag.String("path", "", "port to the server")
	flag.Parse()

	if *path == "" {
		dir, _ = os.Getwd()
	} else {
		dir = *path //Only use Porinter * for the flag package
	}

	http.ListenAndServe(":"+*port, http.FileServer(http.Dir(dir)))
}
