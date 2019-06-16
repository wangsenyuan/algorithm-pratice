package p784

func letterCasePermutation(S string) []string {

	var dfs func(i int)
	res := make([]string, 0, 100)
	bs := []byte(S)
	dfs = func(i int) {
		if i == len(S) {
			res = append(res, makeCopy(bs))
			return
		}
		dfs(i + 1)

		if bs[i] >= 'a' && bs[i] <= 'z' {
			bs[i] = bs[i] - 'a' + 'A'
			dfs(i + 1)
			bs[i] = bs[i] - 'A' + 'a'
		} else if bs[i] >= 'A' && bs[i] <= 'Z' {
			bs[i] = bs[i] - 'A' + 'a'
			dfs(i + 1)
			bs[i] = bs[i] - 'a' + 'A'
		}
	}

	dfs(0)

	return res
}

func makeCopy(bs []byte) string {
	tmp := make([]byte, len(bs))
	copy(tmp, bs)
	return string(tmp)
}
