package chain

/**
 * Dynamically chain the action at runtime
 * Pass a request through a chain of processors
 *
 * --Make ChainLogger's signature return Next(String) error
 * --Changing receivedMessage field to a pointer
 *
 * --Widely used to create Finte State machine(FSM)
 * --Use interchangerabley with decorater patterns

 * It helps building a chain of objects.
 * Request enters from one end and keeps
 * going from object to object till it finds the suitable handler.
 *
 * @7/6/2017 12:32 PM
 */

import (
	"fmt"
	"io"
	"strings"
)

type ChainLogger interface {
	Next(string)
}

type FirstLogger struct {
	NextChain ChainLogger
}

func (f *FirstLogger) Next(s string) {
	fmt.Printf("First logger: %s\n", s)

	if f.NextChain != nil {
		f.NextChain.Next(s)
	}
}

type SecondLogger struct {
	NextChain ChainLogger
}

func (se *SecondLogger) Next(s string) {
	if strings.Contains(strings.ToLower(s), "hello") {
		fmt.Printf("Second logger: %s\n", s)

		if se.NextChain != nil {
			se.NextChain.Next(s)
		}

		return
	}

	fmt.Printf("Finishing in second logging\n\n")
}

type WriterLogger struct {
	NextChain ChainLogger
	Writer    io.Writer
}

func (w *WriterLogger) Next(s string) {
	if w.Writer != nil {
		w.Writer.Write([]byte("WriterLogger: " + s))
	}

	if w.NextChain != nil {
		w.NextChain.Next(s)
	}
}

type ClosureChain struct {
	NextChain ChainLogger
	Closure   func(string)
}

func (c *ClosureChain) Next(s string) {
	if c.Closure != nil {
		c.Closure(s)
	}

	if c.NextChain != nil {
		c.Next(s)
	}
}
