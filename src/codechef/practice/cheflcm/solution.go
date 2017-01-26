package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for t > 0 {
		var n int
		fmt.Scanf("%d", &n)

		var res = int64(0)

		for x := 1; x*x <= n; x++ {
			if n%x == 0 {
				res += int64(x)
				y := n / x
				if y != x {
					res += int64(y)
				}
			}
		}

		fmt.Println(res)
		t--

	}
}
