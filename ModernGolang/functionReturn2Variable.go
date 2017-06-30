package main

import (
	fmt "fmt"


)

func main() {
	fmt.Println(ComapreTwoNumber(12, 12))
	someint := 2
	switch someint {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println("Switch statment with 2")
	case 3:
		fmt.Println(3)
	}
}
func ComapreTwoNumber(i int, i2 int) (bool,int) {
	if i>i2 {
		return false, i-i2
		
	} else if i2 > i {
		 return false, i2-i
	}
	fmt.Println("They are requal")
	return true, 0


}
