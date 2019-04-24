package main

import "fmt"

func permutaionCal(left chan []byte, right chan []byte, s []byte) {
	for v := range left {
		prefixIncrement(v, s, right)
	}
	close(right)
}

func PermutaionChain(s []byte) {
	leftmost := make(chan []byte)
	left := leftmost
	right := leftmost

	for i := 0; i < len(s)-1; i++ {
		right = make(chan []byte)
		go permutaionCal(left, right, s)
		left = right
	}

	go func(c chan []byte) {
		for _, v := range s {
			c <- []byte{v}
		}
		close(c)
	}(leftmost)

	for v := range right {
		fmt.Println(v)
	}
}
