package main

import "fmt"

func numIdenticalPairs(nums []int) int {
	cnt := 0

	for i, num1 := range nums {
		for j, num2 := range nums {
			if num1 == num2 && i < j {
				cnt++
			}
		}
	}

	return cnt
}

func main() {
	nums := []int{1, 1, 1, 1}

	fmt.Println(numIdenticalPairs(nums))
}
