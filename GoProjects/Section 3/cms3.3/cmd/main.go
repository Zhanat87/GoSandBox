package main

import (
	"net/http"

	cms "goprojects/cms3.3"
)

func main() {
	http.HandleFunc("/", cms.ServeIndex)
	http.HandleFunc("/page/", cms.ServePage)
	http.HandleFunc("/post/", cms.ServePost)
	http.HandleFunc("/new", cms.HandleNew)
	http.ListenAndServe(":3000", nil)
}