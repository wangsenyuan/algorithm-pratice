package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for t > 0 {
		t--
		var n, k int64
		fmt.Scanf("%d %d", &n, &k)
		fmt.Println(solve(n, k))
	}
}

const mod = 1000000007

func solve(n int64, k int64) int64 {
	total := (k - 1) / (n - 1)
	first := k - 1
	last := first - total*(n-1)

	ans := first + last

	total++

	if ans%2 == 0 {
		ans /= 2
	} else {
		total /= 2
	}

	ans = ((ans % mod) * (total % mod)) % mod

	return ans
}
