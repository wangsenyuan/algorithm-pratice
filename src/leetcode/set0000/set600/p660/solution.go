package main

import "fmt"

func newInteger(n int) int {
	ans, base := 0, 1

	for n > 0 {
		ans += n % 9 * base
		n /= 9
		base *= 10
	}

	return ans
}

func main() {
	fmt.Println(newInteger(9))
	fmt.Println(newInteger(18))

}
