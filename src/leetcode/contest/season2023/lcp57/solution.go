package lcp57

import "sort"

var dd = []int{-1, 0, 1, 0, -1}

func getMaximumNumber(moles [][]int) int {
	n := len(moles)
	dp := make([]int, 9)
	fp := make([]int, 9)
	gp := make([]int, 9)
	for i := 0; i < 9; i++ {
		dp[i] = -2 * n
	}
	sort.Slice(moles, func(i, j int) bool {
		return moles[i][0] < moles[j][0]
	})
	prev := 0
	dp[4] = 0
	for i := 0; i < n; {
		it := i
		copy(gp, dp)
		for ; i < n && moles[i][0] == moles[it][0]; i++ {
			copy(fp, dp)
			pos := moles[i][1]*3 + moles[i][2]
			cur := moles[i][0]
			diff := cur - prev
			if diff >= 4 {
				// 可以在足够的时间内，移动到任何位置
				x := dp[0]
				for j := 1; j < 9; j++ {
					x = max(x, dp[j])
				}
				for j := 0; j < 9; j++ {
					fp[j] = x
				}
			} else if diff == 3 {
				// 除了对角线不可达，其他位置都可以到达 (0, 8), (2, 6)
				max_at := 1
				for j := 1; j < 8; j++ {
					if j == 2 || j == 6 {
						continue
					}
					if dp[j] > dp[max_at] {
						max_at = j
					}
				}
				v := dp[max_at]
				for j := 1; j < 8; j++ {
					if j == 2 || j == 6 {
						continue
					}
					fp[j] = max_of(fp[j], v, dp[0], dp[2], dp[6], dp[8])
				}
				fp[0] = max_of(dp[0], dp[2], dp[6], v)
				fp[2] = max_of(dp[2], dp[0], dp[8], v)
				fp[6] = max_of(dp[6], dp[0], dp[8], v)
				fp[8] = max_of(dp[8], dp[2], dp[6], v)
			} else if diff == 2 {
				for x := 0; x < 9; x++ {
					a, b := x/3, x%3
					for dx := -2; dx <= 2; dx++ {
						for dy := -2; dy <= 2; dy++ {
							if abs(dx)+abs(dy) > diff {
								continue
							}
							c, d := a+dx, b+dy
							if c >= 0 && c < 3 && d >= 0 && d < 3 {
								fp[x] = max(fp[x], dp[c*3+d])
							}
						}
					}
				}
			} else if diff == 1 {
				for x := 0; x < 9; x++ {
					a, b := x/3, x%3
					for k := 0; k < 4; k++ {
						c, d := a+dd[k], b+dd[k+1]
						if c >= 0 && c < 3 && d >= 0 && d < 3 {
							fp[x] = max(fp[x], dp[c*3+d])
						}
					}
				}
			}
			fp[pos]++
			for x := 0; x < 9; x++ {
				gp[x] = max(gp[x], fp[x])
			}
		}
		copy(dp, gp)
		prev = moles[it][0]
	}
	var res int
	for x := 0; x < 9; x++ {
		res = max(res, dp[x])
	}

	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func max_of(arr ...int) int {
	res := arr[0]
	for j := 1; j < len(arr); j++ {
		res = max(res, arr[j])
	}
	return res
}
