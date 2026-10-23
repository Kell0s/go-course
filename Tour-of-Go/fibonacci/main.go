package main

import "fmt"

func fibonacci() func(int) int {
	ans := 1
	lastAns := 0
	tempData := 0

	return func(x int) int {
		if x == 0 {
			return 0
		} else if x == 1 {
			return 1
		} else {
			tempData = ans
			ans += lastAns
			lastAns = tempData
			return ans
		}
	}
}

func main() {
	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
