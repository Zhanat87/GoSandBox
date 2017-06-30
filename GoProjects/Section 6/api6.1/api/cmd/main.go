package main

import (
	"net/http"

	"goprojects/api"
)

func main() {
	http.HandleFunc("/", api.Doc)
	http.HandleFunc("/newpage", api.CreatePage)
	http.HandleFunc("/pages", api.AllPages)
	http.HandleFunc("/pages/", api.GetPage)

	http.ListenAndServe(":3000", nil)
}
