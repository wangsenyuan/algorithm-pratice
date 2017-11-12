package main

import "fmt"

func main() {
	fmt.Println(isHappy(19))
}

func isHappy(n int) bool {

	for n != 89 {
		if n == 1 {
			return true
		}

		m := 0
		for n > 0 {
			m += (n % 10) * (n % 10)
			n /= 10
		}
		n = m
	}
	return false
}
