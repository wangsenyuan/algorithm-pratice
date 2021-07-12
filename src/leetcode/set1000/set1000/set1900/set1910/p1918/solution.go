package p1918

func countPalindromicSubsequence(s string) int {
	// aba
	n := len(s)
	dp := make([]bool, 26)
	dp2 := make([][]bool, 26)
	for i := 0; i < 26; i++ {
		dp2[i] = make([]bool, 26)
	}

	mem := make(map[int]bool)

	for i := 0; i < n; i++ {
		x := int(s[i] - 'a')
		for j := 0; j < 26; j++ {
			if dp2[x][j] {
				mem[x*26+j] = true
			}
		}

		for j := 0; j < 26; j++ {
			if dp[j] {
				dp2[j][x] = true
			}
		}
		dp[x] = true
	}

	return len(mem)
}
