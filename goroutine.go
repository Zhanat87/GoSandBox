package main

import (
	"fmt"
	"time"
)

/**
 * GoRoutine is kind of crazy
 *
 * @param  url  an absolute URL giving the base location of the image
 * @param  name the location of the image, relative to the url argument
 * @return      the image at the specified URL
 * @see         Image
 * @09/07/2017 00:01
 * Author : cyrsis@github 
 */
func main() {

	go fmt.Println("Go! Goroutine")
    //Didnt print anything until

	time.Sleep(time.Millisecond) //Finally would print this out

	go func() {
		fmt.Println("Go! GoRoutine in Different Way")
	}()

	time.Sleep(time.Millisecond)
}
