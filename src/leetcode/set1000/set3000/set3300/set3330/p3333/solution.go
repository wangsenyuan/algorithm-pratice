package p3333

func possibleStringCount(word string, k int) int {
	if len(word) < k { // 无法满足要求
		return 0
	}

	const mod = 1_000_000_007
	cnts := []int{}
	ans := 1
	cnt := 0
	for i := range word {
		cnt++
		if i == len(word)-1 || word[i] != word[i+1] {
			if len(cnts) < k {
				cnts = append(cnts, cnt)
			}
			ans = ans * cnt % mod
			cnt = 0
		}
	}

	m := len(cnts)
	if m >= k { // 任何输入的字符串都至少为 k
		return ans
	}

	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, k)
	}
	f[0][0] = 1
	s := make([]int, k+1)
	for i, c := range cnts {
		for j, v := range f[i] {
			s[j+1] = (s[j] + v) % mod
		}
		// j <= i 的 f[i][j] 都是 0
		for j := i + 1; j < k; j++ {
			f[i+1][j] = s[j] - s[max(j-c, 0)]
		}
	}

	for _, v := range f[m][m:] {
		ans -= v
	}
	return (ans%mod + mod) % mod // 保证结果非负
}
