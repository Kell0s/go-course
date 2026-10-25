package main

import (
	"fmt"
)

func fibonacci() func() int {
	ans := 1
	lastAns := 0
	totalAns := 0

	return func() int {
		totalAns = lastAns
		lastAns, ans = ans, lastAns+ans
		return totalAns
	}
}

func main() {
	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
