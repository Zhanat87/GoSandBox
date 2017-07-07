package main

import "fmt"

func main() {
	k,v := getMeTwoValue()

	fmt.Println("Key is %s, Value is %s", k,v)
}
func getMeTwoValue() (int, int) {
	return 2,3
}
