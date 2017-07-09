package main

import (
	"fmt"

)

func main() {
	fetchDemo()
	fmt.Println("The main function had been executed")
}
func fetchDemo() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("Recovered a panic. []Index=%d]\n", v)
		}
	}()
	 ss := []string{"A", "B", "C", "D"}
	fmt.Println("Fetching the element in %v one by one...\n",ss)
	fetchElement(ss,0)
	fmt.Println("The element fetch is done")

}
func fetchElement(ss []string, index int) (element string){
 	if index >= len(ss){
		fmt.Println("Occur a panic ![index=%d]\n",index)
		panic(index)
	}
	fmt.Println("Fetching the element... [index = %d]\n",index)
	element= ss[index]
	defer fmt.Println("The element is \"%s\". [index=%d]\n",element,index)
	fetchElement(ss,index+1)
	return
}

