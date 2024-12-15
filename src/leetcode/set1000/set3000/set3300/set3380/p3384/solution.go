package p3384

import "slices"

const inf = 1 << 30

func makeStringGood(s string) int {
	// 在给定长度时，找出最小的代价? 可行吗？
	// 删除和添加时比较好处理的
	// 第三个操作，貌似可以省下一步操作,
	//  一个字符只会变化一次，不需要变化2次上，a -> b -> c
	// 因为变化两次，可以被当作删除一个a, 增加一个c
	// dp[i][j][k] 表示平均值是j,且从i-1变化了k个下来
	// k <= j， 且如果有从i-1变化出来的话，就不可能有从i-2变化到i。
	// 否则的话， 就相当于从i-2变化了一部分到i

	freq := make([]int, 26)
	for _, x := range []byte(s) {
		freq[int(x-'a')]++
	}

	n := len(s)

	x := slices.Max(freq)

	var dfs func(i int, exp int, x int) int

	dp := make([]map[int]int, 26)
	for i := 0; i < 26; i++ {
		dp[i] = make(map[int]int)
	}

	dfs = func(i int, exp int, x int) (ans int) {
		if i == 26 {
			if x == 0 {
				return 0
			}
			return inf
		}

		if v, ok := dp[i][x]; ok {
			return v
		}

		defer func() {
			dp[i][x] = ans
		}()

		if x > 0 {
			if x+freq[i] <= exp {
				ans = min(exp-(x+freq[i]), x+freq[i]) + dfs(i+1, exp, 0)
			} else {
				// x + freq[i] >= exp
				ans = max(0, freq[i]-exp) + dfs(i+1, exp, 0)
			}
		} else {
			if freq[i] == 0 {
				ans = dfs(i+1, exp, 0)
			} else if freq[i] <= exp {
				ans = min(exp-freq[i], freq[i]) + dfs(i+1, exp, 0)
				ans = min(ans, dfs(i+1, exp, freq[i])+freq[i])
			} else {
				ans = min(dfs(i+1, exp, freq[i]-exp), dfs(i+1, exp, 0)) + freq[i] - exp
			}
		}

		return ans
	}

	res := n

	for i := 1; i <= x; i++ {
		for j := 0; j < 26; j++ {
			dp[j] = make(map[int]int)
		}
		res = min(res, dfs(0, i, 0))
	}
	return res
}
