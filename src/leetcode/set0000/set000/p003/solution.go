package p003

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	mp := make(map[byte]int)
	res := 1

	for i, j := 0, 0; j < len(s); j++ {
		c := s[j]

		if k, found := mp[c]; found && k >= i {
			i = k + 1
		}
		mp[c] = j
		res = max(res, j-i+1)
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
