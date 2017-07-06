package main

import (
	"os"
	
)

func main() {
	panic("Panic and stop execution");
	_, error := os.Create("tmp/panicfile")
	if error != nil {
		panic(error)
	}

}
