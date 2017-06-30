package middleware

import (
	"fmt"
	"net/http"
)

// Add is a variadic function that adds up numbers
func Add(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// Chain holds our sum
type Chain struct {
	Sum int
}

// AddNext is chainable sum function
func (c *Chain) AddNext(num int) *Chain {
	c.Sum += num
	return c
}

// Finally is for use at the end of the chain, and returns the sum
func (c *Chain) Finally(num int) int {
	return c.Sum + num
}

// Next runs the next function in the chain
func Next(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Before")
		next.ServeHTTP(w, r)
		fmt.Println("After")
	})
}
