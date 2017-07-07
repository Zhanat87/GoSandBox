package main

import "fmt"

/**
 * Extension Method on Struct,
 * Make sure use struct.f()
 * Otherwise wont work
 *
 * @param  url  an absolute URL giving the base location of the image
 * @param  name the location of the image, relative to the url argument
 * @return      the image at the specified URL
 * @see         Image
 * @7/7/2017 3:33 PM
 * Author : cyrsis@github 
 */
type student2 struct {
	name string
	age  int
	width int
	height int
}

func (data student2) area() int {
	return data.width * data.height
}
func main() {
   victor := student2{"Victor",12, 10,20}
	fmt.Println(victor.area())
}
