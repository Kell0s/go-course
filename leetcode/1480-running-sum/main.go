package main

import "fmt"

func runningSum(nums []int) []int {
	lastNum := 0
	for i, num := range nums {
		nums[i] = num + lastNum
		lastNum += num
	}

	return nums
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(runningSum(nums))
}
