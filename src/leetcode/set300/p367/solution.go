package main

import "fmt"

func main() {
	fmt.Println(isPerfectSquare(16))
	fmt.Println(isPerfectSquare(14))
}

func isPerfectSquare(num int) bool {
	seen := make(map[int]bool)

	guess := 1
	seen[guess] = true

	for guess*guess != num {
		guess = (guess + num/guess) / 2
		if guess <= 0 || seen[guess] {
			return false
		}
		seen[guess] = true
	}

	return true
}

func isPerfectSquare1(n int) bool {
	if n == 1 || n == 4 {
		return true
	}

	if n < 4 {
		return false
	}

	for i := 2; i*2 < n; i++ {
		if n == i*i {
			return true
		}
	}

	return false
}
