package main

import (
	"net/http"

	"goprojects/middleware"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func panicker(w http.ResponseWriter, r *http.Request) {
	panic(middleware.ErrInvalidID)
}

func main() {
	logger := middleware.CreateLogger("section4")

	http.Handle("/", middleware.Time(logger, hello))
	http.Handle("/panic", middleware.Recover(panicker))
	http.ListenAndServe(":3000", nil)
}
