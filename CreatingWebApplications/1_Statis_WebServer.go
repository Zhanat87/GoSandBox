package main

import "net/http"

func main() {  //{ Have to be following after main, otherwise it would add ; to the end to make compiler error

	//Func can live by itself in Go

	http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
		//The funcs in here is annoymouse function and , *http.Request is a point to a object

		//main is limited to local package since, but Write is public for all
	 w.Write([]byte("Hello World and u would like it"))
		w.Write([]byte("What else do u want?"))


	})

	http.ListenAndServe(":8000", nil)

}

type MyHandler struct{
	http.HandlerFunc
}

func (this *MyHandler) ServeHttp(w http.ResponseWriter, req *http.Request){

}