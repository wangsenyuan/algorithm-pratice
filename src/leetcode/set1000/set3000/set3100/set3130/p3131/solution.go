package p3131

func minEnd(n int, x int) int64 {
	// a[0] & a[1] ... & a[n-1] = x
	// 也就是说 a[i] 都是x的超集
	// 且 a[0] < a[1] < ... < a[n-1]
	// 那么 a[0] = x
	// 然后 x的位数去掉后，剩下的其他位递增
	// x + n - 1
	res := x
	delta := n - 1

	// res + delta

	for i, j := 0, 0; delta > 0; i++ {
		if (x>>i)&1 == 1 {
			continue
		}
		v := (delta >> j) & 1
		res += v << i
		delta -= v << j
		j++
	}

	return int64(res)
}
