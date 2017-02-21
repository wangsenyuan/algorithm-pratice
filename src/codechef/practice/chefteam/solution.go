package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for t > 0 {
		var n, k int
		fmt.Scanf("%d%d", &n, &k)

		if k > n {
			fmt.Println(0)
		} else {
			if k > n-k {
				k = n - k
			}
			var c int64 = 1
			for j := 1; j <= k; j++ {
				d := gcd(c, int64(j))
				c /= d
				c *= int64(n-j+1) / (int64(j) / d)
			}
			fmt.Println(c)
		}

		t--
	}
}

func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
