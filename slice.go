package main

import "fmt"

/**
 * Slice is much better than array and it is a interface
 *
 * @param  url  an absolute URL giving the base location of the image
 * @param  name the location of the image, relative to the url argument
 * @return      the image at the specified URL
 * @see         Image
 * @7/7/2017 12:53 PM
 */

func main() {
	sliceString := []string{"a", "b", "b"}
	fmt.Println(sliceString)

	sliceVar := []string{"bb", "cc", "dd"}
	fmt.Println(sliceVar)

	fmt.Println(cap(sliceVar))
}