package main

import (
	"fmt"
)

func cal_single(prefix []byte, s []byte, cur int) {
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
			cal_single(prefix, s, cur+1)
		}

	}
}

