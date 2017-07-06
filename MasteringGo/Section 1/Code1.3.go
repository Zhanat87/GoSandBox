package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	response, err := http.Get("http://api.theysaidso.com/qod")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close() //What is defer?

	//defer用于资源的释放，会在函数返回之前进行调用

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(contents))

}
