package main

import "fmt"

func numIdenticalPairs(nums []int) int {
	seen := map[int]int{}
	cnt := 0

	for _, num := range nums {
		cnt += seen[num]
		seen[num]++
	}

	return cnt
}

func main() {
	nums := []int{1, 2, 3, 1, 1, 3}

	fmt.Println(numIdenticalPairs(nums))
}
