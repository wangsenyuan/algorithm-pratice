package p1864

func minSwaps(s string) int {
	var cnt int
	for _, x := range s {
		if x == '1' {
			cnt++
		}
	}
	n := len(s)

	if n%2 == 0 {
		if cnt == n/2 {
			x := swapCount(s, '1')
			y := swapCount(s, '0')
			return min(x, y)
		}
	} else {
		if cnt == n/2 {
			return swapCount(s, '0')
		}

		if cnt == (n+1)/2 {
			return swapCount(s, '1')
		}
	}

	return -1
}

func swapCount(s string, x byte) int {
	var cnt int
	for i := 0; i < len(s); i += 2 {
		if s[i] != x {
			cnt++
		}
	}
	return cnt
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
