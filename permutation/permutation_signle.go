package main

import (
	"fmt"
)

func permutaionImpl(prefix []byte, s []byte, cur int) {
	if cur == len(s) {
		fmt.Println(prefix)
		return
	}

	for _, b := range s {
		exist := false
		for i := 0; i < cur; i++ {
			if prefix[i] == b {
				exist = true
				break
			}
		}

		if !exist {
			prefix[cur] = b
			permutaionImpl(prefix, s, cur+1)
		}

	}
}

func Permutation(s []byte) {
	//前缀slice
	p := make([]byte, len(s), len(s))
	permutaionImpl(p, s, 0)
}
