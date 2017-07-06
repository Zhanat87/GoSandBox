package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*"))

const SSLport = ":444"

func SocialNetwork(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got a request")
	templates.ExecuteTemplate(w, "socialnetwork.html", nil)
}

func main() {

	Router := mux.NewRouter()
	Router.HandleFunc("/home", SocialNetwork).Methods("GET")
	Router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("js/"))))
	Router.PathPrefix("/views/").Handler(http.StripPrefix("/views/", http.FileServer(http.Dir("views/"))))

	http.ListenAndServeTLS(SSLport, "cert.pem", "key.pem", Router)

}
