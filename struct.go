package main

import "fmt"

type student struct {
	name string
	age  int
}
func main() {

	victor :=student{"Victor",12}
	fmt.Println("The student namme is ", victor.name , victor.age ," and Somethinge lse")
}