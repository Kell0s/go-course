package main

import "fmt"

func numIdenticalPairs(nums []int) int {
	m := map[int]int{}
	cnt := 0

	for _, num := range nums {
		cnt += m[num]
		m[num]++
	}

	return cnt
}

func main() {
	arr := []int{1, 2, 3, 1, 1, 3}

	fmt.Println(numIdenticalPairs(arr))
}
