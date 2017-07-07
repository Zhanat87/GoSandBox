package main

import "fmt"

/**
 * Variadic Function , Just add ...types for vararg
 *
 * @param  url  an absolute URL giving the base location of the image
 * @param  name the location of the image, relative to the url argument
 * @return      the image at the specified URL
 * @see         Image
 * @7/7/2017 1:41 PM
 */
func main() {
	//sum(1,2,3,5)
	total(2,4,6,8)
}

func total(nums ...int) {
	for k, v := range nums{
		fmt.Printf("%s -> %s\n", k, v)
	}
}
func sum(nums ...int) {
	total := 0

	fmt.Println(nums," ")
	for v := range nums {

		total += v
	}
	fmt.Println(total)

	for k, v := range nums {
		fmt.Printf("%s -> %s\n", k, v)
	}

}
