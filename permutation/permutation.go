package main

import (
	"fmt"
)

func cal1(in []byte, s []byte, next chan []byte) {
    if len(in) == len(s) {
        fmt.Println(in)
    } else {
        for _, c := range s {
            exist := false
            for _, e := range in {
                if e == c {
                    exist = true
                    break
                }
            }
            if exist {
                continue
            } 

            temp := make([]byte, 0) 
            temp = append(temp, in...)
            temp = append(temp, c)
            next <- temp
        }
    }
} 

func cal(req chan []byte, out chan []byte, s []byte) {
	go func() {
		//end
        v, ok := <- req

        if !ok {
            close(out)
            return
        }

		next := make(chan []byte)
		cal(next, out, s)

        cal1(v, s, next)
		for in := range req {
            cal1(in, s, next)
		}
		close(next)
	}()
}

func main() {
	s := make([]byte, 0)
	for c := byte('a'); c <= 'd'; c++ {
		s = append(s, c)
	}

	req, out := make(chan []byte), make(chan []byte)
	cal(req, out, s)

	for _, c := range s {
		sl := make([]byte,0)
		sl = append(sl, c)
		req <- sl
	}
	close(req)

	<-out
}
