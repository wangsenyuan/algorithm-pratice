package p1781

func beautySum(s string) int {

	n := len(s)
	cnt := make([][]int, n)
	for i := 0; i < n; i++ {
		cnt[i] = make([]int, 26)
	}
	var res int
	for j := 0; j < n; j++ {
		if j > 0 {
			copy(cnt[j], cnt[j-1])
		}
		cnt[j][idx(s[j])]++
		for i := 0; i < j; i++ {
			var a, b = 1, n
			for k := 0; k < 26; k++ {
				x := cnt[j][k]
				if i > 0 {
					x -= cnt[i-1][k]
				}
				if x == 0 {
					continue
				}
				a = max(a, x)
				b = min(b, x)
			}
			res += a - b
		}
	}

	return res
}

func idx(x byte) int {
	return int(x - 'a')
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	return a + b - max(a, b)
}
