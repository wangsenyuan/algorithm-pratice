package main

import "fmt"

func main() {
	fmt.Println(removeKdigits("1234567890", 9))
	fmt.Println(removeKdigits("1432219", 3))
	fmt.Println(removeKdigits("10200", 1))
	fmt.Println(removeKdigits("10", 2))
}

func removeKdigits(num string, k int) string {
	n := len(num)
	stack := make([]byte, n)
	p := 0

	for i := 0; i < n; i++ {
		for p > 0 && k > 0 && stack[p-1] > num[i] {
			p--
			k--
		}
		stack[p] = num[i]
		p++
	}

	for k > 0 {
		p--
		k--
	}

	m := 0
	for m < p && stack[m] == '0' {
		m++
	}
	if m == p {
		return "0"
	}
	return string(stack[m:p])
}
