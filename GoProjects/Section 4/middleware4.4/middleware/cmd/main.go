package main

import (
	"net/http"

	"golang.org/x/net/context"
	"goprojects/middleware"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func panicker(w http.ResponseWriter, r *http.Request) {
	panic(middleware.ErrInvalidID)
}

func withContext(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	bar := ctx.Value("foo")
	w.Write((bar.([]byte)))
}

func main() {
	logger := middleware.CreateLogger("section4")

	http.Handle("/", middleware.Time(logger, hello))
	http.Handle("/panic", middleware.Recover(panicker))
	http.Handle("/context", middleware.PassContext(withContext))
	http.ListenAndServe(":3000", nil)
}
