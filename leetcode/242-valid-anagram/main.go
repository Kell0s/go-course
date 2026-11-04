package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[rune]int)

	for _, ch := range s {
		m[ch]++
	}

	for _, ch := range t {
		m[ch]--
		if m[ch] < 0 {
			return false
		}
	}

	return true
}

func main() {
	s := "aba"
	t := "aba"
	fmt.Println(isAnagram(s, t))
}
