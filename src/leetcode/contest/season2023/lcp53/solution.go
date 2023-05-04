package lcp53

import "math"

func defendSpaceCity(time, position []int) int {
	const n, m = 100, 1 << 5
	rain := [n + 1]int{}
	for i, t := range time {
		rain[position[i]] |= 1 << (t - 1)
	}

	var union, single [m]int
	for i := 1; i < m; i++ {
		lb := i & -i
		j := i ^ lb
		lb2 := j & -j
		if lb == lb2>>1 { // 两个时间点相邻
			union[i] = union[j] + 1
			single[i] = single[j] + 1 // 递推
		} else {
			// 若 i 的二进制包含 101，对于联合屏障选择继续维持是更优的，
			// 不过下面的 DP 已经枚举了所有的情况，自然会选择更优的方案
			union[i] = union[j] + 3
			single[i] = single[j] + 2
		}
	}

	f := [n + 1][m]int{}
	for j := range f[0] {
		f[0][j] = union[j] + single[(m-1^j)&rain[0]]
	}
	for i := 1; i <= n; i++ {
		for j := range f[i] {
			f[i][j] = math.MaxInt32 / 2
			mask := m - 1 ^ j
			for pre := mask; ; pre = (pre - 1) & mask { // 枚举 j 的补集 mask 中的子集 pre
				cost := f[i-1][pre] + union[j] + single[(mask^pre)&rain[i]]
				f[i][j] = min(f[i][j], cost)
				if pre == 0 {
					break
				}
			}
		}
	}
	return f[n][0]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
