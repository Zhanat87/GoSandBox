package main

import (
	//DEMO 1, 2, 3, 4, & 5
	"net/http"

	//DEMO 2 & 3
	//"io/ioutil"

	//DEMO 3
	//"strings"

	//DEMO 4
	//"os"
	//"bufio"

)

func main() {
	//DEMO 1
	/*
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
			w.Write([]byte("Hello World"))
		})

	http.ListenAndServe(":8000", nil)
	*/

	//DEMO 2, 3, 4
	/*
	http.Handle("/", new(MyHandler))
	*/

	//DEMO 5
	http.ListenAndServe(":8000", http.FileServer(http.Dir("public")))
}

//DEMO 2, 3, 4
type MyHandler struct {
	http.Handler
}

//DEMO 2
/*
func (this *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := "public" + req.URL.Path
	data, err := ioutil.ReadFile(string(path))

	if err == nil {
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 - " + http.StatusText(404)))
	}

}
*/

//DEMO 3
/*
func (this *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := "public" + req.URL.Path
	data, err := ioutil.ReadFile(string(path))

	if err == nil {
		var contentType string
		if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "application/javascript"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "image/png"
		} else {
			contentType = "text/plain"
		}

		w.Header().Add("Content Type", contentType)
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 - " + http.StatusText(404)))
	}
}
*/

//DEMO 4
/*
func (this *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := "public" + req.URL.Path
	f, err := os.Open(path)

	if err == nil {
		var bufferedReader = bufio.NewReader(f)

		var contentType string
		if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "application/javascript"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "image/png"
		} else {
			contentType = "text/plain"
		}

		w.Header().Add("Content Type", contentType)
		bufferedReader.WriteTo(w)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 - " + http.StatusText(404)))
	}
}
*/