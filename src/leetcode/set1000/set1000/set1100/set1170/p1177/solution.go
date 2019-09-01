package p1177

func canMakePaliQueries(s string, queries [][]int) []bool {
	n := len(s)
	cnt := make([][]int, 26)
	for i := 0; i < 26; i++ {
		cnt[i] = make([]int, n+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < 26; j++ {
			cnt[j][i+1] = cnt[j][i]
		}
		cnt[s[i]-'a'][i+1]++
	}

	ans := make([]bool, len(queries))

	for i, query := range queries {
		left, right, k := query[0], query[1], query[2]
		var odd int

		for j := 0; j < 26; j++ {
			tmp := cnt[j][right+1] - cnt[j][left]
			if tmp%2 == 1 {
				odd++
			}
		}

		ans[i] = (odd / 2) <= k
	}

	return ans
}
