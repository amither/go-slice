package main

import (
	"fmt"
)

func cal(req chan int, out chan int) {
	go func() {
		num, ok := <-req
		if !ok {
			close(out)
			return
		}
		fmt.Println(num)
		nr := make(chan int)
		cal(nr, out)
		for x := range req {
			if x%num != 0 {
				nr <- x
			}
		}
		close(nr)

	}()
}
func main() {
	req, out := make(chan int), make(chan int)
	cal(req, out)
	for num := 2; num < 10; num++ {
		req <- num
	}
	close(req)
	<-out
}
