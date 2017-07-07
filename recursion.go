package main

import "fmt"

/**
 * Recursion
 * Function that call itself
 *
 * @param  url  an absolute URL giving the base location of the image
 * @param  name the location of the image, relative to the url argument
 * @return      the image at the specified URL
 * @see         Image
 * @7/7/2017 2:31 PM
 */
func main() {
	fmt.Println(recursion2(7))

}

func recursion2(n int) int {
	if n == 0 {
		return 1
	}
	return n * recursion2(n-1)
}
