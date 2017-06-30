package main

import
(
	"github.com/nkozyra/hyphen/admin"
	"log"
)



func main() {

	log.Println("Starting hyphen servier")
	admin.Init()
	admin.StartServer()

}