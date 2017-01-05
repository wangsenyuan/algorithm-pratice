package main

import "fmt"

func main() {
	for i := 1; i < 9; i++ {
		fmt.Println(largestPalindrome(i))
	}
}

func largestPalindrome(n int) int {
	if n == 1 {
		return 9
	}

	upper := repeat(9, n)

	max := upper * upper
	half := max / (upper + 1)
	pl := int64(0)
	found := false
	for !found {
		pl = palindrome(half, upper+1)

		for x := upper; x*x >= pl && !found; x-- {
			if pl/x > max {
				break
			}
			found = pl%x == 0
		}
		half--
	}
	return int(pl % 1337)
}
func palindrome(first int64, base int64) int64 {
	second := int64(0)
	x := first
	for x > 0 {
		i := x % 10
		second = second*10 + i
		x = x / 10
	}
	return first*base + second
}

func repeat(x int, n int) int64 {
	if n == 0 {
		return 0
	}

	return repeat(x, n-1)*10 + int64(x)
}
