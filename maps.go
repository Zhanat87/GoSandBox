package main

import "fmt"

/**
 * Maps are call dicts or hash in other language
 *
 * @param  url  an absolute URL giving the base location of the image
 * @param  name the location of the image, relative to the url argument
 * @return      the image at the specified URL
 * @see         Image
 * @7/7/2017 1:07 PM
 */
func main() {
	maps := make(map[string]int)
	maps["key1"] = 7
	maps["key2"] = 13
	fmt.Println(maps)

	delete(maps,"key1")
	fmt.Println("After delete the key1",maps)

	mape := map[string]int{"foo": 1, "bar": 2}

	fmt.Println(mape)
}