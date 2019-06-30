package p1103

const INF = 1 << 20

func distributeCandies(candies int, n int) []int {
	// 1, n + 1, 2 * n + 1, 3 * n + 1 ...  ==> 0 * n + 1 + n + 1 + 2 * n + 1 + 3 * n + 1 =>
	// (0 + 1 + 2 + 3 + k - 1) * n + k
	// 2, n + 2, 2 * n + 2, 3 * n + 2 ...
	// k * n * n + (n + 1) * n / 2
	// (0 + k - 1) * k / 2 * n + (1 + 2 + 3 + .. k) * k =>
	// n * (k - 1) * k / 2 + (1 + k) * k / 2
	a := n * n
	b := n * (n + 1) / 2
	k := search(1, INF, func(k int) bool {
		return a*k*(k-1)/2+b*k > candies
	})

	k--

	res := make([]int, n)

	for i := 1; i <= n; i++ {
		res[i-1] = n*k*(k-1)/2 + k*i
		candies -= res[i-1]
	}

	for i := 1; i <= n && candies > 0; i++ {
		x := min(k*n+i, candies)
		res[i-1] += x
		candies -= x
	}

	return res
}

func search(left int, right int, fn func(int) bool) int {
	for left < right {
		mid := (left + right) / 2
		if fn(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
