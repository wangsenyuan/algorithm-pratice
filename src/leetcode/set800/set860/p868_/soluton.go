package p868_

func binaryGap(N int) int {
	var ans int
	j := -1

	for i := 0; i < 32; i++ {
		if N&1 == 1 {
			if j >= 0 && i-j > ans {
				ans = i - j
			}
			j = i
		}
		N >>= 1
	}

	return ans
}
