package main

import "fmt"

func main() {
	int32 := 32
	fmt.Println("Point of Variable ",&int32)

	PassPointertoFunction(&int32)
	fmt.Println("After pass into Pointer", int32)
}
func PassPointertoFunction(pointerOfInt *int) {
	*pointerOfInt = 0
}
