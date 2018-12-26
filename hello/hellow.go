package main

import (
	"fmt"
	"go-slice/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("hello,world"))
    x, _ := stringutil.Swap("hello","world")
	fmt.Println(x)
}
