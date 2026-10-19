package main

import "fmt"

func maximumWealth(accounts [][]int) int {
	maxMoney := 0

	for _, customer := range accounts {
		i := 0

		for _, bank := range customer {
			i += bank
		}

		if maxMoney < i {
			maxMoney = i
		}
	}
	return maxMoney
}

func main() {
	accounts := [][]int{
		{1, 5},
		{7, 3},
		{3, 5},
	}
	fmt.Println(maximumWealth(accounts))
}
