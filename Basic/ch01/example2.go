package main

import "fmt"

var steps int = 21

var somethingese = 21

type Message string

func (msg Message) Say() { fmt.Print(msg) }

func (message Message) iSay() { fmt.Print(message) }
func main() {

	msg := Message("Mastering Go in ")
	msg.iSay()
	fmt.Println(steps, " steps.")
}
