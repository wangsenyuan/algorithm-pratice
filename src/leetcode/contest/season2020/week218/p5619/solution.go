package p5619

import "sort"

const INF = 1 << 30

func minimumIncompatibility1(nums []int, k int) int {
	res := check(nums, k)
	if res == -1 {
		return -1
	}
	if res >= 0 {
		return res
	}
	sort.Ints(nums)
	n := len(nums)
	N := 1 << n
	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = INF
		}
	}
	per := n / k
	for p := 0; p < n; p++ {
		dp[0][p] = 0
	}

	for mask := 1; mask < N; mask++ {
		if digitCount(mask)%per == 0 {
			for p := 0; p < n; p++ {
				if mask&(1<<p) > 0 {
					dp[mask][0] = min(dp[mask][0], dp[mask^(1<<p)][p])
				}
			}
			for p := 1; p < n; p++ {
				dp[mask][p] = dp[mask][0]
			}
		} else {
			for pre := 0; pre < n; pre++ {
				for p := pre + 1; p < n; p++ {
					if mask&(1<<p) > 0 && nums[p] > nums[pre] {
						dp[mask][pre] = min(dp[mask][pre], dp[mask^(1<<p)][p]+nums[p]-nums[pre])
					}
				}
			}
		}
	}
	return dp[N-1][0]
}

func digitCount(num int) int {
	var res int
	for num > 0 {
		res++
		num -= num & -num
	}
	return res
}

func check(nums []int, k int) int {
	n := len(nums)

	if k == n {
		return 0
	}

	// 每个子集中有n/k的元素
	cnt := make(map[int]int)
	var a, b = INF, 0
	for i := 0; i < n; i++ {
		a = min(a, nums[i])
		b = max(b, nums[i])
		cnt[nums[i]]++
		if cnt[nums[i]] > k {
			return -1
		}
	}

	if k == 1 {
		return b - a
	}
	return -2
}

func minimumIncompatibility(nums []int, k int) int {
	res := check(nums, k)
	if res == -1 {
		return -1
	}
	if res >= 0 {
		return res
	}

	n := len(nums)
	N := 1 << n
	var P []Pair
	sort.Ints(nums)
	var pick func(i int, mask int, prev int, a int, b int)
	x := n / k
	pick = func(i int, mask int, prev int, a int, b int) {
		if i == x {
			P = append(P, Pair{mask, b - a})
			return
		}

		for j := prev + 1; j < n; j++ {
			if prev < 0 || nums[j] > nums[prev] {
				pick(i+1, mask|(1<<j), j, min(a, nums[j]), nums[j])
			}
		}
	}
	pick(0, 0, -1, n+1, 0)
	// then need to pick k sets, such that set a & b == 0 && sum(a, b, ..) = 0
	m := len(P)
	// C(m, k)
	dp := make([]int, N)
	for i := 0; i < N; i++ {
		dp[i] = INF
	}
	dp[0] = 0
	for mask := 0; mask < N; mask++ {
		if dp[mask] == INF {
			continue
		}
		for i := 0; i < m; i++ {
			if mask&P[i].first > 0 {
				continue
			}
			dp[mask|P[i].first] = min(dp[mask|P[i].first], dp[mask]+P[i].second)
		}
	}

	return dp[N-1]
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

type Pair struct {
	first  int
	second int
}
