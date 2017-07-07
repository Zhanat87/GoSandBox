package main

/**
 * Error in Go is complicated too
 * https://blog.golang.org/error-handling-and-go
 * Check this out when I have time
 *
 * @param  url  an absolute URL giving the base location of the image
 * @param  name the location of the image, relative to the url argument
 * @return      the image at the specified URL
 * @see         Image
 * @7/7/2017 6:07 PM
 * Author : cyrsis@github 
 */
import "fmt"

func main() {

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
type argError struct {
	arg  int
	prob string
}
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

//In this case we use &argError syntax to build a new struct, supplying values for the two fields arg and prob.
func f2(arg int) (int, error) {
	if arg == 42 {

return -1, &argError{arg, "can't work with it"}
}
return arg + 3, nil
}
