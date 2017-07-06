package main

import "fmt"

func main() {
	s := struct {
		i int
		s string
	}{i: 2, s: "Victor"}
	fmt.Printf("%v \n", s)
	fmt.Println(s.i, s.s)
}
