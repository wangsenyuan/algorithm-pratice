package p3438

import "sort"

const inf = 1 << 60

func minimumIncrements(nums []int, target []int) int {
	sort.Ints(nums)
	sort.Ints(target)

	m := len(target)
	M := 1 << m
	// len(target) <= 4
	dp := make([]int, M)
	fp := make([]int, M)

	for i := range M {
		dp[i] = inf
		fp[i] = inf
	}
	dp[0] = 0

	for _, num := range nums {
		// 如果使用num得到一个结果
		for state := 0; state < M; state++ {
			sub := (M - 1) ^ state
			for next := sub; next > 0; next = (next - 1) & sub {
				var l = 1
				for j := 0; j < m; j++ {
					if (next>>j)&1 == 1 {
						l = lcm(l, target[j])
					}
				}
				if l < num {
					l = (num + l - 1) / l * l
				}
				fp[next|state] = min(fp[next|state], dp[state]+l-num)
			}
		}
		for state := range M {
			dp[state] = min(dp[state], fp[state])
			fp[state] = inf
		}
	}

	return dp[M-1]
}

func lcm(a, b int) int {
	g := gcd(a, b)
	return a / g * b
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
