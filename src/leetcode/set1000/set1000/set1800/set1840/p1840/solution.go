package p1840

import "sort"

/*
n ~~ 1e9
h[i] <= 1e5
*/
func maxBuilding(n int, restrictions [][]int) int {
	sort.Slice(restrictions, func(i, j int) bool {
		return restrictions[i][0] < restrictions[j][0]
	})

	if len(restrictions) == 0 {
		return n - 1
	}

	// 假设有两个点，(xh, yh) 表示最高点所在的坐标,
	// 并且比当前限制点要高，(xi, yi), 那么满足限制的条件为 xi - xh >= yh - yi
	// 否则 yh = yi + xi - xh
	// 假设前面有个点 (xl, 0), 表示最低点所在的坐标,
	// 满足当前点(xi, yi) 的条件为 xi - xl >= hi
	m := len(restrictions)
	xl := make([]int, m+1)
	xl[0] = 1

	for i := 0; i < len(restrictions); i++ {
		cur := restrictions[i]
		x, h := cur[0], cur[1]
		//stack stores the previous
		xl[i+1] = max(xl[i], x-h)
	}
	best := n - xl[m]
	// xi - xr = h
	var xr = restrictions[m-1][0] + restrictions[m-1][1]
	for i := m - 1; i >= 0; i-- {
		cur := restrictions[i]
		x, h := cur[0], cur[1]
		if xl[i+1] < xr {
			tmp := (xr - xl[i+1]) / 2
			// tmp - xh <= prev - x
			// xh >= tmp - (prev - x)
			// xl[i+1] + tmp < n
			if xl[i+1]+tmp <= n {
				best = max(best, tmp)
			}
		}
		xr = min(xr, x+h)
	}
	return best
}

type Pair struct {
	first, second int
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
