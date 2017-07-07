package main

import "fmt"

/**
 * Interface is really hard in Go, come back when I have time
 * https://gobyexample.com/interfaces
 *
 * @param  url  an absolute URL giving the base location of the image
 * @param  name the location of the image, relative to the url argument
 * @return      the image at the specified URL
 * @see         Image
 * @7/7/2017 4:02 PM
 * Author : cyrsis@github 
 */
type gemetry interface {
	area() float64
	perim() float64
}

type rect struct {
	width float64

	height float64
}

func (structData rect) area() float64 {
	return structData.width * structData.height
}
func (structData rect) perim() float64 {
	return 2 * structData.width * 2 * structData.height
}
func measure(g gemetry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	measure(r)
}
