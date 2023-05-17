package p014

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	n := len(strs)
	for j := 0; ; j++ {
		for i := 0; i < n; i++ {
			if len(strs[i]) == j || i > 0 && strs[i][j] != strs[0][j] {
				return strs[i][:j]
			}
		}
	}

	return ""
}
