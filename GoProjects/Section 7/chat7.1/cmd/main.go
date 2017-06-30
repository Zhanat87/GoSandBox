package main

import (
	"fmt"
	"net/http"
	"os"
)

var sites = []string{
	"https://github.com",
	"https://google.com",
	"https://stackoverflow.com",
	"https://facebook.com",
	"https://twitter.com",
	"https://golang.org",
	"https://forum.golangbridge.org",
	"https://packtpub.com/",
}

func get() {
	for _, s := range sites {
		res, _ := http.Get(s)
		fmt.Printf("%s %d\n", s, res.StatusCode)
		res.Body.Close()
	}
}

func getConcurrently() {
	ch := make(chan string)

	for _, s := range sites {
		go func(s string) {
			res, _ := http.Get(s)
			ch <- fmt.Sprintf("%s %d", s, res.StatusCode)
			res.Body.Close()
		}(s)
	}

	for range sites {
		fmt.Println(<-ch)
	}
}

func main() {
	switch os.Args[1] {
	case "seq":
		get()
	case "conc":
		getConcurrently()
	default:
		fmt.Println("Please choose `seq` or `conc`.")
	}
}
