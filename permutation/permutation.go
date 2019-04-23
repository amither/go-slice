package main

import (
	"fmt"
	"sync"
)

func prefixIncrement(in []byte, s []byte, next chan []byte) {
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

func permutaionConImpl(req chan []byte, out chan []byte, s []byte) {
	go func() {
		//递归退出条件: len(v) == len(s)-1
		v, ok := <-req
		if !ok {
			return
		}

		next := out
		if len(v) != len(s)-1 {
			next = make(chan []byte)
			permutaionConImpl(next, out, s)
		}

		prefixIncrement(v, s, next)
		for in := range req {
			prefixIncrement(in, s, next)
		}
		close(next)
	}()
}

// PermutationConcurrency  并发计算全排列
func PermutationConcurrency(s []byte) {
	req, out := make(chan []byte), make(chan []byte)

	//开启goroutine计算
	permutaionConImpl(req, out, s)

	over := make(chan struct{})
	go func() {
		for res := range out {
			fmt.Println(res)
		}
		close(over)
	}()

	for _, c := range s {
		sl := []byte{c}
		req <- sl
	}
	close(req)

	<-over

}

func PermutationConcurrencyVertical(s []byte) {
	var sg sync.WaitGroup
	for _, c := range s {
		sg.Add(1)

		//要传参数进去，避免只取到最后一个值
		go func(k byte) {
			p := make([]byte, len(s), len(s))
			p[0] = k
			permutaionImpl(p, s, 1)
			defer sg.Done()
		}(c)
	}

	sg.Wait()
}
