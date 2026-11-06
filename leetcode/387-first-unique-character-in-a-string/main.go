package main

import "fmt"

func firstUniqChar(s string) int {
	seen := make(map[rune]int)

	for _, letter := range s {
		seen[letter]++
	}

	for i, k := range s {
		if seen[k] == 1 {
			return i
		}
	}

	return -1
}

func main() {
	str := "loveleetcode"
	fmt.Println(firstUniqChar(str))
}
