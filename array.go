package main

import "fmt"

func main() {
	victor := [5]int{1, 2, 3, 4, 5}

	fmt.Println(victor)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
