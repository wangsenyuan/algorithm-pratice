package main

import "fmt"

func main() {
	s := "D"
	fmt.Println(findPermutation(s))
}

func findPermutation(s string) []int {
	s = s + "I"
	n := len(s)
	res := make([]int, n)
	idx := 1

	descCount := 0
	for i := 0; i < n; i++ {
		if s[i] == 'I' {
			for j := i; descCount >= 0; j-- {
				res[j] = idx
				idx++
				descCount--
			}
		}
		descCount++
	}

	return res
}
