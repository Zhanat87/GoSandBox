package main

/**
 * Range can loop over all kind of structure
 *
 * @param  url  an absolute URL giving the base location of the image
 * @param  name the location of the image, relative to the url argument
 * @return      the image at the specified URL
 * @see         Image
 * @7/7/2017 1:33 PM
 */
import "fmt"

func main() {
	mapKStringVInt := map[string]string{"foo": "foov1", "bar": "barv1"}
	for k, v := range mapKStringVInt {
		fmt.Printf("%s -> %s\n", k, v)
	}
}