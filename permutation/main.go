package main

import "fmt"

func main() {
	s := []byte{'1', '2', '3', '4'}

	fmt.Println("normal version")
	Permutation(s)

	// fmt.Println("cocurrency version")
	// PermutationConcurrency(s)

	fmt.Println("concurrency vertical version")
	PermutationConcurrencyVertical(s)
}
z