package main

import "fmt"

func main() {
	var t, n int
	fmt.Scanf("%d", &t)

	for t > 0 {
		fmt.Scanf("%d", &n)

		cnt := 0
		for x := 1; x*x <= n; x++ {
			if n%x != 0 {
				continue
			}

			if isOverLucky(x) {
				cnt++
			}

			y := n / x
			if y != x && isOverLucky(y) {
				cnt++
			}
		}

		fmt.Println(cnt)

		t--
	}
}
func isOverLucky(x int) bool {
	for x > 0 {
		m := x % 10
		if m == 4 || m == 7 {
			return true
		}
		x = x / 10
	}

	return false
}
