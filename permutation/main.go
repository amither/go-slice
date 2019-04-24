package main

import "fmt"

func testChain() {
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

func main() {
	s := []byte{'1', '2', '3', '4'}

	// fmt.Println("normal version")
	// Permutation(s)

	// // fmt.Println("cocurrency version")
	// // PermutationConcurrency(s)

	// fmt.Println("concurrency vertical version")
	// PermutationConcurrencyVertical(s)

	// testChain()

	fmt.Println("concurrency chain version")
	PermutaionChain(s)
}
