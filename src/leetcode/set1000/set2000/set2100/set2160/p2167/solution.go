package p2167

const INF = 1 << 30

func minimumTime(s string) int {
	n := len(s)
	// n <= 2 * 100000
	// n**2 won't pass
	// 应该是在某个i处，删除所有的头部，cost (i + 1), 在j后面删除所有的尾部, cost(n - j - 1)
	// 然后删除i 和 j中间的1，cost 2 * cnt
	// i + 1 + 2 * (sum(j) - sum[i]) + n - j - 1
	// i + 2 * sum(j) - 2 * sum[i]
	// i - 2 * sumn[i] 最小
	var minValue int
	best := n
	var sum int
	for i := 1; i <= n; i++ {
		sum += int(s[i-1] - '0')

		best = min(best, n-i+2*sum+minValue)
		if i-2*sum < minValue {
			minValue = i - 2*sum
		}
	}
	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
