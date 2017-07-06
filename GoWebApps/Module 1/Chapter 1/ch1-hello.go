package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"syscall"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/test", TestHandler)
	http.Handle("/", router)
	fmt.Println("Everything is set up!")

}

//Some comment
