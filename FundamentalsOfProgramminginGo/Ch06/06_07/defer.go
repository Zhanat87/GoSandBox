package main

/**
 * Defer is like using in C#, make sure it execute at the end
 *
 * @param  url  an absolute URL giving the base location of the image
 * @param  name the location of the image, relative to the url argument
 * @return      the image at the specified URL
 * @see         Image
 * @08/07/2017 16:33
 * Author : cyrsis@github 
 */
import (
	"fmt"
)

func main() {
	defer fmt.Println("Close the file!")
	fmt.Println("Open the file!")

	defer fmt.Println("Statement 1")
	defer fmt.Println("Statement 2")
	
	myFunc()
	
	defer fmt.Println("Statement 3")
	defer fmt.Println("Statement 4")
	fmt.Println("Undeferred statement")
	
	x := 1000
	defer fmt.Println("Value of x:", x)
	x++
	fmt.Println("Value of x after incrementing:", x)
}

func myFunc() {
	defer fmt.Println("")
	defer fmt.Println("Deferred in the function")
	fmt.Println("Not deferred in the function")
}
