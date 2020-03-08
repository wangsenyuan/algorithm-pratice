package p1371

func findTheLongestSubstring(s string) int {
	cnt := make([]int, 5)

	m := 5

	pos := make([]int, 1<<m)

	for i := 0; i < len(pos); i++ {
		pos[i] = -2
	}

	pos[0] = -1

	var res int

	for i := 0; i < len(s); i++ {
		j := vowelIndex(s[i])
		if j >= 0 {
			cnt[j]++
		}

		var f int

		for k := 0; k < 5; k++ {
			x := cnt[k] & 1
			if x == 1 {
				f |= 1 << k
			}
		}

		if pos[f] == -2 {
			pos[f] = i
		}

		res = max(res, i-pos[f])
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func vowelIndex(b byte) int {
	if b == 'a' {
		return 0
	}
	if b == 'e' {
		return 1
	}

	if b == 'i' {
		return 2
	}

	if b == 'o' {
		return 3
	}

	if b == 'u' {
		return 4
	}
	return -1
}
