package main

import (
	"strconv"
	"fmt"
)

func main() {
	fmt.Println(nextGreaterElement(12))
	fmt.Println(nextGreaterElement(21))
	fmt.Println(nextGreaterElement(132))
	fmt.Println(nextGreaterElement(123))

	fmt.Println(nextGreaterElement(110))
}

func nextGreaterElement(n int) int {
	s := []byte(strconv.Itoa(n))

	i := len(s) - 2
	for i >= 0 && s[i] >= s[i+1] {
		i--
	}

	if i < 0 {
		return -1
	}

	j := len(s) - 1
	for s[j] <= s[i] {
		j--
	}

	s[i], s[j] = s[j], s[i]

	i, j = i+1, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	x, _ := strconv.Atoi(string(s))
	return x
}
