package main

/**
 * Learning Switch Statment
 * https://gobyexample.com/switch
 * @param  url  an absolute URL giving the base location of the image
 * @param  name the location of the image, relative to the url argument
 * @return      the image at the specified URL
 * @see         Image
 * @7/6/2017 6:16 PM
 */
import (
	"time"
	"fmt"
)

func main() {

	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday:
		fmt.Println("This is weekends")
	default:
		fmt.Println("This is Weekdays")
	}

}
