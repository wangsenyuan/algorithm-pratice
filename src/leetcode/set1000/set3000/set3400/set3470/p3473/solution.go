package p3473

const inf = 1 << 60

func maxSum(nums []int, k, m int) int {
	n := len(nums)
	s := make([]int, n+1)
	for i, x := range nums {
		s[i+1] = s[i] + x
	}

	f := make([]int, n+1)
	d := make([]int, n+1)
	for i := 1; i <= k; i++ {
		for j, v := range f {
			d[j] = v - s[j]
		}
		for j := i*m - m; j < i*m; j++ {
			f[j] = -inf // 即使 [0,j) 全选，也没有 i 个长为 m 的子数组
		}
		mx := -inf
		// 左右两边留出足够空间
		for j := i * m; j <= n-(k-i)*m; j++ {
			// mx 表示最大的 f[L]-s[L]，其中 L 在区间 [(i-1)*m, j-m] 中
			mx = max(mx, d[j-m])
			f[j] = max(f[j-1], mx+s[j]) // 不选 vs 选
		}
	}
	return f[n]
}
