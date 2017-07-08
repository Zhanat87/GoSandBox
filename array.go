package main

import "fmt"

func main() {
	victor := [5]int{1, 2, 3, 4, 5}

	jenny := [...]int{2, 3, 5, 6, 6}// Confirm the Dot Dot Dot is working in array list

	fmt.Println(victor)
	fmt.Println(jenny)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
