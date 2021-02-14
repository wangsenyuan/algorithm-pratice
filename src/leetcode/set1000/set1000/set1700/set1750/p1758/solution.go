package p1758

func minOperations(s string) int {
	return min(find(s, '0'), find(s, '1'))
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func find(s string, x byte) int {
	var res int

	for i := 0; i < len(s); i++ {
		if i&1 == 0 {
			if s[i] != x {
				res++
			}
		} else {
			if s[i] == x {
				res++
			}
		}
	}

	return res
}
