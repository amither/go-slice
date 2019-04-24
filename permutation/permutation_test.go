package main

import (
	"fmt"
	"testing"
)

func TestChain(t *testing.T) {
	leftmost := make(chan int)
	left := leftmost
	right := leftmost
	for i := 0; i < 10; i++ {
		right = make(chan int)
		go func(c1, c2 chan int) {
			c1 <- 1 + <-c2
		}(left, right)
		left = right
	}

	go func(c chan int) {
		c <- 1
	}(right)

	fmt.Println(<-leftmost)
}
