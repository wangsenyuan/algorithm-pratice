package main

import "fmt"

func main() {
	fmt.Println(removeKdigits("10200", 1))
}

func removeKdigits(num string, k int) string {
	bytes := []byte(num)
	n := len(bytes)
	removeOneDigit := func() {
		i := 0
		for i < n-1 && bytes[i] <= bytes[i+1] {
			i++
		}
		leading := true
		m := 0
		for j := 0; j < n; j++ {
			if (leading && bytes[j] == '0') || j == i {
				continue
			}
			leading = false
			bytes[m] = bytes[j]
			m++
		}
		n = m
	}

	for i := 0; i < k; i++ {
		removeOneDigit()
	}

	if n == 0 {
		return "0"
	}
	return string(bytes[:n])
}
