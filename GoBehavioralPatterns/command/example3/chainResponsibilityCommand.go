package main

/* Chan of Responsibility of Command Patterns
 * Pass a string message between links
 *
 * @param  url  an absolute URL giving the base location of the image
 * @param  name the location of the image, relative to the url argument
 * @return      the image at the specified URL
 * @see         Image
 * @7/6/2017 3:19 PM
 */

import (
	"fmt"
	"time"
)

type Command interface {
	Info() string
}

type ChainLogger interface {
	Next(Command)
}

type TimePassed struct {
	start time.Time
}

func (t *TimePassed) Info() string {
	return time.Since(t.start).String()
}

type Logger struct {
	NextChain ChainLogger
}

func (f *Logger) Next(c Command) {
	time.Sleep(time.Second)

	fmt.Printf("Elapsed time from creation: %s\n", c.Info())

	if f.NextChain != nil {
		f.NextChain.Next(c)
	}
}

func main() {
	second := new(Logger)
	first := Logger{NextChain: second}

	command := &TimePassed{start: time.Now()}

	first.Next(command)
}
