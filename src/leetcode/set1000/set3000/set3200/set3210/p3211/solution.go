package p3211

func validStrings(n int) []string {
	buf := make([]byte, n)

	var res []string

	var dfs func(i int)

	dfs = func(i int) {
		if i == n {
			res = append(res, string(buf))
			return
		}
		if i == 0 || buf[i-1] == '1' {
			buf[i] = '0'
			dfs(i + 1)
			buf[i] = '1'
			dfs(i + 1)
		} else {
			buf[i] = '1'
			dfs(i + 1)
		}
	}

	dfs(0)

	return res
}
