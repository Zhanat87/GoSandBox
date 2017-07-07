package main

import "fmt"

/**
 * 
 * Go supports anonymous functions, which can form closures. Anonymous functions are useful * when you want to define a function inline without having to name it.

 * The effect in the state of the function variable, How does it userful?
 *
 * @param  url  an absolute URL giving the base location of the image
 * @param  name the location of the image, relative to the url argument
 * @return      the image at the specified URL
 * @see         Image
 * @7/6/2017 6:20 PM
 */
func main() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

}

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1

		return i
	}
}
