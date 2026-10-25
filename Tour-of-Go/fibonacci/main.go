package main

import (
	"fmt"
)

func fibonacci() func() int {
	ans := 1
	lastAns := 0

	return func() int {
		res := lastAns
		lastAns, ans = ans, lastAns+ans
		return res
	}
}

func main() {
	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
